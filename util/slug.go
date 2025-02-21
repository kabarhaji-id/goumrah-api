package util

import "strings"

func Slug(raw string) (slug string) {
	for _, r := range raw {
		if r == ' ' {
			slug += "-"
		} else {
			slug += strings.ToLower(string(r))
		}
	}

	return
}
