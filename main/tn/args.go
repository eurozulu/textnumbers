package main

import "strings"

type asArgs []string

func (a asArgs) Parameters() []string {
	i := a.asIndex()
	if i < 0 {
		i = len(a)
	}
	return a[:i]
}

func (a asArgs) Language() string {
	i := a.asIndex()
	if i < 0 || i+1 >= len(a) {
		return ""
	}
	return a[i+1]
}

func (a asArgs) asIndex() int {
	for i, arg := range a {
		if strings.EqualFold(arg, "as") {
			return i
		}
	}
	return -1
}
