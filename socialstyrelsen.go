package main

import (
	"github.com/kracekumar/go-mwapi"
	"net/url"
	"io/ioutil"
	"log"
	"fmt"
	"encoding/json"
	"strings"
)

type GetCategoriesResult struct{
	BatchComplete	string `json:"batchcomplete"`
	Continue	ContinueToken `json:"continue"`
	QueryResult	QueryResult `json:"query"`
}

type ContinueToken struct{
	Token		string `json:"cmcontinue"`
	Continue	string `json:"continue"`
}

type QueryResult struct{
	CategoryMembers		[]CategoryPage `json:categorymebers`
}

type CategoryPage struct{
	PageID	int64 `json:pageid`
	Title	string `json:title`
}

func getCategories(wiktionaryApiUrl url.URL, category string, continueToken string) {
	params := url.Values{
		"action":	{"query"},
		"list":		{"categorymembers"},
		"cmtitle":	{category},
		"format":	{"json"},
		"cmlimit":	{"500"},
		"cmcontinue":	{continueToken},
	}

	api := mwapi.NewMWApi(wiktionaryApiUrl)
	res := api.Get(params)

	if res.StatusCode != 200 {
		log.Fatal()
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if (err != nil) {
		log.Fatal()
	}

	c := GetCategoriesResult{}
	err = json.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Fatal(err)
	}

	for _, word := range c.QueryResult.CategoryMembers {
		if !strings.HasPrefix(word.Title, "Kategori:") {
			fmt.Printf("%s\n", word.Title)
		}
	}

	if c.Continue.Token != "" {
		getCategories(wiktionaryApiUrl, "Kategori:Svenska/Substantiv", c.Continue.Token)
	}
}

func main() {
	wiktionaryApiUrl := url.URL{
		Scheme:	"https",
		Host:	"sv.wiktionary.org",
		Path:	"/w/api.php",
	}

	getCategories(wiktionaryApiUrl, "Kategori:Svenska/Substantiv", "")
}
