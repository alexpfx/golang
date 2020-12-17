package bookmark

import "regexp"

// BookmarkHolder representa um conjunto de bookmarks e operação que podem ser realizadas sobre eles
type BookmarkHolder interface {
	Update(item Item)
	Add(item Item)
	Search(item Item) (i *Item, found bool)
	All() []Item
	AllUrls() []string
}

// BookmarkCollection armazena uma lista de bookmarks
type BookmarkCollection struct {
	Bookmarks []Item `toml:"bookmark" json:"bookmark"`
}

type ChromeCollection struct {
	Roots Roots `json:"roots"`
}

type Roots struct {
	BookmarkBar ChromeItem `json:"bookmark_bar"`
	Other       ChromeItem `json:"other"`
	Synced      ChromeItem `json:"synced"`
}

type ChromeItem struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Id   string `json:"id"`
	Url  string `json:"url"`

	Children []ChromeItem `json:"children"`
}

// Update atualiza um bookmark
func (c *BookmarkCollection) Update(item Item) {
	p, found := c.Search(item)
	if found {
		(*p).update(item)
	}

}
func (i *Item) update(item Item) {
	i.Desc = item.Desc
	i.Tags = item.Tags
	i.Category = item.Category

}

//Add adiciona um bookmark a collection
func (c *BookmarkCollection) Add(item Item) {
	c.Bookmarks = append(c.Bookmarks, item)
}

// All retorna todos os Bookmarks
func (c BookmarkCollection) All() []Item {
	return c.Bookmarks
}

var reg, err = regexp.Compile("[^a-zA-Z0-9]+")

// AllUrls retorna um slice com todas as URLs
func (c BookmarkCollection) AllUrls() []string {
	urls := make([]string, 0)

	for _, it := range c.Bookmarks {
		url := it.Url
		if url == "" {
			continue
		}
		//withoutSpecialChars := reg.ReplaceAllString(url, " ")
		urls = append(urls, url)
	}
	return urls
}

//Search busca um bookmark pelo url
func (c BookmarkCollection) Search(item Item) (*Item, bool) {
	for i, it := range c.Bookmarks {
		if it.Url == item.Url {
			return &c.Bookmarks[i], true
		}
	}

	return nil, false
}

type Item struct {
	Url      string   `toml:"url" json:"url"`
	Desc     string   `toml:"desc" json:"desc"`
	Tags     []string `toml:"tags" json:"tags"`
	Category string   `toml:"category" json:"category"`
}
