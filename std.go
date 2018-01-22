package ustd

import (
	"encoding/json"
	"os"
)

// JsonDecodeFromFile opens the specified file and attempts to JSON-decode into the specified destination location.
func JsonDecodeFromFile(fromFilePath string, intoDestination interface{}) (err error) {
	var f *os.File
	if f, err = os.Open(fromFilePath); err == nil {
		defer f.Close()
		err = json.NewDecoder(f).Decode(intoDestination)
	}
	return
}

// JsonEncodeToFile creates the specified file and attempts to JSON-encode into it from the specified source.
func JsonEncodeToFile(from interface{}, toFilePath string) (err error) {
	var f *os.File
	if f, err = os.Create(toFilePath); err == nil {
		defer f.Close()
		err = json.NewEncoder(f).Encode(from)
	}
	return
}
