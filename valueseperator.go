package textnumbers

type valueSeperator struct {
	Name          string `json:"name"`
	Value         uint64 `json:"value"`
	ReverseDigits bool   `json:"reverse"`
}

func (v valueSeperator) String() string {
	return v.Name
}
