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
	if args.isMinus {
		fmt.Printf("%s ", l.MinusLabel())
	}
	fmt.Println(l.Format(args.value))
}

func readArgs(args []string) (*myargs, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("provide a number to convert")
	}

	var found myargs
	found.isMinus = strings.HasPrefix(args[0], "-")
	if found.isMinus {
		args[0] = strings.TrimLeft(args[0], "-")
	}
	i, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		if strings.HasSuffix(err.Error(), strconv.ErrRange.Error()) {
			return nil, fmt.Errorf("The number %s is too big to parse. Maximum value is: %v", args[0], uint64(math.MaxUint64))
		}
		return nil, err
	}
	found.value = i
	ix := findIndex("as", args)
	if ix > 0 {
		if ix+1 >= len(args) {
			return nil, fmt.Errorf(("must provide a language name."))
		}
		found.language = args[ix+1]
	} else {
		found.language = defaultLanguage
	}
	return &found, nil
}

type myargs struct {
	value    uint64
	isMinus  bool
	language string
}

func findIndex(s string, ss []string) int {
	for i, sz := range ss {
		if strings.EqualFold(sz, s) {
			return i
		}
	}
	return -1
}
