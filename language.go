package textnumbers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const languageFolderName = "languages"

type Language interface {
	Title() string
	Format(i int64) string
}

type language struct {
	title        string
	names        []*numberName
	labels       []*numberLabel
	separator    string
	invertDigits bool
	minusLabel   string
}

func (l language) Title() string {
	return l.title
}

func (l language) Format(i int64) string {
	var minus string
	isMinus := i < 0
	if isMinus {
		i = -i
		minus = l.minusLabel
	}
	// trim off the higher digits which have labels
	lb, r := l.writeLabels(i)
	var n string
	// write digits only if it's > 0 unless requesting zero.
	if i == 0 || r > 0 {
		n = l.writeName(r)
	}
	var sp string
	if len(lb) > 0 && len(n) > 0 {
		sp = l.separator
	}

	return NewStringBuffer().Append(minus).Append(lb).Append(sp).Append(n).String()
}

func (l language) writeName(i int64) string {
	var names []string
	for {
		nm := l.nameFor(i)
		n := nm.Name
		i = i - nm.Value
		names = append(names, n)
		if i == 0 {
			break
		}
	}
	if l.invertDigits {
		names = l.invertedLastToFirst(names)
	}
	return strings.Join(names, " ")
}

func (l language) writeLabels(i int64) (string, int64) {
	sb := NewStringBuffer()
	lb := l.labelFor(i)
	for ; lb != nil; lb = l.labelFor(i) {
		lv := valueDigits(lb.Value, i)
		// recursive call to get the value string to label
		vs := l.Format(lv)
		sb.Append(fmt.Sprintf("%s %s", vs, lb.Label))
		i = valueRemain(lb.Value, i)
	}
	return sb.String(), i
}

func (l language) nameFor(i int64) *numberName {
	name := l.names[0]
	for _, lb := range l.names[1:] {
		if lb.Value > i {
			break
		}
		name = lb
	}
	return name
}

func (l language) labelFor(i int64) *numberLabel {
	var label *numberLabel
	for _, lb := range l.labels {
		if lb.Value > i {
			break
		}
		label = lb
	}
	return label
}

func (l language) invertedLastToFirst(s []string) []string {
	if len(s) < 2 {
		return s
	}
	last := len(s) - 1
	return append([]string{s[last], l.separator}, s[:last]...)
}

func (l *language) UnmarshalJSON(bytes []byte) error {
	// using standard json decoding, into a psudo instance, then sorts bases before assigning to this.
	var lp struct {
		Title        string         `json:"title"`
		Names        []*numberName  `json:"names"`
		Labels       []*numberLabel `json:"labels,omitempty"`
		Separator    string         `json:"separator,omitempty"`
		InvertDigits bool           `json:"invert-digits,omitempty"`
		MinusLabel   string         `json:"minus"`
	}
	if err := json.Unmarshal(bytes, &lp); err != nil {
		return err
	}
	// sort names lowest value first
	sort.Slice(lp.Names, func(i, j int) bool {
		return lp.Names[i].Value < lp.Names[j].Value
	})

	// sort labels lowest value first
	sort.Slice(lp.Labels, func(i, j int) bool {
		return lp.Labels[i].Value < lp.Labels[j].Value
	})

	l.title = lp.Title
	l.names = lp.Names
	l.labels = lp.Labels
	l.separator = lp.Separator
	l.invertDigits = lp.InvertDigits
	l.minusLabel = lp.MinusLabel
	return l.validate()
}

func (l *language) validate() error {
	if len(l.names) == 0 {
		return fmt.Errorf("no names found")
	}
	if l.names[0].Value != 0 {
		return fmt.Errorf("no zero value title found")
	}

	// ensure all labels are higher than names
	if len(l.labels) > 0 {
		if l.names[len(l.names)-1].Value > l.labels[0].Value {
			return fmt.Errorf("label values must be higher than all title values")
		}
	}
	m := map[int64]bool{}
	for _, n := range l.names {
		if m[n.Value] {
			return fmt.Errorf("title and label values must be unique, %d already exists", n.Value)
		}
		m[n.Value] = true
	}
	for _, lb := range l.labels {
		if m[lb.Value] {
			return fmt.Errorf("title and label values must be unique, %d already exists", lb.Value)
		}
		m[lb.Value] = true
	}
	return nil
}

func valueDigits(value, i int64) int64 {
	base := Base(Number(value).DigitCount() - 1)
	return Number(i).DigitsAt(base)
}

func valueRemain(value, i int64) int64 {
	b := Base(Number(value).DigitCount() - 1)
	if b == 0 {
		return 0
	}
	ii := Number(i).ValueAt(b - 1)
	return ii
}

func cleanName(name string) string {
	if filepath.Ext(name) == "" {
		name = strings.Join([]string{name, "json"}, ".")
	}
	if filepath.Dir(name) != languageFolderName {
		name = filepath.Join(languageFolderName, name)
	}
	return name
}

func OpenLanguage(name string) (Language, error) {
	f, err := os.Open(cleanName(name))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	var l language
	if err = json.NewDecoder(f).Decode(&l); err != nil {
		return nil, err
	}
	return l, nil
}
