package main

import (
	//"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const Version = "0.1.3"

type runeinfo map[string]int

func filenameFromArgs(args []string) (error, string) {
	filename := os.Getenv("FILE")
	maybePath := flag.String("f", filename, "file to be processed")
	flag.Parse()
	if *maybePath != "" {
		filename = *maybePath
	}
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return err, ""
		}
	}
	return nil, filename
}

func main() {
	fmt.Printf("runecounter2 %s\n", Version)
	err, filename := filenameFromArgs(os.Args)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	err, rawtext := readText(filename)
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
	lx := len(dat)
	proverbs := ""
	if err == nil {
		maxChars := lx
		if string(dat[lx-1]) == "\n" {
			maxChars--
		}
		proverbs = string(dat[:maxChars])
	}
	return err, proverbs
}
