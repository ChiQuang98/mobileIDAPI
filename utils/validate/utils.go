package validate

import "regexp"

func ValidatePhoneNumber(phoneNumber string) bool {
	pattern := `^84\d{9}$`
	matched, err := regexp.MatchString(pattern, phoneNumber)
	return err == nil && matched
}
