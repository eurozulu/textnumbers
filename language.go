package textnumbers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const languageFolderName = "languages"

type Language interface {
	Title() string
	Format(i uint64) string
	MinusLabel() string
}

type language struct {
	title        string
	names        []*valueName
	separators   []*valueName
	invertDigits bool
	minusLabel   string
}

func (l language) Title() string {
	return l.title
}

func (l language) MinusLabel() string {
	return l.minusLabel
}

func (l language) Format(i uint64) string {
	var digits []string
	for {
		vn := l.valueNameFor(i)
		vb := vn.ValueBase()

		var vs string // formatted string of value
		var vf uint64 // value that's been formatted
		var sp *valueName
		if vn.IsLabel() {
			// Format the value at the base defined by valueName (12345 @base 3 = 12, @base 1 = 45)
			d := Number(i).DigitsAt(vb)
			vs = fmt.Sprintf(vn.String(), l.Format(d))
			// Formatted only the digits of the value-name base or above.
			vf = d * uint64(math.Pow10(int(vb)))
			sp = l.seperatorFor(i - vf)
		} else {
			// Not a label, maps directly to the name
			vs = vn.String()
			// formatted entire value
			vf = vn.Value
		}

		digits = append(digits, vs)
		// insert seperator if needed
		if sp != nil {
			digits = append(digits, sp.Name)
		}

		i = (i - vf)
		if i == 0 {
			// Last base been processed
			break
		}

	}
	return strings.Join(digits, " ")
}

func (l language) valueNameFor(i uint64) *valueName {
	name := l.names[0]
	for _, lb := range l.names[1:] {
		if lb.Value > i {
			break
		}
		name = lb
	}
	return name
}

func (l language) seperatorFor(v uint64) *valueName {
	if v == 0 {
		return nil
	}
	for _, dl := range l.separators {
		if dl.Value >= v {
			return dl
		}
	}
	return nil
}

func (l *language) UnmarshalJSON(bytes []byte) error {
	// using standard json decoding, into a psudo instance, then sorts bases before assigning to this.
	var lp struct {
		Title        string       `json:"title"`
		Names        []*valueName `json:"names"`
		Separator    []*valueName `json:"separator,omitempty"`
		InvertDigits bool         `json:"invert-digits,omitempty"`
		MinusLabel   string       `json:"minus"`
	}
	if err := json.Unmarshal(bytes, &lp); err != nil {
		return err
	}
	// sort names lowest value first
	sort.Slice(lp.Names, func(i, j int) bool {
		return lp.Names[i].Value < lp.Names[j].Value
	})
	// sort delimiters lowest value first
	sort.Slice(lp.Separator, func(i, j int) bool {
		return lp.Separator[i].Value < lp.Separator[j].Value
	})

	l.title = lp.Title
	l.names = lp.Names
	l.separators = lp.Separator
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
	return nil
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
