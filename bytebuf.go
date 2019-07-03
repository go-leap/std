package ustd

import (
	"io"
	"unsafe"
)

// BytesReader implements `io.Reader` over a slice of `byte`s in a leaner manner than the reader-and-writer `bytes.Buffer`.
type BytesReader struct {
	Data []byte
	pos  int
}

// Read implements `io.Reader`.
func (me *BytesReader) Read(p []byte) (n int, err error) {
	if ld := len(me.Data); ld == me.pos {
		err = io.EOF
	} else if lp := len(p); lp > 0 {
		end := me.pos + lp
		if n = lp; end > ld {
			end = ld
			if n = end - me.pos; n == 0 {
				err = io.EOF
			}
		}
		if n > 0 {
			copy(p[:n], me.Data[me.pos:end])
			me.pos += n
		}
	}
	return
}

const bytesWriterPadding = 88

// BytesWriter implements `io.Writer` over a slice of `byte`s in a leaner manner than the reader-and-writer `bytes.Buffer`.
type BytesWriter struct{ Data []byte }

// Bytes returns `me.Data` and aids in compatibility with `bytes.Buffer`.
func (me *BytesWriter) Bytes() []byte { return me.Data }

// Reset sets the `len` of `me.Data` to 0.
func (me *BytesWriter) Reset() { me.Data = me.Data[0:0] }

func (me *BytesWriter) String() string {
	return *(*string)(unsafe.Pointer(&me.Data))
}

// WriteByte writes a single `byte` to `me.Data`.
func (me *BytesWriter) WriteByte(b byte) {
	l, c := len(me.Data), cap(me.Data)
	if l == c {
		old := me.Data
		me.Data = make([]byte, l+1, l+l+bytesWriterPadding) // the constant extra padding: if l is tiny, it helps some; if large, it hurts none
		copy(me.Data[:l], old)
	} else {
		me.Data = me.Data[:l+1]
	}
	me.Data[l] = b
}

// Write implements `io.Writer` and writes the specified `byte`s to `me.Data`, always returning `len(b), nil`.
func (me *BytesWriter) Write(b []byte) (int, error) {
	l, c, n := len(me.Data), cap(me.Data), len(b)
	if ln := l + n; ln > c {
		old := me.Data
		me.Data = make([]byte, ln, ln+ln+bytesWriterPadding)
		copy(me.Data[:l], old)
	} else {
		me.Data = me.Data[:ln]
	}
	copy(me.Data[l:], b)
	return n, nil
}

// WriteString writes the specified `string` to `me.Data`.
func (me *BytesWriter) WriteString(b string) {
	l, c, n := len(me.Data), cap(me.Data), len(b)
	if ln := l + n; ln > c {
		old := me.Data
		me.Data = make([]byte, ln, ln+ln+bytesWriterPadding)
		copy(me.Data[:l], old)
	} else {
		me.Data = me.Data[:ln]
	}
	copy(me.Data[l:], b)
}

// WriteTo implements the `io.WriterTo` interface.
func (me *BytesWriter) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(me.Data)
	return int64(n), err
}

// TrimSuffix ensures that `me.Data` is not suffixed by the specified `suffix` byte.
func (me *BytesWriter) TrimSuffix(suffix byte) {
	for n := len(me.Data) - 1; n >= 0 && suffix == me.Data[n]; n = len(me.Data) - 1 {
		me.Data = me.Data[:n]
	}
}

// ReadAll is a somewhat leaner alternative to `ioutil.ReadAll`, for `io.Reader`s known to be of limited (not potentially-infinite or RAM-exceeding) size.
func ReadAll(r io.Reader, initialBufSize int64) (data []byte, err error) {
	if initialBufSize <= 0 {
		initialBufSize = 128
	} else if m := initialBufSize % 8; m != 0 {
		initialBufSize += (8 - m)
	}
	data = make([]byte, initialBufSize)
	var n, i int
	for err == nil {
		if n, err = r.Read(data[i:]); n > 0 {
			if i += n; i >= len(data) && err == nil {
				data = append(data, make([]byte, len(data)*2)...)
			}
		}
	}
	if data = data[:i]; err == io.EOF {
		err = nil
	}
	return
}
