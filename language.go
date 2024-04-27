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

// Language represents a number formater for single Language
type Language interface {
	// Title is the name of the Language
	Title() string

	// Format formats the given number into the language
	Format(i uint64) string

	// MinusLabel is the string used to indicate a negative number in the language
	MinusLabel() string
}

type language struct {
	title      string
	minusLabel string
	digitSpace string
	names      []*valueName
	separators []*valueSeperator
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
		nb := vn.ValueBase()

		var ns string // formatted string of number
		var nf uint64 // number that's been formatted
		if vn.IsLabel() {
			// Format the value at the base defined by valueName (12345 @base 3 = 12, @base 1 = 45)
			d := Number(i).DigitsAt(nb)
			ns = fmt.Sprintf(vn.String(), l.Format(d))
			// Formatted only the digits of the value-name base or above.
			nf = d * uint64(math.Pow10(int(nb)))

		} else {
			// Not a label, maps directly to the name
			ns = vn.String()
			// formatted entire value
			nf = vn.Value
		}
		digits = append(digits, ns)
		nn := i - nf // The next number to format

		// insert separator if needed
		if nn > 0 {
			if sp := l.seperatorFor(i); sp != nil && nn < sp.Value {
				// If separator reverse, collect next value and insert with separator, prior to last value (just inserted)
				if sp.ReverseDigits {
					digits = insertIntoSlice(digits, []string{l.Format(nn), sp.String()}, len(digits)-1)
					nn = 0
				} else {
					digits = append(digits, sp.String())
				}
			}
		}

		if nn == 0 {
			// Last base been processed
			break
		}
		i = nn
	}
	return strings.Join(digits, l.digitSpace)
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

// seperatorFor returns the separator with the biggest value available which is <= given v
func (l language) seperatorFor(v uint64) *valueSeperator {
	var found *valueSeperator
	for _, sp := range l.separators {
		if sp.Value > v {
			break
		}
		found = sp
	}
	return found
}

func (l *language) UnmarshalJSON(bytes []byte) error {
	// using standard json decoding, into a psudo instance, then sorts bases before assigning to this.
	var lp struct {
		Title        string            `json:"title"`
		Names        []*valueName      `json:"names"`
		Separators   []*valueSeperator `json:"separators,omitempty"`
		MinusLabel   string            `json:"minus"`
		NoDigitSpace bool              `json:"no-digit-space"`
	}
	if err := json.Unmarshal(bytes, &lp); err != nil {
		return err
	}
	// sort names lowest value first
	sort.Slice(lp.Names, func(i, j int) bool {
		return lp.Names[i].Value < lp.Names[j].Value
	})
	// sort delimiters lowest value first
	sort.Slice(lp.Separators, func(i, j int) bool {
		return lp.Separators[i].Value < lp.Separators[j].Value
	})
	var ds string
	if !lp.NoDigitSpace {
		ds = " "
	}
	l.title = lp.Title
	l.names = lp.Names
	l.separators = lp.Separators
	l.minusLabel = lp.MinusLabel
	l.digitSpace = ds

	// validate
	if len(l.names) == 0 {
		return fmt.Errorf("no names found")
	}
	if l.names[0].Value != 0 {
		return fmt.Errorf("no zero value title found")
	}
	return nil
}

func insertIntoSlice(s, insert []string, index int) []string {
	if index < len(s) {
		insert = append(insert, s[index:]...)
	}
	return append(s[:index], insert...)
}

func nameToFilePath(name string) string {
	if filepath.Ext(name) == "" {
		name = strings.Join([]string{name, "json"}, ".")
	}
	if filepath.Dir(name) != languageFolderName {
		name = filepath.Join(languageFolderName, name)
	}
	return name
}

func OpenLanguage(name string) (Language, error) {
	f, err := os.Open(nameToFilePath(name))
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
