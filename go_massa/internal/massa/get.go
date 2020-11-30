package massa

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const path_recupera_recente = "massa/recuperaMaisRecente"

type MassaGetter interface {
	GetRecent(catalogo, ambiente int) (Massa, error)
}

type massaGetter struct {
	client *http.Client
}

func NewMassaGetter() MassaGetter {
	return massaGetter{
		client: createClient(),
	}
}

func createClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr}
}

func (c massaGetter) GetRecent(catalogo, ambiente int) (Massa, error) {
	url := strings.Join([]string{baseUrl, path_recupera_recente, strconv.Itoa(catalogo), strconv.Itoa(ambiente)}, "/")
	req, _ := http.NewRequest(http.MethodPut, url, nil)
	response, err := c.client.Do(req)
	if err != nil {
		return Massa{}, fmt.Errorf("erro ao recuperar massa ambiente %d e catalogo %d \n%v", ambiente, catalogo, err.Error())
	}

	var massa Massa
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &massa)
	if err != nil {
		log.Fatal(err)
	}

	return massa, nil
}
