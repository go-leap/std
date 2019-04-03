# ustd
--
    import "github.com/go-leap/std"


## Usage

#### func  For

```go
func For(numIter int, on func(int))
```

#### func  JsonDecodeFromFile

```go
func JsonDecodeFromFile(fromFilePath string, intoDestination interface{}) (err error)
```
JsonDecodeFromFile opens the specified file and attempts to JSON-decode into the
specified destination location.

#### func  JsonEncodeToFile

```go
func JsonEncodeToFile(from interface{}, toFilePath string) (err error)
```
JsonEncodeToFile creates the specified file and attempts to JSON-encode into it
from the specified source.

#### func  ReadAll

```go
func ReadAll(r io.Reader, initialBufSize int64) (data []byte, err error)
```

#### func  Time

```go
func Time() func() time.Duration
```

#### type BytesReader

```go
type BytesReader struct {
	Data []byte
}
```


#### func (*BytesReader) Read

```go
func (me *BytesReader) Read(p []byte) (n int, err error)
```

#### type BytesWriter

```go
type BytesWriter struct{ Data []byte }
```


#### func (*BytesWriter) Bytes

```go
func (me *BytesWriter) Bytes() []byte
```

#### func (*BytesWriter) TrimSuffix

```go
func (me *BytesWriter) TrimSuffix(suffix byte)
```

#### func (*BytesWriter) Write

```go
func (me *BytesWriter) Write(b []byte) (int, error)
```

#### func (*BytesWriter) WriteByte

```go
func (me *BytesWriter) WriteByte(b byte)
```

#### func (*BytesWriter) WriteString

```go
func (me *BytesWriter) WriteString(b string)
```

#### func (*BytesWriter) WriteTo

```go
func (me *BytesWriter) WriteTo(w io.Writer) (int64, error)
```
WriteTo implements the `io.WriterTo` interface.
