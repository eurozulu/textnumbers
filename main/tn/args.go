package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// myargs parses the standard os.Args slice into values
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

func newArgs(args []string) (*myargs, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("provide a number to convert")
	}

	var found myargs
	ix := findIndex("as", args)
	if ix < 0 {
		ix = findIndex("in", args)
	}
	if ix > 0 {
		if ix+1 >= len(args) {
			return nil, fmt.Errorf(("must provide a language name."))
		}
		found.language = args[ix+1]
		args = args[:ix]
	} else {
		found.language = defaultLanguage
	}
	ix = findIndex("-q", args)
	found.quiet = ix >= 0
	if found.quiet {
		args = args[:ix]
	}

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
	return &found, nil
}
