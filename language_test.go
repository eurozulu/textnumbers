package textnumbers

import (
	"math"
	"testing"
)

func TestOpenLanguage(t *testing.T) {
	if _, err := openTestLanguage(); err != nil {
		t.Fatal(err)
	}
}

func TestLanguage_Title(t *testing.T) {
	l, _ := openTestLanguage()
	if l.Title() != "English" {
		t.Errorf("unexpected title.  Expected '%s' found '%s'", "English", l.Title())
	}
}

func TestLanguage_Format_Names(t *testing.T) {
	l, _ := openTestLanguage()
	e := "zero"
	f := l.Format(0)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "one"
	f = l.Format(1)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "five"
	f = l.Format(5)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "nine"
	f = l.Format(9)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "ten"
	f = l.Format(10)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "eleven"
	f = l.Format(11)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "fifteen"
	f = l.Format(15)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "nineteen"
	f = l.Format(19)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "twenty"
	f = l.Format(20)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}

	e = "twenty one"
	f = l.Format(21)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "twenty nine"
	f = l.Format(29)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "thirty"
	f = l.Format(30)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "thirty one"
	f = l.Format(31)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "fourty"
	f = l.Format(40)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "fifty"
	f = l.Format(50)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "sixty"
	f = l.Format(60)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "seventy"
	f = l.Format(70)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "eighty"
	f = l.Format(80)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "ninety"
	f = l.Format(90)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "ninety nine"
	f = l.Format(99)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
}

func TestLanguage_Format_Labels(t *testing.T) {
	l, _ := openTestLanguage()
	e := "one hundred"
	f := l.Format(100)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "five hundred"
	f = l.Format(500)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "nine hundred"
	f = l.Format(900)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "one hundred and one"
	f = l.Format(101)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "five hundred and fifty five"
	f = l.Format(555)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "nine hundred and ninety nine"
	f = l.Format(999)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}

	e = "one thousand"
	f = l.Format(1000)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}

	e = "one thousand and one"
	f = l.Format(1001)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "one thousand and ten"
	f = l.Format(1010)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "one thousand one hundred"
	f = l.Format(1100)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "five thousand five hundred and fifty five"
	f = l.Format(5555)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "nine thousand nine hundred and ninety nine"
	f = l.Format(9999)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}

	e = "nineteen thousand nine hundred and ninety nine"
	f = l.Format(19999)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
	e = "ninety nine thousand nine hundred and ninety nine"
	f = l.Format(99999)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}

	e = "nine hundred and ninety nine thousand nine hundred and ninety nine"
	f = l.Format(999999)
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}

	e = "eighteen quintillion four thousand four hundred and sixty seven quadrillion fourty four trillion seventy three billion seven hundred and nine million five hundred and fifty one thousand six hundred and fifteen"
	f = l.Format(uint64(math.MaxUint64)) // 18446744073709551615
	if f != e {
		t.Errorf("unexpected number format.  Expected '%s' found '%s'", e, f)
	}
}

func openTestLanguage() (Language, error) {
	return OpenLanguage("english")
}
