package go_task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type RtcSession interface {
	Get(taskId string) []Task
}

const (
	workItemPath string = "/jts/identity"
	//workItemPath  string = "/ccm/rpt/repository/workitem"
	authPath      string = "/ccm/jts/j_security_check"
	authUrl       string = "https://alm.dataprev.gov.br/ccm/auth/j_security_check"
	xComIbmHeader string = "x-com-ibm-team-repository-web-auth-msg"

	curl         string = `https://alm.dataprev.gov.br/ccm/rpt/repository/workitem?fields=workitem/workItem[id=%s]/(id|summary|type/name|href|description|owner/userId|resolver/userId|creator/userId|category/name|reportableUrl|plannedEndDate|foundIn/name|timeSpent|duration|plannedStartDate|activationDate|reportableUrl|creationDate|teamArea/name|comments/content|timeSheetEntries)`
	pathWorkItem string = `/ccm/rpt/repository/workitem`

	pathWorkItemFields string = "fields=workitem/workItem[id=%s]/(%s)"
)

func NewCcmSession(username string, password string, host string) RtcSession {
	return ccmSession{
		username: username,
		password: password,
		host:     host,
	}
}

type ccmSession struct {
	username  string
	password  string
	host      string
	cookieJar *http.CookieJar
}

func (c ccmSession) Get(taskId string) []Task {
	j, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: nil})

	result, connected, jar := c.get(taskId, "id|summary", j)
	if !connected {
		jar = c.auth(nil)
		result, connected, jar = c.get(taskId, "id|summary", jar)
		if !connected {

			fmt.Println("Não pode conectar")
			return nil
		}
		return result
	}
	return result
}

func logUrl(url url.URL) {
	fmt.Println(url.String())
}

func (c *ccmSession) get(id string, filter string, jar http.CookieJar) (result []Task, isConnected bool, rJar http.CookieJar) {
	url := url.URL{
		Scheme:   "https",
		Host:     c.host,
		Path:     pathWorkItem,
		RawQuery: fmt.Sprintf(pathWorkItemFields, id, filter),
	}
	client := clientWithJar(jar, false)

	request, _ := http.NewRequest(http.MethodGet, url.String(), strings.NewReader(""))

	resp, err := client.Do(request)

	check(err)
	if needAuth(resp) {
		return nil, false, client.Jar
	}

	return convertResult(resp), true, client.Jar

}

func (c *ccmSession) auth(jar http.CookieJar) http.CookieJar {
	values := url.Values{
		"j_username": {c.username},
		"j_password": {c.password},
	}

	client := clientWithJar(jar, false)

	resp, err := client.PostForm(authUrl, values)
	h := resp.Header.Get("set-cookie")
	all, err := ioutil.ReadAll(resp.Body)
	fmt.Println("all: ", string(all))

	fmt.Println(resp.Status)
	fmt.Println("cookie: ", h)
	check(err)
	return client.Jar
}

func convertResult(resp *http.Response) []Task {
	return nil
}

func needAuth(resp *http.Response) bool {
	h := resp.Header.Get("X-com-ibm-team-repository-web-auth-msg")
	if h == "authrequired" {
		return true
	}
	return false
}

func newClient(jar http.CookieJar, redirect bool) *http.Client {
	var checkRedirect func(*http.Request, []*http.Request) error

	if !redirect {
		checkRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	c := &http.Client{
		CheckRedirect: checkRedirect, Jar: jar,
	}

	return c
}
func clientWithJar(jar http.CookieJar, redirect bool) *http.Client {
	return newClient(jar, false)
}

func client() *http.Client {
	return newClient(nil, true)
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}

}
