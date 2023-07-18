package misc

import "flag"

func IsIdent(value byte) bool {
	return (value >= 'A' && value <= 'Z') || (value >= 'a' && value <= 'z') || (value == '_')
}

func FlagExists(name string) bool {
	exists := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			exists = true
		}
	})
	return exists
}
