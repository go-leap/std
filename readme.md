# ustd
--
    import "github.com/go-leap/std"


## Usage

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
