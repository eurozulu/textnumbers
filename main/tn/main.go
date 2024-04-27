package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"textnumbers"
)

const defaultLanguage = "english"

func main() {
	args, err := newArgs(os.Args[1:])
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
