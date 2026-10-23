package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var digits strings.Builder

	for _, r := range phoneNumber {
		switch {
		case r >= '0' && r <= '9':
			digits.WriteRune(r)
		case r == '+' || r == '(' || r == ')' ||
			r == '-' || r == ' ' || r == '.':
			// Ignore common formatting characters.
		default:
			return "", errors.New("phone number contains an invalid character")
		}
	}

	number := digits.String()

	// Remove the NANP country code when present.
	if len(number) == 11 && number[0] == '1' {
		number = number[1:]
	}

	if len(number) != 10 {
		return "", errors.New("phone number must contain 10 digits")
	}

	// The area code and exchange code must each start with 2–9.
	if number[0] < '2' || number[0] > '9' ||
		number[3] < '2' || number[3] > '9' {
		return "", errors.New("area and exchange codes must start with 2–9")
	}

	return number, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
