package regexpattern

import "regexp"

var (
	phoneNumber *regexp.Regexp = nil
)

func PhoneNumber() *regexp.Regexp {
	if phoneNumber == nil {
		phoneNumber = regexp.MustCompile(`^(\+62|0)[0-9]{9,15}$`)
	}

	return phoneNumber
}
