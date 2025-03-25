package validator

import "fmt"

const (
	mustBeNotEmpty = "Must be not empty"

	invalidUrl         = "Invalid URL"
	invalidPhoneNumber = "Invalid phone number"
	invalidEmail       = "Invalid email"
)

func invalidDate(format string) string {
	return fmt.Sprintf("Date must be in %s format", format)
}

func maxChars(total int) string {
	return fmt.Sprintf("Maximum %d characters", total)
}

func mustBeChars(total int) string {
	return fmt.Sprintf("Must be %d characters", total)
}

func mustBeGt(value int) string {
	return fmt.Sprintf("Must be greater than %d", value)
}

func mustBeGte(value int) string {
	return fmt.Sprintf("Must be greater than or equal to %d", value)
}

func mustBeLt(value int) string {
	return fmt.Sprintf("Must be less than %d", value)
}

func mustBetween(min, max int) string {
	return fmt.Sprintf("Must between %d and %d", min, max)
}

func mustBeLte(value int) string {
	return fmt.Sprintf("Must be less than or equal to %d", value)
}

func mustBe(values ...string) string {
	valuesLen := len(values)
	if valuesLen == 0 {
		return "Invalid value"
	}

	result := "Must be "

	for i, value := range values {
		result += fmt.Sprintf("'%v'", value)

		if i+1 < valuesLen {
			result += ", "
		}
	}

	return result
}

type Error struct {
	Field   string
	Message string
}

func newError(field, message string) Error {
	return Error{
		Field:   field,
		Message: message,
	}
}
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
