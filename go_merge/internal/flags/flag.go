package flags

import "flag"

type BoolSlice map[string]bool

func (v BoolSlice) String() string {
	return "true"
}

func (v *BoolSlice) Set(id string) error {
	return nil

}

func main() {
	var bs BoolSlice

	flag.Var(&bs, "te", "t")


}

