package textnumbers

import (
	"math"
)

// Number is a value which can be formatted into text
type Number int64

// Base is the base of a number, representing the position of a digit.
type Base byte

// DigitAt returns the single digit at the given base position.
// base zero returns the right most digit, base n returns the left most digit, where n is the basecount of the Number.
// a given base greater than the Number Base count (e.g. base 5 on '123') will return zero.
func (n Number) DigitAt(base Base) int {
	return int(n.ValueAt(base) / int64(math.Pow10(int(base))))
}

func (n Number) DigitsAt(base Base) int64 {
	p := math.Pow10(int(base))
	return int64(n) / int64(p)
}

// ValueAt returns the value of the number trimmed at the given base.
// any digits with a base highter than the given base are subtracted, leaving just the digits
// at the given base and lower.
// e.g. for the Number 1234:
// base		returns
// 3		1234
// 2		234
// 1		34
// 0		4
func (n Number) ValueAt(base Base) int64 {
	p := int64(math.Pow10(int(base) + 1))
	return int64(n) % p
}

// DigitCount retruns the number of digits in the Number.
// All Numbers have one or more digits.
func (n Number) DigitCount() int {
	return int(math.Log10(float64(n))) + 1
}
