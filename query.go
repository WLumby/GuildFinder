package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"strconv"
	"strings"
)

type GuildInfo struct {
	GuildData Guild
	Progress  string
}

var (
	bow = surf.NewBrowser()
	guildList []GuildInfo
)

func queryGuilds(realm, class string, progress int) ([]GuildInfo, error) {
	jsonData, err := getRealmData(realm)
	if err != nil {
		panic(err)
	}
	guilds, err := unmarshalRealm(jsonData)

	for i := 0; i < len(guilds); i++ {
		guildProgress, err := findGuildProgress(guilds[i].Url)
		if err != nil {
			return guildList, err
		}

		progressValue, err := strconv.ParseInt(guildProgress[:strings.IndexByte(guildProgress, '/')], 10, 64)
		if err != nil {
			return guildList, err
		}

		classWanted, err := findClassWanted(guilds[i].Url, class)
		if err != nil {
			return guildList, err
		}

		if int(progressValue) < progress {
			return guildList, nil
		}

		if classWanted  {
			guildList = append(guildList, GuildInfo{GuildData: guilds[i], Progress:  guildProgress,})
		}
	}

	return guildList, nil
}

func findClassWanted(address, class string) (bool, error) {
	foundClass := false

	err := bow.Open(address)
	if err != nil {
		return false, err
	}

	findString := "span[class]." + class
	bow.Dom().Find(findString).Each(func(_ int, s *goquery.Selection) {
		if s.Text() == class {
			foundClass = true
		}
	})

	return foundClass, nil
}

func findGuildProgress(address string) (string, error) {
	var foundProgress string

	err := bow.Open(address)
	if err != nil {
		return "", err
	}

	findString := "span[class].innerLink"
	bow.Dom().Find(findString).Each(func(_ int, s *goquery.Selection) {
		foundProgress = s.Text()
	})

	return foundProgress, nil
}
