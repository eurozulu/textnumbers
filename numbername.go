package textnumbers

// numberName represents a direct mapping of a single value to a Name
type numberName struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

func (v numberName) String() string {
	return v.Name
}

// numberLabel represents a string label which is applied to values within the same Base (number of digits)
type numberLabel struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

func (v numberLabel) String() string {
	return v.Label
}
