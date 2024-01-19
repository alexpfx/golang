package util

import (
	"encoding/json"
	"sort"
)

func FormatJson(results interface{}) (string, error) {
	bytes, err := json.MarshalIndent(results, "", "   ")

	return string(bytes), err
}

func IntUniqueSorted(ids []int) []int {
	sort.Ints(ids)
	j := 0
	for i := 0; i < len(ids); i++ {
		if ids[i] == ids[j] {
			continue
		}
		j++
		ids[j] = ids[i]
	}
	return ids[:j+1]

}
