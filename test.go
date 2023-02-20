package main

import (
	"fmt"
	"regexp"
)

func validatePhoneNumber(phoneNumber string) bool {
	pattern := `^84\d{9}$`
	matched, err := regexp.MatchString(pattern, phoneNumber)
	return err == nil && matched
}
func main() {
	phoneNumber := "8470348l0999"
	if validatePhoneNumber(phoneNumber) {
		fmt.Printf("Phone number %s is valid\n", phoneNumber)
	} else {
		fmt.Printf("Phone number %s is invalid\n", phoneNumber)
	}
}
