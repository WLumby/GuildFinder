package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

func main() {

	classes := []string{
		"deathknight",
		"demon",
		"druid",
		"hunter",
		"mage",
		"monk",
		"paladin",
		"priest",
		"rogue",
		"shaman",
		"warlock",
		"warrior",
	}

	prompt := promptui.Select{
		Label: "Select Class",
		Items: classes,
	}

	_, class, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	promptP := promptui.Prompt{
		Label: "Enter Realm (e.g. eu_draenor)",
	}

	realm, err := promptP.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	progress := []string{
		"8/8 (M)",
	}

	prompt = promptui.Select{
		Label: "Select Minimum Progress",
		Items: progress,
	}

	_, progressFound, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	progressValue, err := strconv.ParseInt(progressFound[:strings.IndexByte(progressFound, '/')], 10, 64)
	if err != nil {
		return
	}

	guilds, err := queryGuilds(realm, class, int(progressValue))

	fmt.Println(guilds)
}

