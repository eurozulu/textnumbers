package textnumbers

// valueName represents a mapping of a value to a string
type valueName struct {
	Name  string `json:"name"`
	Value uint64 `json:"value"`
}

func (v valueName) String() string {
	return v.Name
}
