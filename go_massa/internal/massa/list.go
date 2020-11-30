package massa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)


const path_listagem_catalogos = "/bmfrontend/listagemCatalogosComEstoquesFrontend"

type CatalogoList interface {
	Catalogos(ambiente int) ([]Catalogo, error)
}

func NewCatalogoList() CatalogoList{
	return catalogoList{}
}

type catalogoList struct {

}

func (c catalogoList) Catalogos(ambiente int) ([]Catalogo, error) {
	url := strings.Join([]string{baseUrl, path_listagem_catalogos, strconv.Itoa(ambiente)}, "/")
	get, err := http.Get(url)
	if err != nil {
		return []Catalogo{}, fmt.Errorf("erro ao obter lista de catálogos para o ambiente %d\n%v", ambiente, err.Error())
	}

	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return []Catalogo{}, fmt.Errorf("erro ao transformar string json \n%v", err.Error())
	}

	var list []Catalogo
	json.Unmarshal(body, &list)
	return list, nil
}

