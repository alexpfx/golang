package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const Base = "https://www-scm.prevnet/api/v3/projects/754"
const MergeRequestsPath = "merge_requests"
const CommitsPath = "repository/commits"

var (
	isDebug = false
	token   string
	isRange bool
	infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Obtém informações sobre um ou mais merges requests e retorna um json",
		Args: func(cmd *cobra.Command, args []string) error { // validador
			isDebug, _ = cmd.Root().PersistentFlags().GetBool("debug")
			isRange, _ = cmd.Flags().GetBool("range")

			if isRange {

				if len(args) < 2 {
					return fmt.Errorf("requer um intervalo, ex: 'merge info -i 7004 7009'")
				}
				p0, e0 := strconv.Atoi(args[0])
				p1, e1 := strconv.Atoi(args[1])

				if e0 != nil || e1 != nil {
					return fmt.Errorf("requer um intervalo, ex: 'merge info -i 7004 7009'")
				}

				if p0 > p1 {
					return fmt.Errorf("requer um intervalo, onde o primeiro elemento é menor ou igual ao segundo: 'merge info -i 7004 7009'")
				}

				lim := 30
				if (p1 - p0) > lim {
					return fmt.Errorf("máximo de %d MRs por consulta'", lim)
				}

			}

			for _, n := range args {
				if _, e := strconv.Atoi(n); e != nil {
					return fmt.Errorf("parámetro não é um número: %s", n)
				}

			}

			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			token = getPrivateKeyArgument(cmd)
			targetBranchs := getTargetBranchsArgument(cmd)

			mergeIds := extractMrs(args, isRange)

			client := getClient()
			okMrList := make([]main.MRResult, 0)
			errorMrList := make([]main.MRErrResult, 0)
			if isDebug {
				fmt.Println("private token: ", token)
			}

			for _, mergeId := range mergeIds {

				slowdown := rand.Intn(500)
				time.Sleep(time.Duration(slowdown))

				func() {

					url := fmt.Sprintf("%s/%s/?iid=%d", Base, MergeRequestsPath, mergeId)

					if isDebug {
						fmt.Println("obtendo MR: ", url)
					}

					req := createRequest(url, token)

					r, err := client.Do(req)

					if err != nil {
						e := fmt.Sprintf("Erro na requisição: %s\n", url)
						errorMrList = appendError(errorMrList, mergeId, e)
						return
					}

					defer r.Body.Close() //TODO ver onde fechar isso

					body, err := ioutil.ReadAll(r.Body)

					var merges []main.Merge

					err = json.Unmarshal(body, &merges)

					if err != nil {
						e := fmt.Sprintf("Erro na conversão de objeto para json: %s", url)
						errorMrList = appendError(errorMrList, mergeId, e)
						return
					}

					if len(merges) < 1 {
						e := fmt.Sprintf("MR inexistente: %d", mergeId)
						errorMrList = appendError(errorMrList, mergeId, e)
						return

					}

					var merge = merges[0]
					author, _ := cmd.Flags().GetString("author")

					if !shouldAddThisBranch(targetBranchs, merge) {
						return
					}

					if shouldAddThisAuthor(author, merge) {
						mergeCommit, err := getCommit(merge.MergeCommitSha)
						if err != nil {
							e := fmt.Sprintf("Merge request ainda não aceito: %v", mergeId)
							errorMrList = appendError(errorMrList, mergeId, e)
							return
						}
						okMrList = appendResult(okMrList, merge, mergeCommit)
					} else if isDebug {
						fmt.Println("ignorando MR de autor filtrado: ", author)
					}
				}()

			}

			outputJson(okMrList)

		},
	}
)

func getTargetBranchsArgument(cmd *cobra.Command) []string {
	targetBranchs, _ := cmd.Flags().GetStringSlice("targets")
	sort.Strings(targetBranchs)
	return targetBranchs
}

func shouldAddThisBranch(branchs []string, merge main.Merge) bool {
	sliceLen := len(branchs)
	if sliceLen == 0 {
		return true
	}

	// se encontrou retorna o índice senão retorna o tamanho do slice.
	return sort.SearchStrings(branchs, merge.TargetBranch) != sliceLen
}

func outputJson(result []main.MRResult) {
	for i, mrResult := range result {
		createdAt := mrResult.MergeCommit.CreatedAt
		if createdAt != "" {
			t, _ := time.Parse(time.RFC3339, createdAt)
			result[i].MergeCommit.CreatedAt = t.Format("2006-01-02T15:04:05Z")
			//2006-11-12T11:45:26.000Z

		}
	}



	bytes, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func getPrivateKeyArgument(cmd *cobra.Command) string {
	token, _ := cmd.Flags().GetString("PRIVATE_TOKEN")
	return token
}

func shouldAddThisAuthor(author string, merge main.Merge) bool {
	if author == "" {
		return true
	}

	return merge.Author.Username == author
}

func createRequest(url string, token string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("PRIVATE-TOKEN", token)
	return req
}

func getClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr}
}

func getCommit(commitSha string) (main.MergeCommit, error) {
	url := fmt.Sprintf("%s/%s/%s", Base, CommitsPath, commitSha)
	client := getClient()
	var commit main.MergeCommit

	req := createRequest(url, token)

	r, e := client.Do(req)
	if e != nil {
		return commit, e
	}

	body, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return commit, e
	}

	e = json.Unmarshal(body, &commit)
	if e != nil {
		return commit, e
	}

	return main.MergeCommit{
		Id:        commitSha,
		CreatedAt: commit.CreatedAt,
		Email:     commit.Email,
		Username: strings.FieldsFunc(commit.Email, func(r rune) bool {
			return r == '@'
		})[0],
	}, e
}

func appendError(results []main.MRErrResult, mergeId int, err string) []main.MRErrResult {
	return append(results, main.MRErrResult{
		MergeId: mergeId,
		Err:     err,
	})
}

func appendResult(results []main.MRResult, merge main.Merge, mergeCommit main.MergeCommit) []main.MRResult {

	return append(results, main.MRResult{
		Merge:       merge,
		MergeCommit: mergeCommit,
	})
}

func extractMrs(args []string, isRange bool) []int {
	var numbers []int

	if isRange {
		p0, _ := strconv.Atoi(args[0])
		p1, _ := strconv.Atoi(args[1])

		if p1 < p0 {
			panic("ops!")
		}

		for i := p0; i <= p1; i++ {
			numbers = append(numbers, i)
		}

		return numbers
	}

	for _, s := range args {
		i, _ := strconv.Atoi(s)
		numbers = append(numbers, i)

	}
	return numbers

}

func init() {

	infoCmd.Flags().BoolP("range", "i", false, "-i <mergeId_inicial> <mergeId_final>")
	infoCmd.Flags().StringP("author", "a", "", "-a <author>")
	infoCmd.Flags().StringSliceP("targets", "b", []string{}, "-b <target_branch>")
	infoCmd.Flags().String("PRIVATE_TOKEN", os.Getenv("PRIVATE_TOKEN"), "--PRIVATE_TOKEN <seu_token> ou passar via variável de ambiente.")

	rootCmd.AddCommand(infoCmd)
}

//PRIVATE_TOKEN=xZgEwsr2nhz121qmiFLd ./merge info 8868 | jq -r '.[]|add|"\(.web_url)\t\(.author.username)\t\(.username)\t\(.created_at)"' | xsel -b
