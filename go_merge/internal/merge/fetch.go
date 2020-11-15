package merge

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const mergeRequestsPath = "merge_requests"
const mergeRequestQuery = "?iid="
const commitsPath = "repository/commits"

// args:
// cmd <url> <url>
// cmd <mergeId> <mergeId>
// cmd mergeId:mergeId
var lastNumberRegex = regexp.MustCompile("[0-9]+")
var colonInMid = regexp.MustCompile(`[0-9]+\:[0-9]+`)

func createUrl(base, project, path, query string) string {
	return strings.Join([]string{base, project, path, query}, "/")
}

func GetFetchMode(args []string) (ranges []int, err error) {
	for i, arg := range args {
		if i == 0 {
			continue
		}

		if !isRange(arg) {
			valid, id := validateSingleMrId(arg)
			if !valid {
				continue
			}
			ranges = append(ranges, id)
			continue
		}

		valid, first, last := validateRange(arg)
		if !valid {
			continue
		}

		ranges = append(ranges, buildRange(first, last)...)
	}
	return nil, nil
}

func validateSingleMrId(arg string) (bool, int) {
	return false, -1
}

func buildRange(first int, last int) []int {
	return []int{}
}

func validateRange(arg string) (valid bool, first, last int) {
	valid = false

	splitted := strings.Split(arg, ":")
	mi := splitted[0]
	mf := splitted[2]

	first, err := strconv.Atoi(mi)
	if err != nil {
		return
	}

	last, err = strconv.Atoi(mf)
	if err != nil {
		return
	}

	if last <= first {
		return
	}
	valid = true
	return
}

func isRange(arg string) bool {
	return colonInMid.MatchString(arg)
}

func extractIds(urls []string) (result []int, err error) {

	for _, url := range urls {
		mrIdStr := lastNumberRegex.FindAllString(url, -1)
		n := len(mrIdStr)
		if n < 1 {
			return nil, fmt.Errorf("url de merge inválida: %s", url)
		}
		mrIdInt, err := strconv.Atoi(mrIdStr[n-1])
		if err != nil {
			return nil, err
		}
		result = append(result, mrIdInt)
	}

	return
}

func fetch(token, baseUrl, project string, mrs []int, filter map[string]string) (result []MRResult, err error) {
	sort.Ints(mrs)

	mrList := make([]MRResult, 0)
	errMrList := make([]MRErrResult, 0)

	client := createClient()

	for _, mrId := range mrs {
		url := createUrl(baseUrl, project, mergeRequestsPath, mergeRequestQuery+strconv.Itoa(mrId))
		req := createRequest(url, token)

		resp, err := client.Do(req)

		if err != nil {
			errOut := fmt.Errorf("não foi possível obter %d: ", mrId)
			errMrList = appendError(errMrList, mrId, errOut.Error())
			continue
		}

		var merges []Merge
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &merges)
		if err != nil {
			errOut := fmt.Errorf("erro na conversão de objeto para json: %d", mrId)
			errMrList = appendError(errMrList, mrId, errOut.Error())
			continue
		}
		if len(merges) < 1 {
			err := fmt.Errorf("objeto MR inexistente: %d", mrId)
			errMrList = appendError(errMrList, mrId, err.Error())
			continue
		}

		for _, merge := range merges {
			mrList, err = addOrDiscard(baseUrl, project, merge, mrList, filter, token)
			if err != nil {
				errMrList = appendError(errMrList, mrId, err.Error())
			}
		}
	}
	return mrList, nil

}

func addOrDiscard(baseUrl string, project string, merge Merge, mrList []MRResult, filter map[string]string, token string) ([]MRResult, error) {

	if filter == nil {
		commit, err := fetchCommit(baseUrl, project, merge.MergeCommitSha, token)
		if err != nil {
			return mrList, err
		}
		return appendResult(mrList, merge, commit), nil
	}

	for k, v := range filter {

		if strings.EqualFold(k, "author") {
			if merge.Author.Username != v {
				continue
			}
		}
		if strings.EqualFold(k, "target_branch") {
			if !strings.EqualFold(v, merge.TargetBranch) {
				continue
			}
		}
		// filtros
		commit, err := fetchCommit(baseUrl, project, merge.MergeCommitSha, token)
		if err != nil {
			return mrList, err
		}
		return appendResult(mrList, merge, commit), nil
	}

	return mrList, nil

}

func createRequest(url string, token string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("PRIVATE-TOKEN", token)
	return req
}
func appendResult(results []MRResult, merge Merge, commit Commit) []MRResult {

	return append(results, MRResult{
		Merge:       merge,
		MergeCommit: commit,
	})
}

func appendError(results []MRErrResult, mergeId int, err string) []MRErrResult {
	return append(results, MRErrResult{
		MergeId: mergeId,
		Err:     err,
	})
}

func createClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr}
}

func fetchCommit(baseUrl, project, commitSha, token string) (Commit, error) {
	url := createUrl(baseUrl, project, commitsPath, commitSha)
	client := createClient()
	var commit Commit

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

	return Commit{
		Id:        commitSha,
		CreatedAt: commit.CreatedAt,
		Email:     commit.Email,
		Username: strings.FieldsFunc(commit.Email, func(r rune) bool {
			return r == '@'
		})[0],
	}, e
}
