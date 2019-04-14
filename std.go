package ustd

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
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

func For(numIter int, on func(int)) {
	var wait sync.WaitGroup
	wait.Add(numIter)
	do := func(i int) { on(i); wait.Done() }
	for i := 0; i < numIter; i++ {
		go do(i)
	}
	wait.Wait()
}

func IfNil(val interface{}, thenFallbackTo interface{}) interface{} {
	if val == nil {
		return thenFallbackTo
	}
	return val
}

func Time() func() time.Duration {
	starttime := time.Now().UnixNano()
	return func() time.Duration {
		return time.Duration(time.Now().UnixNano() - starttime)
	}
}

func Write(to io.Writer) func(string, int) {
	return func(s string, n int) {
		b := make([]byte, 0, len(s)*n)
		for i := 0; i < n; i++ {
			b = append(b, s...)
		}
		_, _ = to.Write(b)
	}
}

func WriteLines(to io.Writer) func(...string) {
	return func(lns ...string) {
		if len(lns) > 0 {
			b := make([]byte, 0, len(lns)*len(lns[0]))
			for i := range lns {
				b = append(append(b, lns[i]...), '\n')
			}
			_, _ = to.Write(b)
		}
	}
}
