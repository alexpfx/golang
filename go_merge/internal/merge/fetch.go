package merge

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const Base = "https://www-scm.prevnet/api/v3/projects/754"
const MergeRequestsPath = "merge_requests"
const CommitsPath = "repository/commits"

func Fetch(token, baseUrl string, mrs []int, filter map[string]string) (result []MRResult, err error) {
	sort.Ints(mrs)

	mrList := make([]MRResult, 0)
	errMrList := make([]MRErrResult, 0)

	client := createClient()

	for _, mrId := range mrs {
		url := strings.Join([]string{baseUrl, strconv.Itoa(mrId)}, "")
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
			mrList, err = addOrDiscard(merge, mrList, filter, token)
			if err != nil {
				errMrList = appendError(errMrList, mrId, err.Error())
			}
		}
	}
	return mrList, nil

}

func addOrDiscard(merge Merge, mrList []MRResult, filter map[string]string, token string) ([]MRResult, error) {

	if filter == nil {
		commit, err := fetchCommit(merge.MergeCommitSha, token)
		if err != nil{
			return mrList, err
		}
		return appendResult(mrList, merge, commit), nil
	}

	for k, v := range filter {

		if strings.EqualFold(k, "author") {
			if merge.Author.Username != v{
				continue
			}
		}
		if strings.EqualFold(k, "target_branch"){
			if !strings.EqualFold(v, merge.TargetBranch){
				continue
			}
		}
		// filtros
		commit, err := fetchCommit(merge.MergeCommitSha, token)
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

func fetchCommit(commitSha, token string) (Commit, error) {
	url := fmt.Sprintf("%s/%s/%s", Base, CommitsPath, commitSha)
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
