package textnumbers

import "strings"

const ValueToken = "%v"

// valueName represents a mapping of a value to a string
type valueName struct {
	Name  string `json:"name"`
	Value uint64 `json:"value"`
}

func (v valueName) String() string {
	return v.Name
}

func (v valueName) IsLabel() bool {
	return strings.Contains(v.Name, ValueToken)
}

func (v valueName) ValueBase() Base {
	return Number(v.Value).Base()
}
