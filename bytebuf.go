package ustd

import (
	"io"
)

const bufPad = 128

type Buf struct{ b []byte }

func NewBuf(b []byte) *Buf {
	return &Buf{b: b}
}

func (me *Buf) Bytes() []byte { return me.b }

func (me *Buf) WriteByte(b byte) {
	l, c := len(me.b), cap(me.b)
	if l == c {
		old := me.b
		me.b = make([]byte, l+1, l+l+bufPad) // the constant extra padding: if l is tiny, it helps much; if large, it hurts none
		copy(me.b[:l], old)
	} else {
		me.b = me.b[:l+1]
	}
	me.b[l] = b
}

func (me *Buf) Write(b []byte) (int, error) {
	l, c, n := len(me.b), cap(me.b), len(b)
	if ln := l + n; ln > c {
		old := me.b
		me.b = make([]byte, ln, ln+ln+bufPad)
		copy(me.b[:l], old)
	} else {
		me.b = me.b[:ln]
	}
	copy(me.b[l:], b)
	return n, nil
}

func (me *Buf) WriteString(b string) {
	l, c, n := len(me.b), cap(me.b), len(b)
	if ln := l + n; ln > c {
		old := me.b
		me.b = make([]byte, ln, ln+ln+bufPad)
		copy(me.b[:l], old)
	} else {
		me.b = me.b[:ln]
	}
	copy(me.b[l:], b)
}

// WriteTo implements the `io.WriterTo` interface.
func (me *Buf) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(me.b)
	return int64(n), err
}

func (me *Buf) TrimSuffix(suffix byte) {
	for n := len(me.b) - 1; n >= 0 && suffix == me.b[n]; n = len(me.b) - 1 {
		me.b = me.b[:n]
	}
}
