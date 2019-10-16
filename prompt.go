package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
	"time"
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
		"7/8 (M)",
		"6/8 (M)",
		"5/8 (M)",
		"4/8 (M)",
		"3/8 (M)",
		"2/8 (M)",
		"1/8 (M)",
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

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	guilds, err := queryGuilds(realm, class, int(progressValue))
	s.Stop()

	createTable(guilds)
}

func createTable(guilds []GuildInfo) {
	var data [][]string

	for i := 0; i < len(guilds); i++ {
		data = append(data, []string{guilds[i].GuildData.Name, guilds[i].Progress, guilds[i].GuildData.Url})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Guild Name", "Progress", "WoWProgress Link"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
