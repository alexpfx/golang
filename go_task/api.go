package go_task

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RtcSession interface {
	Connect()
}

const (
	workItemPath string = "/ccm/rpt/repository/workitem"
	authPath     string = "/ccm/auth/j_security_check"
)

func NewCcmSession(username string, password string, host string) RtcSession {
	return ccmSession{
		username: username,
		password: password,
		host:     host,
	}
}

type ccmSession struct {
	username string
	password string
	host     string
	cookies  []*http.Cookie
}

func (c ccmSession) Connect() {
	authUrl := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   authPath,
	}

	fmt.Println(authUrl.String())

	values := url.Values{
		"j_username": {c.username},
		"j_password": {c.password},
	}
	response, err := http.PostForm(authUrl.String(), values)

	//buf := bytes.NewBuffer([]byte(fmt.Sprintf(`
	//	"j_username": {%s},
	//	"j_password": {%s}`, c.username, c.password)))

	fmt.Println(authUrl.String())
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, authUrl.String(), strings.NewReader(values.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	response, err = client.Do(request)

	check(err)
	defer response.Body.Close()

	c.cookies = response.Cookies()
	for i, cookie := range c.cookies {
		fmt.Print(i)
		fmt.Println(cookie.Name)
	}

	body, err := ioutil.ReadAll(response.Body)

	check(err)

	log.Println(string(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}

}
