package output

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"strings"
)

func Filter(jsonInput string, filter []string) string {
	if len(filter) == 0 {
		return jsonInput
	}
	var aux interface{}
	json.Unmarshal([]byte(jsonInput), &aux)
	m := aux.(map[string]interface{})

	for i, s := range filter {
		split := strings.Split(s, " ")
		if len(split) < 1 || len(split) > 2 {
			log.Fatal("filtro inválido: ", s)
		}

		field := split[0]

		if _, ok := m[field]; !ok {
			continue
		}
		if len(split) > 1 {
			alias := split[1]
			filter[i] = alias
			m[alias] = m[field]
			delete(m, field)
		} else {
			filter[i] = field
		}
	}
	for key := range m {
		if !contains(key, filter) {
			delete(m, key)
		}
	}
	marshal, _ := json.Marshal(m)
	return string(marshal)
}

func contains(key string, filter []string) bool {
	for _, s := range filter {
		if key == s {
			return true
		}
	}
	return false
}

func format(jsonInput string, output ...string) string {
	get := gjson.GetMany(jsonInput, output...)

	sb := new(strings.Builder)

	colLen := len(get)

	lines := get[0].Array()

	for i := 0; i < len(lines); i++ {
		for j := 0; j < colLen; j++ {
			lineArray := get[j].Array()
			if cap(lineArray) == 0 {
				continue
			}
			tab, _ := strconv.Unquote(`"` + "\t" + `"`)
			sb.WriteString(lineArray[i].String() + tab)
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
