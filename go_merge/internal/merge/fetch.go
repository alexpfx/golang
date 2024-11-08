package merge

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const mergeRequestsPath = "merge_requests"
const mergeRequestQuery = "?iid="
const commitsPath = "repository/commits"

var colonInMid = regexp.MustCompile(`[0-9]+[:|-][0-9]+`)

// args:
// cmd <url> <url>
// cmd <mergeId> <mergeId>
// cmd mergeId:mergeId
//var lastNumberRegex = regexp.MustCompile("[0-9]+")

var lastNumberRegex = regexp.MustCompile("[0-9]+")

func createUrl(base, project, path, query string) string {
	return strings.Join([]string{base, project, path, query}, "/")
}

func ParseIds(args []string) (ranges []int, err error) {

	for _, arg := range args {

		if isRange(arg) {
			valid, first, last := validateRange(arg)
			if valid {
				ranges = append(ranges, buildRange(first, last)...)
			}
			continue
		}
		id, valid := parseUrlOrSingleId(arg)

		if !valid {
			continue
		}
		ranges = append(ranges, id)

	}
	if len(ranges) == 0 {
		return nil, fmt.Errorf("passe um ou mais merge ids ou urls como parâmetro")
	}

	return ranges, nil
}

func parseUrlOrSingleId(arg string) (result int, valid bool) {
	var parsedUrl, err = url.ParseRequestURI(arg)
	if err != nil {
		return parseSingleId(arg)
	}

	var split = lastNumberRegex.FindAllString(parsedUrl.Path, -1)

	size := len(split)
	if size < 1 {
		return -1, false
	}

	a := split[size-1]
	n, err := strconv.Atoi(a)
	if err != nil || n < 0 {
		return -1, false
	}
	return n, true
}

func parseSingleId(arg string) (int, bool) {
	n, err := strconv.Atoi(arg)
	if err != nil || n < 0 {
		return -1, false
	}
	return n, true
}

func buildRange(first int, last int) []int {
	var r []int
	for i := first; i <= last; i++ {
		r = append(r, i)
	}
	return r
}

func validateRange(arg string) (valid bool, first, last int) {
	valid = false

	var splitted []string
	splitted = strings.Split(arg, ":")
	if len(splitted) != 2 {
		splitted = strings.Split(arg, "-")
	}
	mi := splitted[0]
	mf := splitted[1]

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

func ToJsonStr(results interface{}) string {
	bytes, err := json.MarshalIndent(results, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func Fetch(token, baseUrl, project string, mrs []int, filter map[string]string) ([]MRResult, []MRErrResult, error) {
	sort.Ints(mrs)

		mrList := make([]MRResult, 0)
	errMrList := make([]MRErrResult, 0)

	client := createClient()

	for _, mrId := range mrs {
		mrUrl := createUrl(baseUrl, project, mergeRequestsPath, mergeRequestQuery+strconv.Itoa(mrId))

		req := createRequest(mrUrl, token)

		resp, err := client.Do(req)

		if err != nil {
			errOut := fmt.Errorf("%v", err.Error())
			errMrList = appendError(errMrList, mrId, errOut.Error())
			continue
		}

		var merges []Merge
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &merges)
		if err != nil {
			errOut := fmt.Errorf("erro na conversão de objeto para json: %d %v", mrId, err)
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
	return mrList, errMrList, nil

}

func addOrDiscard(baseUrl string, project string, merge Merge, mrList []MRResult, filter map[string]string, token string) ([]MRResult, error) {

	if filter == nil || len(filter) < 1 {
		commit, err := fetchCommit(baseUrl, project, merge.MergeCommitSha, token)
		if err != nil {
			return mrList, err
		}
		return appendResult(mrList, merge, commit), nil
	}

	for k, v := range filter {

		if strings.EqualFold(k, "author") {
			if merge.Author.Username != v {
				return mrList, nil
			}
		}
		if strings.EqualFold(k, "target_branch") {
			if !strings.EqualFold(v, merge.TargetBranch) {
				return mrList, nil
			}
		}
	}

	if merge.MergeCommitSha == "" {
		return mrList, nil
	} // filtros

	commit, err := fetchCommit(baseUrl, project, merge.MergeCommitSha, token)
	if err != nil {
		return mrList, err
	}
	return appendResult(mrList, merge, commit), nil
}

func createRequest(url string, token string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("PRIVATE-TOKEN", token)
	return req
}
func appendResult(results []MRResult, merge Merge, commit Commit) []MRResult {
	merge.Commit = commit
	return append(results, MRResult{
		Merge: merge,
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

	createdAt, _ := time.Parse(time.RFC3339, commit.CreatedAt)

	return Commit{
		Id:        commitSha,
		CreatedAt: createdAt.Format("2006-01-02T15:04:05Z"),
		Email:     commit.Email,
		Username: strings.FieldsFunc(commit.Email, func(r rune) bool {
			return r == '@'
		})[0],
	}, e
}
