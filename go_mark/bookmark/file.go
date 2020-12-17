package bookmark

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configDirName = "/bmark"
const bookmarkFilename = "/bookmarks.json"
const envVarCustomConfigDir = "BMARK_CONFIG_DIR"

const configFileName = "/conf.toml"
const bookmarkBkpFilename = "/bookmarks.toml.bk"

// LoadBookmarks lê os bookmarks do arquivo
func LoadBookmarks(storageFile string) BookmarkHolder {
	//bookmarkFilePath := StorageFileFullPath(os.Getenv(envVarCustomConfigDir))

	//log.Println("bookmarkFilePath ", bookmarkFilePath)

	bytes, err := ioutil.ReadFile(storageFile)
	checkErr(err)

	var collection Collection

	//err = tom.Unmarshal(bytes, &collection)

	err = json.Unmarshal(bytes, &collection)
	checkErr(err)

	return &collection
}

func ReadFromChromeBookmarkFile(filePath string) (bookmarks BookmarkHolder) {
	fileBytes, err := ioutil.ReadFile(filePath)
	checkErr(err)

	jMap := ChromeCollection{}

	err = json.Unmarshal(fileBytes, &jMap)
	checkErr(err)

	bookmarks = extractFromChromeMap(jMap)

	return
}

func extractFromChromeMap(collection ChromeCollection) BookmarkHolder {
	bookmarker := Collection{
		[]Item{},
	}

	children := traverse(collection.Roots.BookmarkBar.Children, []string{})

	for _, item := range children {
		bookmarker.Add(item)
	}

	return &bookmarker
}

func traverse(children []ChromeItem, parents []string) (items []Item) {
	for _, ch := range children {
		if ch.Type == "url" {
			var category string

			if len(parents) > 0 {
				category = parents[0]
			} else {
				category = ""
			}

			items = append(items, Item{
				Url:      ch.Url,
				Desc:     ch.Name,
				Tags:     parents,
				Category: category,
			})
		} else if ch.Type == "folder" {

			items = append(items, traverse(ch.Children, append(parents, ch.Name))...)
		}
	}
	return
}

//StoreBookmarks é responsavel por armazenar os bookmarks em arquivo
func StoreBookmarks(bookmarks BookmarkHolder) {
	bytes, err := json.MarshalIndent(bookmarks, " ", "   ")
	checkErr(err)

	fullPath := StorageFileFullPath(os.Getenv(envVarCustomConfigDir))

	err = ioutil.WriteFile(fullPath, bytes, 0777)
	checkErr(err)

	//log.Printf("dados gravados. arquivo contém agora %d bytes", len(bytes))
}

// CheckCreateStorageFile verifica a existencia do arquivo de armazenamento,
// caso não existe o arquivo é criado.
func CheckCreateStorageFile() {
	customConfigDir := os.Getenv(envVarCustomConfigDir)
	bookmarkFilePath := StorageFileFullPath(customConfigDir)
	_, err := os.Stat(bookmarkFilePath)

	if os.IsNotExist(err) {
		//log.Printf("arquivo %s não existe, criando", bookmarkFilePath)
		configDir := GetConfigDir(customConfigDir)
		err = os.MkdirAll(configDir, 0755)
		check(err, fmt.Sprintf("não pode criar arquivo de bookmars em %s. forneça o diretório de configuração através da variável de ambiente %s", configDir, envVarCustomConfigDir))

		f, err := os.Create(bookmarkFilePath)
		checkErr(err)
		defer f.Close()
	} else {
		//log.Println("arquivo existe: ", bookmarkFilePath)
	}

}

// StorageFileFullPath monta o caminho do arquivo de armazenamento.
func StorageFileFullPath(customConfigDir string) string {
	bookmarkFilePath := GetConfigDir(customConfigDir) + bookmarkFilename
	return bookmarkFilePath
}

// GetConfigDir monta o caminho do diretório de configuração da aplicação
func GetConfigDir(customConfigDir string) string {
	var configDir string

	if customConfigDir != "" {
		configDir = customConfigDir
	} else {
		userConfigDir, err := os.UserConfigDir()
		checkErr(err)
		configDir = userConfigDir + configDirName
	}
	//log.Println("Config dir: ", configDir)

	return configDir
}
