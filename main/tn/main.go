package main

import (
	"fmt"
	"github.com/eurozulu/textnumbers"
	"log"
	"os"
	"strings"
)

func main() {
	args, err := newArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	if args.language == "" {
		args.language = textnumbers.DefaultLanguage()
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
