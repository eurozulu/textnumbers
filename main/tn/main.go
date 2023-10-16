package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"textnumbers"
)

const defaultLanguage = "english"

func main() {
	n, lang, err := readArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := textnumbers.OpenLanguage(lang)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(l.Format(n))
}

func readArgs(args []string) (int64, string, error) {
	if len(args) < 1 {
		return 0, "", fmt.Errorf("provide a number to convert")
	}
	i, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return 0, "", err
	}
	var lang string
	ix := findIndex("as", args)
	if ix > 0 {
		if ix+1 >= len(args) {
			return 0, "", fmt.Errorf(("must provide a language name."))
		}
		lang = args[ix+1]
	} else {
		lang = defaultLanguage
	}
	return i, lang, nil
}

func findIndex(s string, ss []string) int {
	for i, sz := range ss {
		if strings.EqualFold(sz, s) {
			return i
		}
	}
	return -1
}
