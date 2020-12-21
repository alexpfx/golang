package bookmark_test

import (
	"bmark/bookmark"
	"os"
	"testing"
)

var chromeBookmarkFilePath = "Bookmarks"

func TestMain(m *testing.M) {
	ex := m.Run()
	os.Exit(ex)
}

func TestReadFromChromeBookmarkFile1(t *testing.T) {
	bookmarker := bookmark.ReadFromChromeBookmarkFile("./data/Bookmarks.json")

	for _, bm := range bookmarker.All() {
		t.Log(bm.Url)
	}

}


func TestReadFromChromeBookmarkFile2(t *testing.T) {
	bookmarker := bookmark.ReadFromChromeBookmarkFile("./data/Bookmarks2.json")

	for _, bm := range bookmarker.All() {
		t.Log(bm.Url)
	}

}