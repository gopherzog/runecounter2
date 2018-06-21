package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const Version = "0.1.1"

type runeinfo map[string]int

func main() {
	fmt.Printf("runecounter2 %s\n", Version)
	err, rawtext := readText("proverbs.txt")
	if err != nil {
		problem := fmt.Sprintf("%v", err)
		panic(problem)
	}
	lines := strings.Split(rawtext, "\n")
	for idx, line := range lines {
		runemap := mapRunes(line)
		fmt.Printf("%2d) %s\n", idx, line)
		formattedMap := formatMap(runemap)
		fmt.Printf("%s\n\n", formattedMap)
	}
}

func mapRunes(line string) runeinfo {
	runes := make(map[string]int)
	for _, r := range line {
		s := string(r)
		runes[s] = runes[s] + 1
	}
	return runes
}

func formatMap(info runeinfo) string {
	var allItems []string
	for k, v := range info {
		item := fmt.Sprintf("'%s'=%d", k, v)
		allItems = append(allItems, item)
	}
	textRepr := strings.Join(allItems, ", ")
	return textRepr
}

func readText(filename string) (error, string) {
	dat, err := ioutil.ReadFile(filename)
	proverbs := ""
	if err == nil {
		proverbs = string(dat)
	}
	return err, proverbs
}
