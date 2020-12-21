package cmd_test_test

import (
	"bmark/bookmark"
	"bmark/cmd"
	"bmark/cmd_test"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
)

const testFailMessage = "para o termo %v, esperava %v, recebeu %v.\n"

func TestMain(m *testing.M) {
	_ = os.Setenv("BMARK_CONFIG_DIR", "./data")
	ex := m.Run()
	os.Exit(ex)
}

func TestSearchCmd(t *testing.T) {
	var table = map[string]string{
		"volut":        "http://51.la/nec/nisi/volutpat/eleifend/donec.jsp",
		"big":          "http://bigcartel.com/at/nunc/commodo.html",
		"quis":         "https://posterous.com/donec/quis.aspx",
		"png":          "https://people.com.cn/ultrices/vel/augue/vestibulum.png",
		"eget":         "https://phpbb.com/eget/tincidunt/eget.aspx",
		"http":         "https://redcross.org/donec/odio/justo/sollicitudin/ut/suscipit/a.json",
		"ba8":          "http://ba8.net/ante/vel/ipsum.aspx",
		"iaculis":      "http://tiny.cc/iaculis/justo/in/hac/habitasse.png",
		"mac":          "https://mac.com/tempus.json",
		"wisc":         "https://wisc.edu/pede/venenatis/non/sodales/sed/tincidunt/eu.aspx",
		"tortor":       "http://moonfruit.com/accumsan/tortor.jsp",
		"cisco":        "https://cisco.com/etiam/vel/augue.aspx",
		"163":          "http://163.com/iaculis/diam/erat/fermentum/justo/nec.js",
		"de.vu":        "https://de.vu/id/ornare/imperdiet/sapien.aspx",
		"freewebs.com": "http://freewebs.com/non/mi.jpg",
		"state":        "https://state.gov/mauris/vulputate/elementum.jsp",
		"diam":         "http://bbb.org/at/diam/nam.png",
		"nisl":         "http://arizona.edu/nisl/ut/volutpat/sapien.png",
		"ante":         "https://disqus.com/at/lorem/integer/tincidunt/ante.png",
		"usatoday":     "https://usatoday.com/rhoncus/sed/vestibulum/sit/amet/cursus/id.xml",
		"dyndns":       "https://dyndns.org/in/purus/eu/magna.aspx",
		"nps.gov":      "http://nps.gov/orci/pede/venenatis/non/sodales/sed.xml",
		"dropbox":      "http://dropbox.com/consectetuer/adipiscing.js",
		"epa":          "https://epa.gov/morbi/ut/odio/cras.json",
		"ox.ac":        "https://ox.ac.uk/ligula/vehicula/consequat.png",
		"homestead":    "http://homestead.com/mi/in/porttitor/pede/justo/eu.aspx",
		"ls.com":       "https://parallels.com/ante/vivamus/tortor/duis/mattis.html",
		"nibh":         "http://dedecms.com/nibh/ligula/nec/sem.js",
		"justo":        "http://gmpg.org/id/justo/sit/amet/sapien/dignissim.png",
		"sfgate":       "http://sfgate.com/velit/nec/nisi/vulputate/nonummy/maecenas.json",
		"icq.com":      "http://icq.com/mi/integer/ac/neque/duis/bibendum.js",
		"free.fr":      "http://free.fr/commodo.js",
		"cursus":       "https://prweb.com/vivamus/in/felis/eu/sapien/cursus.xml",
		"click":        "https://clickbank.net/ipsum.xml",
		"wikia":        "https://wikia.com/ipsum/praesent.xml",
		"flavors":      "https://flavors.me/at/velit/vivamus/vel.xml",
		"goo.ne":       "http://goo.ne.jp/aliquam/sit/amet/diam.json",
		"salon":        "https://salon.com/quis.jpg",
		"blogs":        "http://blogs.com/nam/dui/proin/leo/odio.png",
		"sitemeter":    "http://sitemeter.com/sed.js",
		"nbcnews":      "http://nbcnews.com/interdum/mauris/non/ligula/pellentesque/ultrices.aspx",
		"fermentum":    "https://goodreads.com/iaculis/diam/erat/fermentum/justo/nec.jpg",
		"https":        "https://hubpages.com/tortor/sollicitudin/mi/sit/amet/lobortis/sapien.json",
		"tellus":       "https://parallels.com/at/ipsum/ac/tellus.xml",
	}
	var nok int

	funcao := func(term string, output string) {
		if len(output) == 0 {
			//fmt.Printf("term: %v output: %v\n", term, output)
			nok++
			fmt.Printf("não encontrou: %v: %v\n", term, nok)

		}

		/*if table[term] != output{
			fmt.Printf(testFailMessage, term, table[term], output)
			er ++
		}else{
			fmt.Printf("ok: %v %v\n", term, output)
			ok ++
		}
		println("*********************** ok: ", ok)
		println("*********************** er: ", er)*/
	}

	cmd.Init([]*cobra.Command{cmd.SearchCmdWithOutputListener(funcao)})

	for key := range table {
		os.Args = []string{
			"bmark", "search", key,
		}

		cmd.Execute()
		fmt.Println("")
		break

	}
}

func tTestGenData(t *testing.T) {
	_ = os.Setenv("BMARK_CONFIG_DIR", "./data")

	customConfigDir := bookmark.GetConfigDir("./data")
	bytes, err := ioutil.ReadFile(customConfigDir + string(os.PathSeparator) + "MOCK_DATA.json")
	check(err)

	var mock []cmd_test.MockData
	err = json.Unmarshal(bytes, &mock)
	check(err)

	lim := 80
	count := 0

	for i, data := range mock {

		if rand.Intn(20) == 1 {
			log.SetOutput(os.Stdout)
			sample(count, lim, data)
			log.SetOutput(ioutil.Discard)
		}

		if i%100 == 0 {
			//log.SetOutput(os.Stdout)
			log.SetOutput(ioutil.Discard)
			t.Log("********************************************************")
			t.Log("************************** ", i, "**********************")
			t.Log("********************************************************")
		}

		args := []string{"bmark",
			"add",
			"-d", data.Desc,
			"-c", data.Category,
			"-t", splitJoin(data.Tags),
			"-u",
			data.Url}

		os.Args = args
		cmd.Execute()

	}

}
func sample(m, n int, data cmd_test.MockData) {
	fmt.Printf(`{"", "%s"}, %s`, data.Url, "\n")
}
func splitJoin(tags string) string {
	//return strings.Join(strings.Split(tags, " "), ",")
	return tags
}

func check(err error) {
	if err != nil {
		panic("erro " + err.Error())
	}
}
