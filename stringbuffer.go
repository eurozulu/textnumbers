package textnumbers

import (
	"bytes"
	"fmt"
	"strings"
)

// StringBuffer is a simple wrapper of a bytes Buffer to control spacing with added strings.
type StringBuffer interface {
	fmt.Stringer
	Append(s string) StringBuffer
	Reset()
}

type stringBuffer struct {
	buf *bytes.Buffer
}

func (sb stringBuffer) String() string {
	return sb.buf.String()
}

func (sb stringBuffer) Reset() {
	sb.buf.Reset()
}

func (sb stringBuffer) Append(s string) StringBuffer {
	if s == "" {
		return sb
	}
	if sb.buf.Len() > 0 && !strings.HasPrefix(s, " ") {
		sb.buf.WriteRune(' ')
	}
	sb.buf.WriteString(s)
	return sb
}

func NewStringBuffer(s ...string) StringBuffer {
	sb := &stringBuffer{buf: bytes.NewBuffer(nil)}
	for _, ss := range s {
		sb.Append(ss)
	}
	return sb
}
