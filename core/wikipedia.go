package core

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
)

type WikipediaLanguage string

const (
	JapaneseInWikipedia WikipediaLanguage = "ja"
	EnglishInWikipedia  WikipediaLanguage = "en"
)

func getCorrespondingWikipediaLanguage(lang Language) WikipediaLanguage {
	switch lang {
	case Japanese:
		return JapaneseInWikipedia
	case English:
		return EnglishInWikipedia
	default:
		panic("unknown language: " + string(lang))
	}
}

/*
func getCorrespondingWikipediaLanguages(langs []Language) []WikipediaLanguage {
	var wikiLangs = make([]WikipediaLanguage, 0)

	for _, l := range langs {
		wikiLangs = append(wikiLangs, getCorrespondingWikipediaLanguage(l))
	}

	return wikiLangs
}
*/

func getBaseEndPoint(lang Language) string {
	var wikiLang = getCorrespondingWikipediaLanguage(lang)
	return "https://" + string(wikiLang) + ".wikipedia.org/w/api.php"
}

func createRandomQueryEndPoint(lang Language) string {
	var url = getBaseEndPoint(lang) +
		"?format=json" +
		"&action=query" +
		"&generator=random" +
		"&grnlimit=1" +
		"&grnnamespace=0" +
		"&prop=extracts" +
		"&exintro" +
		"&explaintext"

	return url
}

type PageInfo struct {
	PageId  int64  `json:"pageid"`
	Ns      int    `json:"ns"`
	Title   string `json:"title"`
	Summary string `json:"extract"`
}

func GetRandomArticleInfo(lang Language) (PageInfo, error) {
	var url = createRandomQueryEndPoint(lang)
	fmt.Println(url)

	type RandomRes struct {
		Query struct {
			Pages map[string]PageInfo `json:"pages"`
		} `json:"query"`
	}

	res, err := http.Get(url)
	if err != nil {
		return PageInfo{}, err
	}

	defer res.Body.Close()

	var randomRes RandomRes
	var buf = new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	var resByte = buf.Bytes()
	err = json.Unmarshal(resByte, &randomRes)
	if err != nil {
		return PageInfo{}, err
	}

	var ret PageInfo
	for _, page := range randomRes.Query.Pages {
		ret = page
		break
	}

	return ret, nil
}
