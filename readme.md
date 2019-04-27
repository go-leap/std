# ustd
--
    import "github.com/go-leap/std"


## Usage

```go
var (
	Flags              []Flag
	FlagsOnErr         func(string, string, error)
	FlagsAddShortNames bool
)
```

#### func  DoNowAndThenEvery

```go
func DoNowAndThenEvery(interval time.Duration, should func() bool, do func()) (start func(), stop func())
```

#### func  FlagOfBool

```go
func FlagOfBool(name string, defaultVal bool, desc string) bool
```

#### func  FlagOfDuration

```go
func FlagOfDuration(name string, defaultVal time.Duration, desc string) (val time.Duration)
```

#### func  FlagOfString

```go
func FlagOfString(name string, defaultVal string, desc string) string
```

#### func  FlagOfStrings

```go
func FlagOfStrings(name string, defaultVal []string, sep string, desc string) []string
```

#### func  FlagOfUint

```go
func FlagOfUint(name string, defaultVal uint64, desc string) uint64
```

#### func  FlagOther

```go
func FlagOther(name string, defaultVal interface{}, desc string, fromString func(string) (interface{}, error), toString func(interface{}) string) (val interface{})
```

#### func  For

```go
func For(numIter int, on func(int))
```

#### func  HashTwo

```go
func HashTwo(oldSum1, oldSum2 uint64, data []byte) (newSum1, newSum2 uint64, bothUnchanged bool)
```

#### func  IfNil

```go
func IfNil(val interface{}, thenFallbackTo interface{}) interface{}
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

#### func  WriteLines

```go
func WriteLines(to io.Writer) func(...string)
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

#### type Flag

```go
type Flag struct {
	Name    string
	Desc    string
	Default string
}
```


#### func (*Flag) ShortName

```go
func (me *Flag) ShortName() (shortForm string)
```

#### type Writer

```go
type Writer struct {
	io.Writer
	On struct {
		AfterEveryNth                int
		Byte                         byte
		Do                           func(int) bool
		ButDontCountImmediateRepeats bool
	}
}
```


#### func (*Writer) RestartOnDo

```go
func (me *Writer) RestartOnDo()
```

#### func (*Writer) SuspendOnDo

```go
func (me *Writer) SuspendOnDo()
```

#### func (*Writer) Write

```go
func (me *Writer) Write(p []byte) (n int, err error)
```
