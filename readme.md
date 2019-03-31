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

#### type Buf

```go
type Buf struct {
}
```


#### func  NewBuf

```go
func NewBuf(b []byte) *Buf
```

#### func (*Buf) Bytes

```go
func (me *Buf) Bytes() []byte
```

#### func (*Buf) TrimSuffix

```go
func (me *Buf) TrimSuffix(suffix byte)
```

#### func (*Buf) Write

```go
func (me *Buf) Write(b []byte) (int, error)
```

#### func (*Buf) WriteByte

```go
func (me *Buf) WriteByte(b byte)
```

#### func (*Buf) WriteString

```go
func (me *Buf) WriteString(b string)
```

#### func (*Buf) WriteTo

```go
func (me *Buf) WriteTo(w io.Writer) (int64, error)
```
WriteTo implements the `io.WriterTo` interface.
