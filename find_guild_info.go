package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

var bow = surf.NewBrowser()

func main() {
	jsonData, err := getRealmData("eu_draenor")
	if err != nil {
		panic(err)
	}
	realm, err := unmarshalRealm(jsonData)

	fmt.Println(realm)
}

func findGuildInfo(address, class string) (string, error) {

	err := bow.Open(address)
	if err != nil {
		return "", err
	}

	findString := "span[class]." + class
	bow.Dom().Find(findString).Each(func(_ int, s *goquery.Selection) {
		if s.Text() == "priest" {
			// Do something
		}
	})

	return "", nil
}
