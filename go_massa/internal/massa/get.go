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

const pathRecuperaRecente = "massa/recuperaMaisRecente"
const pathRecuperaAntiga = "massa/recuperaMaisAntiga"

type Retriever interface {
	Newest(catalogo, ambiente int) (Massa, error)
	Oldest(catalogo, ambiente int) (Massa, error)
}

type retriever struct {
	client *http.Client
}

func (c retriever) Newest(catalogo, ambiente int) (Massa, error) {
	url := strings.Join([]string{baseUrl, pathRecuperaRecente, strconv.Itoa(catalogo), strconv.Itoa(ambiente)}, "/")
	return c.retrieve(catalogo, ambiente, url)
}

func (c retriever) Oldest(catalogo, ambiente int) (Massa, error) {
	url := strings.Join([]string{baseUrl, pathRecuperaAntiga, strconv.Itoa(catalogo), strconv.Itoa(ambiente)}, "/")
	return c.retrieve(catalogo, ambiente, url)
}

func NewRetriever() Retriever {
	return retriever{
		client: createClient(),
	}
}

func createClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr}
}

func (c retriever) retrieve(catalogo int, ambiente int, url string) (Massa, error) {
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
