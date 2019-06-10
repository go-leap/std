package ustd

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Flag represents a named command-line args flag, in the fashion of either
// `--flag-name value` or `--flag-name=value`, as well as both `--fn value` and `--fn=value`.
type Flag struct {
	// Name, not including any leading dashes
	Name string
	// Desc, the help text
	Desc string
	// Default, the fall-back value if none was specified by the user
	Default string
}

var Flags struct {
	// Flags that were registered via any of the `FlagOf*` funcs.
	Known []Flag
	// OnErr is called when parsing of `value` for flag `name` failed.
	OnErr func(name string, value string, err error)
	// AddShortNames adds support for name abbreviations, ie. for `--long-named-flag` we also check for `--lnf`.
	AddShortNames bool
}

// FlagOfDuration obtains `val` from the command-line argument flag named `name` or the `defaultVal`.
func FlagOfDuration(name string, defaultVal time.Duration, desc string) (val time.Duration) {
	return FlagOfOther(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := time.ParseDuration(s); return v, e },
		func(v interface{}) string { return v.(time.Duration).String() },
	).(time.Duration)
}

// FlagOfBool obtains `val` from the command-line argument flag named `name` or the `defaultVal`.
func FlagOfBool(name string, defaultVal bool, desc string) (val bool) {
	return FlagOfOther(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := strconv.ParseBool(s); return v, e },
		func(v interface{}) string { return strconv.FormatBool(v.(bool)) },
	).(bool)
}

// FlagOfUint obtains `val` from the command-line argument flag named `name` or the `defaultVal`.
func FlagOfUint(name string, defaultVal uint64, desc string) (val uint64) {
	return FlagOfOther(name, defaultVal, desc,
		func(s string) (interface{}, error) { v, e := strconv.ParseUint(s, 10, 64); return v, e },
		func(v interface{}) string { return strconv.FormatUint(v.(uint64), 10) },
	).(uint64)
}

// FlagOfStrings obtains `val` from the command-line argument flag named `name` (items joined by `sep`) or the `defaultVal`.
func FlagOfStrings(name string, defaultVal []string, sep string, desc string) (val []string) {
	return FlagOfOther(name, defaultVal, desc,
		func(s string) (interface{}, error) {
			if s == "" { // dont want a slice of 1 empty string usually..
				return []string{}, nil
			}
			return strings.Split(s, sep), nil
		},
		func(v interface{}) string { return strings.Join(v.([]string), sep) },
	).([]string)
}

// FlagOfString obtains `val` from the command-line argument flag named `name` or the `defaultVal`.
func FlagOfString(name string, defaultVal string, desc string) (val string) {
	return flagOfString(flagReg(name, desc, defaultVal))
}

// FlagOfOther obtains `val` from the command-line argument flag named `name` or the `defaultVal`.
func FlagOfOther(name string, defaultVal interface{}, desc string, fromString func(string) (interface{}, error), toString func(interface{}) string) (val interface{}) {
	if str := flagOfString(flagReg(name, desc, toString(defaultVal))); str == "" {
		val = defaultVal
	} else if v, e := fromString(str); e != nil {
		if val = defaultVal; Flags.OnErr != nil {
			Flags.OnErr(name, str, e)
		}
	} else {
		val = v
	}
	return
}

func flagOfString(f *Flag) string {
	n2, n1 := "", "--"+f.Name
	_n2, _n1 := "", n1+"="
	if Flags.AddShortNames {
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
	for i := range Flags.Known {
		if Flags.Known[i].Name == name {
			return &Flags.Known[i]
		}
	}
	Flags.Known = append(Flags.Known, Flag{Name: name, Desc: desc, Default: defaultVal})
	return &Flags.Known[len(Flags.Known)-1]
}

// ShortName returns an abbreviation of `Name`, ie. `long-flag-name` will return `lfn`.
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
