package ustd

import (
	"os"
	"strconv"
	"time"
)

type Flag struct {
	Name    string
	Desc    string
	Default string
}

var Flags []Flag
var FlagsOnErr func(string, string, error)

func flagAdd(name string, desc string, defaultVal string) {
	for i := range Flags {
		if Flags[i].Name == name {
			return
		}
	}
	Flags = append(Flags, Flag{Name: name, Desc: desc, Default: defaultVal})
}

func FlagOfString(name string, defaultVal string, desc string) string {
	flagAdd(name, desc, defaultVal)
	n := "--" + name
	_n := n + "="
	for i, l, il := 1, len(os.Args), len(os.Args)-1; i < l; i++ {
		if arg := os.Args[i]; arg == n && i < il {
			val := os.Args[i+1]
			if len(val) >= 2 && val[0] == '-' && val[1] == '-' {
				continue
			}
			return val
		} else if len(arg) > len(_n) && arg[:len(_n)] == _n {
			return arg[len(_n):]
		}
	}
	return defaultVal
}

func FlagOfDuration(name string, defaultVal time.Duration, desc string) (val time.Duration) {
	return FlagX(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := time.ParseDuration(s); return v, e },
		func(v interface{}) string { return v.(time.Duration).String() },
	).(time.Duration)
}

func FlagOfBool(name string, defaultVal bool, desc string) bool {
	return FlagX(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := strconv.ParseBool(s); return v, e },
		func(v interface{}) string { return strconv.FormatBool(v.(bool)) },
	).(bool)
}

func FlagX(name string, defaultVal interface{}, desc string, fromString func(string) (interface{}, error), toString func(interface{}) string) (val interface{}) {
	flagAdd(name, desc, toString(defaultVal))
	if str := FlagOfString(name, "", desc); str == "" {
		val = defaultVal
	} else if v, e := fromString(str); e != nil {
		if FlagsOnErr != nil {
			FlagsOnErr(name, str, e)
		} else {
			panic(e)
		}
	} else {
		val = v
	}
	return
}
