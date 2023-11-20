package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"textnumbers"
)

const defaultLanguage = "english"

func main() {
	args, err := readArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := textnumbers.OpenLanguage(args.language)
	if err != nil {
		log.Println(err)
		return
	}
	v := l.Format(args.value)
	if args.isMinus {
		v = strings.Join([]string{l.MinusLabel(), v}, " ")
	}

	if !args.quiet {
		fmt.Printf("%s in %s:\n%s\n", os.Args[1], l.Title(), v)
	} else {
		fmt.Println(v)
	}

}

func readArgs(args []string) (*myargs, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("provide a number to convert")
	}

	var found myargs
	number := strings.Join(args, "")
	found.isMinus = strings.HasPrefix(number, "-")
	if found.isMinus {
		number = strings.TrimLeft(number, "-")
	}
	if strings.Contains(number, ",") {
		number = strings.Replace(number, ",", "", -1)
	}
	i, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		if strings.HasSuffix(err.Error(), strconv.ErrRange.Error()) {
			return nil, fmt.Errorf("The number %s is too big to parse. Maximum value is: %v", number, uint64(math.MaxUint64))
		}
		return nil, err
	}
	found.value = i
	ix := findIndex("as", args)
	if ix < 0 {
		ix = findIndex("in", args)
	}
	if ix > 0 {
		if ix+1 >= len(args) {
			return nil, fmt.Errorf(("must provide a language name."))
		}
		found.language = args[ix+1]
	} else {
		found.language = defaultLanguage
	}
	ix = findIndex("-q", args)
	found.quiet = ix >= 0

	return &found, nil
}

type myargs struct {
	value    uint64
	isMinus  bool
	language string
	quiet    bool
}

func findIndex(s string, ss []string) int {
	for i, sz := range ss {
		if strings.EqualFold(sz, s) {
			return i
		}
	}
	return -1
}
