package textnumbers

import (
	"testing"
)

const testValue1Count = 7

var testValue1 Number = 1234560
var testValue1Values = BaseNumbers(testValue1)

//[]Number{testValue1, 234560, 34560, 4560, 560, 60, 0}

func TestNumber_DigitCount(t *testing.T) {
	dc := testValue1.DigitCount()
	if dc != testValue1Count {
		t.Errorf("Unexpected digit count  expected %d, found %d", testValue1Count, dc)
	}
}

func TestNumber_DigitAt(t *testing.T) {
	dc := int(testValue1.DigitCount())
	// assumes testvalue is 1234560
	for i := 0; i < dc; i++ {
		expect := i + 1

		b := Base(dc - expect)
		d := testValue1.DigitAt(b)
		if expect == dc {
			// last digit is zero
			expect = 0
		}
		if d != expect {
			t.Errorf("Unexpected digit expected %d, found %d", expect, d)
		}
	}
}

func TestNumber_ValueAt(t *testing.T) {
	dc := int(testValue1.DigitCount())
	if len(testValue1Values) != dc {
		t.Fatal("Test values do not align. testValue1Values must have number of elements equal to testValue1 base count")
	}
	for i := 0; i < dc; i++ {
		b := Base(dc - i - 1)
		v := Number(testValue1.ValueAt(b))
		if v != testValue1Values[i] {
			t.Errorf("unexpected valueAt base %d.  Expected %d, found %d", b, testValue1Values[i], v)
		}
	}
}

func BaseNumbers(num Number) []Number {
	dc := num.DigitCount()
	nums := make([]Number, dc)
	for i := 0; i < dc; i++ {
		b := Base(dc - i - 1)
		nums[i] = Number(num.ValueAt(b))
	}
	return nums
}
