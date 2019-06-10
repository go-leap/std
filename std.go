package ustd

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

// DoNowAndThenEvery invokes `do` immediately, then if `interval` is given, returns a `start` func that
// invokes `do` every `interval` `time.Duration` together with a `stop` func to `time.Ticker.Stop` doing so.
func DoNowAndThenEvery(interval time.Duration, should func() bool, do func()) (start func(), stop func()) {
	if do(); interval > 0 {
		ticker := time.NewTicker(interval)
		stop, start = ticker.Stop, func() {
			for range ticker.C {
				if should == nil || should() {
					do()
				}
			}
		}
	}
	return
}

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

// For is a `sync.WaitGroup`-based parallel `for` loop equivalent.
func For(numIter int, on func(int)) {
	var wait sync.WaitGroup
	wait.Add(numIter)
	do := func(i int) { on(i); wait.Done() }
	for i := 0; i < numIter; i++ {
		go do(i)
	}
	wait.Wait()
}

// IfNil returns `val` if not `nil`, else `thenFallbackTo`.
func IfNil(val interface{}, thenFallbackTo interface{}) interface{} {
	if val == nil {
		return thenFallbackTo
	}
	return val
}

// Time records a start time and returns a `func` to compute the `time.Duration` since then.
func Time() func() time.Duration {
	starttime := time.Now().UnixNano()
	return func() time.Duration {
		return time.Duration(time.Now().UnixNano() - starttime)
	}
}

// WriteLines returns a `func` that writes the specified `lines` to `w`.
func WriteLines(w io.Writer) func(...string) {
	return func(lines ...string) {
		if len(lines) > 0 {
			b := make([]byte, 0, len(lines)*len(lines[0]))
			for i := range lines {
				b = append(append(b, lines[i]...), '\n')
			}
			_, _ = w.Write(b)
		}
	}
}
