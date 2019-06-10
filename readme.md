# ustd
--
    import "github.com/go-leap/std"


## Usage

```go
var Flags struct {
	// Flags that were registered via any of the `FlagOf*` funcs.
	Known []Flag
	// OnErr is called when parsing of `value` for flag `name` failed.
	OnErr func(name string, value string, err error)
	// AddShortNames adds support for name abbreviations, ie. for `--long-named-flag` we also check for `--lnf`.
	AddShortNames bool
}
```

#### func  DoNowAndThenEvery

```go
func DoNowAndThenEvery(interval time.Duration, should func() bool, do func()) (start func(), stop func())
```
DoNowAndThenEvery invokes `do` immediately, then if `interval` is given, returns
a `start` func that invokes `do` every `interval` `time.Duration` together with
a `stop` func to `time.Ticker.Stop` doing so.

#### func  FlagOfBool

```go
func FlagOfBool(name string, defaultVal bool, desc string) (val bool)
```
FlagOfBool obtains `val` from the command-line argument flag named `name` or the
`defaultVal`.

#### func  FlagOfDuration

```go
func FlagOfDuration(name string, defaultVal time.Duration, desc string) (val time.Duration)
```
FlagOfDuration obtains `val` from the command-line argument flag named `name` or
the `defaultVal`.

#### func  FlagOfOther

```go
func FlagOfOther(name string, defaultVal interface{}, desc string, fromString func(string) (interface{}, error), toString func(interface{}) string) (val interface{})
```
FlagOfOther obtains `val` from the command-line argument flag named `name` or
the `defaultVal`.

#### func  FlagOfString

```go
func FlagOfString(name string, defaultVal string, desc string) (val string)
```
FlagOfString obtains `val` from the command-line argument flag named `name` or
the `defaultVal`.

#### func  FlagOfStrings

```go
func FlagOfStrings(name string, defaultVal []string, sep string, desc string) (val []string)
```
FlagOfStrings obtains `val` from the command-line argument flag named `name`
(items joined by `sep`) or the `defaultVal`.

#### func  FlagOfUint

```go
func FlagOfUint(name string, defaultVal uint64, desc string) (val uint64)
```
FlagOfUint obtains `val` from the command-line argument flag named `name` or the
`defaultVal`.

#### func  For

```go
func For(numIter int, on func(int))
```
For is a `sync.WaitGroup`-based parallel `for` loop equivalent.

#### func  HashTwo

```go
func HashTwo(oldSum1, oldSum2 uint64, data []byte) (newSum1, newSum2 uint64, bothUnchanged bool)
```
HashTwo computes at least `newSum1` and (if equal to non-0 `oldSum1`) also
`newSum2`. The first uses wyhash, the second xxhash.

#### func  IfNil

```go
func IfNil(val interface{}, thenFallbackTo interface{}) interface{}
```
IfNil returns `val` if not `nil`, else `thenFallbackTo`.

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
ReadAll is a somewhat leaner alternative to `ioutil.ReadAll`, for `io.Reader`s
known to be of limited (not potentially-infinite or RAM-exceeding) size.

#### func  Time

```go
func Time() func() time.Duration
```
Time records a start time and returns a `func` to compute the `time.Duration`
since then.

#### func  WriteLines

```go
func WriteLines(w io.Writer) func(...string)
```
WriteLines returns a `func` that writes the specified `lines` to `w`.

#### type BytesReader

```go
type BytesReader struct {
	Data []byte
}
```

BytesReader implements `io.Reader` over a slice of `byte`s in a leaner manner
than the reader-and-writer `bytes.Buffer`.

#### func (*BytesReader) Read

```go
func (me *BytesReader) Read(p []byte) (n int, err error)
```
Read implements `io.Reader`.

#### type BytesWriter

```go
type BytesWriter struct{ Data []byte }
```

BytesWriter implements `io.Writer` over a slice of `byte`s in a leaner manner
than the reader-and-writer `bytes.Buffer`.

#### func (*BytesWriter) Bytes

```go
func (me *BytesWriter) Bytes() []byte
```
Bytes returns `me.Data` and aids in compatibility with `bytes.Buffer`.

#### func (*BytesWriter) Reset

```go
func (me *BytesWriter) Reset()
```
Reset sets the `len` of `me.Data` to 0.

#### func (*BytesWriter) TrimSuffix

```go
func (me *BytesWriter) TrimSuffix(suffix byte)
```
TrimSuffix ensures that `me.Data` is not suffixed by the specified `suffix`
byte.

#### func (*BytesWriter) Write

```go
func (me *BytesWriter) Write(b []byte) (int, error)
```
Write implements `io.Writer` and writes the specified `byte`s to `me.Data`,
always returning `len(b), nil`.

#### func (*BytesWriter) WriteByte

```go
func (me *BytesWriter) WriteByte(b byte)
```
WriteByte writes a single `byte` to `me.Data`.

#### func (*BytesWriter) WriteString

```go
func (me *BytesWriter) WriteString(b string)
```
WriteString writes the specified `string` to `me.Data`.

#### func (*BytesWriter) WriteTo

```go
func (me *BytesWriter) WriteTo(w io.Writer) (int64, error)
```
WriteTo implements the `io.WriterTo` interface.

#### type Flag

```go
type Flag struct {
	// Name, not including any leading dashes
	Name string
	// Desc, the help text
	Desc string
	// Default, the fall-back value if none was specified by the user
	Default string
}
```

Flag represents a named command-line args flag, in the fashion of either
`--flag-name value` or `--flag-name=value`, as well as both `--fn value` and
`--fn=value`.

#### func (*Flag) ShortName

```go
func (me *Flag) ShortName() (shortForm string)
```
ShortName returns an abbreviation of `Name`, ie. `long-flag-name` will return
`lfn`.

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

Writer wraps any given `io.Writer` with additional byte-based processing on
`Write`s.

#### func (*Writer) RestartOnDo

```go
func (me *Writer) RestartOnDo()
```
RestartOnDo resets the counter for `On.Do` calls during `Write`s.

#### func (*Writer) SuspendOnDo

```go
func (me *Writer) SuspendOnDo()
```
SuspendOnDo causes `On.Do` not to be called during `Write`s until `RestartOnDo`.

#### func (*Writer) Write

```go
func (me *Writer) Write(p []byte) (n int, err error)
```
Write calls `me.Writer.Write` and if `me.On.AfterEveryNth` and `me.On.Do` are
set, traverses `p` to conditionally call `me.On.Do` with an internal counter.
