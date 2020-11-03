package cmd

import (
	"bmark/bookmark"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"

	"fmt"

	"github.com/spf13/cobra"

	"log"
	"os"
)

func NewAddCmd() *cobra.Command {
	return addCmd
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona um bookmark",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if _, isFile := readString(cmd, "file"); isFile {
			log.Println(cmd.Usage())
			checkArgCount(args, 0, "file")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		bookmark.CheckCreateStorageFile()

		bookmarks := bookmark.ReadBookmarks()

		var desc, category string
		var tags []string
		url := args[0]

		clipboard := readBool(cmd, "clipboard")
		update := readBool(cmd, "update")
		extract := readBool(cmd, "extract")
		filename, file := readString(cmd, "file")

		if extract {
			var err error
			desc, category, tags, err = extractFromURL(url)
			checkPrint(err)

		} else if file {
			extractFromJsonFile(filename)
		} else if clipboard {
			url, desc, category, tags = readFromClipboard()
		} else {
			desc, category, tags = readFromFlags(cmd)
		}

		bookmarks, err := addBookmark(bookmarks, url, desc, category, tags, update)

		checkPrint(err)

		backupBookmarkFile()
		bookmark.StoreBookmarks(bookmarks)
	},
}

func extractFromJsonFile(fileName string) {


}

//TODO Separar
func extractFromURL(url string) (desc, category string, tags []string, err error) {
	if url == "" {
		err = fmt.Errorf("url não definida")
	}

	fmt.Println("extracting from url: ", url)

	desc = ""
	category = ""
	tags = []string{}
	err = nil

	client := getClient()

	r, err := client.Do(createRequest(url))
	checkPrint(err)

	defer r.Body.Close()

	htmlInfo := extractHtmlInfo(r.Body)

	desc = extractDesc(htmlInfo)
	category = "[CATEGORY]"
	tags = []string{"[TAG]"}

	return
}

func extractDesc(info *HtmlInfo) string {
	if info.Title != "" {
		return info.Title
	}

	desc := extractFirst(info.Meta, "title", "og:title", "twitter:title")

	return desc

}

func extractFirst(meta map[string]string, keys ...string) string {
	for _, key := range keys {
		f, ok := meta[key]
		if ok {
			return f
		}
	}
	return ""
}

type HtmlInfo struct {
	Title string
	Meta  map[string]string
}

func extractHtmlInfo(body io.Reader) *HtmlInfo {
	tkz := html.NewTokenizer(body)

	htmlInfo := new(HtmlInfo)
	htmlInfo.Meta = make(map[string]string, 0)

	isTitle := false
	for {
		tokenType := tkz.Next()

		if tokenType == html.ErrorToken {
			return htmlInfo
		}

		switch tokenType {

		case html.TextToken:
			if isTitle {
				token := tkz.Token()
				htmlInfo.Title = token.Data
				isTitle = false
				continue
			}

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tkz.Token()
			tokenData := token.Data
			if tokenData == `body` {
				return htmlInfo
			}

			switch tokenData {
			case "meta":
				key, value := extractMetaProp(token)
				if key != "" {
					htmlInfo.Meta[key] = value
				}
			case "title":
				isTitle = true
			}
		}
	}

}

func extractMetaProp(token html.Token) (key string, value string) {

	for _, att := range token.Attr {
		if att.Key == "property" {
			key = att.Val
		} else if att.Key == "content" {
			value = att.Val
		}
		if key != "" && value != "" {
			return
		}
	}
	return
}

func createRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func getClient() *http.Client {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Transport: tr, Timeout: time.Second + 3}
}

func readFromClipboard() (url, desc, category string, tags []string) {
	//todo()
	return "", "", "", []string{}
}

func addBookmark(bookmarks bookmark.Bookmarker, url string, desc string, category string, tags []string, updateIfExists bool) (collection bookmark.Bookmarker, err error) {
	item := bookmark.Item{
		Url:      url,
		Category: category,
		Desc:     desc,
		Tags:     tags,
	}

	_, found := bookmarks.Search(item)

	if found {
		if !updateIfExists {
			err = fmt.Errorf("bookmark já existe: %s\nuse -u para atualizar", item.Url)
			return
		}
		bookmarks.Update(item)
	} else {
		bookmarks.Add(item)
	}
	collection = bookmarks
	err = nil
	return
}

// TODO implement
func backupBookmarkFile() {

}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}
func checkExitOk(err error) {
	if err != nil {
		os.Exit(0)
	}
}

func checkArgCount(args []string, n int, argName string) {
	if len(args) != n {
		err := fmt.Errorf("número de argumentos inválido, esperava %d para o parâmetro %s", n, argName)
		log.Fatal(err)
	}
}
func checkPrint(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
func checkPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func readBool(cmd *cobra.Command, name string) bool {
	b, err := cmd.Flags().GetBool(name)
	checkPanic(err)
	return b
}

func readString(cmd *cobra.Command, name string) (str string, hasStr bool) {
	hasStr = false
	str, err := cmd.Flags().GetString(name)
	if err != nil || str == "" {
		return
	}
	hasStr = true
	return
}

func readFromFlags(cmd *cobra.Command) (desc, category string, tags []string) {
	desc, _ = cmd.Flags().GetString("desc")
	category, _ = cmd.Flags().GetString("category")
	tags = splitTags(cmd.Flags().GetString("tags"))

	return
}

func splitTags(tags string, err error) []string {
	if err != nil || tags == "" {
		return []string{}
	}

	tags = strings.ReplaceAll(tags, ",", " ")

	return strings.Split(tags, " ")
}

func init() {

	addCmd.Flags().BoolP("extract", "x", false, "-x")
	addCmd.Flags().StringP("file", "f", "", "-f <json_file_full_path>")

	addCmd.Flags().BoolP("clipboard", "b", false, "-b")
	addCmd.Flags().StringP("desc", "d", "", "--desc/-d <desc> ex: -d 'UOL'")
	addCmd.Flags().StringP("category", "c", "", "--category/-c <category> ex: -c 'portal'")
	addCmd.Flags().StringP("tags", "t", "", "--tags/-t [\"<tag>, <tag>\"] ex: -t 'noticias, esportes'")
	addCmd.Flags().BoolP("update", "u", false, "--update")

}
