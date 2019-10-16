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

	for i:=0;i<len(realm);i++{
		fmt.Println(realm[i].Name)
		_, _ = findGuildInfo(realm[i].Url, "priest")
	}
}

func findGuildInfo(address, class string) (string, error) {

	err := bow.Open(address)
	if err != nil {
		return "", err
	}

	findString := "span[class]." + class
	bow.Dom().Find(findString).Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

	findString = "span[class].innerLink"
	bow.Dom().Find(findString).Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

	return "", nil
}
