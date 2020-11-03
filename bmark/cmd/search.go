package cmd

import (
	"bmark/bookmark"
	"fmt"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"runtime"
)

func NewSearchCmd() *cobra.Command {
	return searchCmd(func(term string, output string) {
		fmt.Println(output)
	})
}

func SearchCmdWithOutputListener(listener func(string, string)) *cobra.Command {
	sc := searchCmd(listener)
	return sc

}

func searchCmd(out func(string, string)) *cobra.Command {
	use := "search"
	brief := "Abre a busca por bookmark"

	args := func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("número de argumentos inválidos")
		}

		return nil
	}

	run := func(cmd *cobra.Command, args []string) {
		bookmarker := bookmark.ReadBookmarks()
		items := bookmarker.All()
		all := items
		if len(all) == 0 {
			return
		}

		option := withPreview(items)

		selected, err := fuzzyfinder.Find(items, func(index int) string {
			selected := items[index]
			desc := selected.Desc

			return fmt.Sprintf("%v", desc)

		}, option,
		)
		check(err, "nenhum item selecionado")

		openBrowser(items[selected])

	}

	var cm = cobra.Command{
		Use:   use,
		Short: brief,
		Run:   run,
		Args:  args,
	}

	return &cm

}

func openBrowser(item bookmark.Item) {
	url := item.Url
	if url == "" {
		return
	}

	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func withPreview(items []bookmark.Item) fuzzyfinder.Option {
	preview := func(index, w, h int) string {
		if index == -1 {
			return ""
		}

		selected := items[index]
		url := selected.Url
		desc := selected.Desc
		tags := selected.Tags
		category := selected.Category

		return fmt.Sprintf("Selecionado:\n%v\n%v\n%v\n%v", desc, category, url, tags)

	}

	return fuzzyfinder.WithPreviewWindow(preview)
}
