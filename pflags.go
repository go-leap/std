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

func flagReg(name string, desc string, defaultVal string) *Flag {
	for i := range Flags {
		if Flags[i].Name == name {
			return &Flags[i]
		}
	}
	Flags = append(Flags, Flag{Name: name, Desc: desc, Default: defaultVal})
	return &Flags[len(Flags)-1]
}

func flagOfString(name string, defaultVal string, desc string) (string, bool) {
	n2, n1 := "", "--"+name
	_n2, _n1, nameexists := "", n1+"=", false
	if f := flagReg(name, desc, defaultVal); FlagsAddShortNames {
		n2 = "--" + f.ShortName()
		_n2 = n2 + "="
	}
	for i, l, il := 1, len(os.Args), len(os.Args)-1; i < l; i++ {
		if arg := os.Args[i]; arg == n1 || arg == n2 {
			if nameexists = true; i < il {
				val := os.Args[i+1]
				if len(val) >= 2 && val[0] == '-' && val[1] == '-' {
					continue
				}
				return val, nameexists
			}
		} else if len(arg) > len(_n1) && arg[:len(_n1)] == _n1 {
			return arg[len(_n1):], nameexists
		} else if len(arg) > len(_n2) && arg[:len(_n2)] == _n2 {
			return arg[len(_n2):], nameexists
		}
	}
	return defaultVal, nameexists
}

func FlagOfString(name string, defaultVal string, desc string) string {
	s, _ := flagOfString(name, defaultVal, desc)
	return s
}

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

func FlagOfStrings(name string, defaultVal []string, sep string, desc string) []string {
	return FlagOther(name, defaultVal, desc,
		func(s string) (interface{}, error) { return strings.Split(s, sep), nil },
		func(v interface{}) string { return strings.Join(v.([]string), sep) },
	).([]string)
}

func FlagOther(name string, defaultVal interface{}, desc string, fromString func(string) (interface{}, error), toString func(interface{}) string) (val interface{}) {
	flagReg(name, desc, toString(defaultVal))
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
