package main

import (
	"fmt"
	"log"
	"os"
	"textnumbers"
)

func main() {
	l, err := textnumbers.OpenLanguage("english")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(l.Title())

	if true {
		return
	}
	args := asArgs(os.Args[1:])
	p := args.Parameters()
	if len(p) != 1 {
		fmt.Println("Give a number as a parameter")
		return
	}
	//i, err := strconv.Atoi(p[0])
	//if err != nil {
	//	fmt.Printf("failed to read parameter as a  number  %v", err)
	//	return
	//}

}
