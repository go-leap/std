package ustd

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Flag struct {
	Name    string
	Desc    string
	Default string
}

var (
	Flags              []Flag
	FlagsOnErr         func(string, string, error)
	FlagsAddShortNames bool
)

func FlagOfDuration(name string, defaultVal time.Duration, desc string) (val time.Duration) {
	return FlagOther(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := time.ParseDuration(s); return v, e },
		func(v interface{}) string { return v.(time.Duration).String() },
	).(time.Duration)
}

func FlagOfBool(name string, defaultVal bool, desc string) bool {
	return FlagOther(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := strconv.ParseBool(s); return v, e },
		func(v interface{}) string { return strconv.FormatBool(v.(bool)) },
	).(bool)
}

func FlagOfUint(name string, defaultVal uint64, desc string) uint64 {
	return FlagOther(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := strconv.ParseUint(s, 10, 64); return v, e },
		func(v interface{}) string { return strconv.FormatUint(v.(uint64), 10) },
	).(uint64)
}

func FlagOfStrings(name string, defaultVal []string, sep string, desc string) []string {
	return FlagOther(name, defaultVal, desc,
		func(s string) (interface{}, error) {
			if s == "" { // dont want a slice of 1 empty string usually..
				return []string{}, nil
			}
			return strings.Split(s, sep), nil
		},
		func(v interface{}) string { return strings.Join(v.([]string), sep) },
	).([]string)
}

func FlagOfString(name string, defaultVal string, desc string) string {
	return flagOfString(flagReg(name, desc, defaultVal))
}

func FlagOther(name string, defaultVal interface{}, desc string, fromString func(string) (interface{}, error), toString func(interface{}) string) (val interface{}) {
	if str := flagOfString(flagReg(name, desc, toString(defaultVal))); str == "" {
		val = defaultVal
	} else if v, e := fromString(str); e != nil {
		if val = defaultVal; FlagsOnErr != nil {
			FlagsOnErr(name, str, e)
		}
	} else {
		val = v
	}
	return
}

func flagOfString(f *Flag) string {
	n2, n1 := "", "--"+f.Name
	_n2, _n1 := "", n1+"="
	if FlagsAddShortNames {
		n2 = "--" + f.ShortName()
		_n2 = n2 + "="
	}
	for i, l, il := 1, len(os.Args), len(os.Args)-1; i < l; i++ {
		if arg := os.Args[i]; arg == n1 || arg == n2 {
			if i < il {
				val := os.Args[i+1]
				if len(val) >= 2 && val[0] == '-' && val[1] == '-' {
					continue
				}
				return strings.TrimSpace(val)
			}
		} else if len(arg) > len(_n1) && arg[:len(_n1)] == _n1 {
			return strings.TrimSpace(arg[len(_n1):])
		} else if len(arg) > len(_n2) && arg[:len(_n2)] == _n2 {
			return strings.TrimSpace(arg[len(_n2):])
		}
	}
	return strings.TrimSpace(f.Default)
}

func flagReg(name string, desc string, defaultVal string) *Flag {
	for i := range Flags {
		if Flags[i].Name == name {
			return &Flags[i]
		}
	}
	Flags = append(Flags, Flag{Name: name, Desc: desc, Default: defaultVal})
	return &Flags[len(Flags)-1]
}

func (me *Flag) ShortName() (shortForm string) {
	lastwasdash := true
	for _, r := range me.Name {
		if r == '-' {
			lastwasdash = true
		} else {
			if lastwasdash {
				shortForm += string(r)
			}
			lastwasdash = false
		}
	}
	return
}
