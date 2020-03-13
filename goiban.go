package goiban

import (
	"fmt"
	"strconv"
	"strings"
)

// Validate Checks if provided IBAN is valid
func Validate(iban string) (bool, string) {

	var trimmed string = strings.ToUpper(strings.Replace(iban, " ", "", -1))

	if len(trimmed) > 34 || len(trimmed) <= 5 {
		return false, "Invalid length" // wrong length
	}

	countryCode := trimmed[0:2]
	first := trimmed[0:4]
	last := trimmed[4:]

	if !isLetter(countryCode) {
		return false, "No country code" // no country code
	}

	output := mod(alphToDec(last+first), 97)

	if output == 1 {
		return true, "Valid IBAN"
	}
	return false, "Invalid IBAN"
}

func isLetter(str string) bool {
	for _, character := range str {
		if character < 'A' || character > 'Z' {
			return false
		}
	}
	return true
}

func alphToDec(str string) string {
	var decString strings.Builder
	for _, val := range str {
		if isLetter(string(val)) {
			tmpAlphabet, _ := strconv.Atoi(fmt.Sprintf("%d", val))
			output := tmpAlphabet - 55
			decString.WriteString(strconv.Itoa(output))
		} else {
			tmpDecimal, _ := strconv.Atoi(fmt.Sprintf("%c", val))
			decString.WriteString(strconv.Itoa(tmpDecimal))
		}
	}
	return decString.String()
}

func mod(num string, a int) int {
	res := 0
	for _, val := range num {
		workVal, _ := strconv.Atoi(fmt.Sprintf("%c", val))
		res = (res*10 + workVal) % a
	}
	return res
}
