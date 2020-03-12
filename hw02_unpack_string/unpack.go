package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var isDigit = regexp.MustCompile(`[0-9][0-9]`)
	var isOnlyNumber = regexp.MustCompile(`^[0-9]+$`)
	var containsDigit = regexp.MustCompile(`[0-9]`)
	var result strings.Builder
	var repeat bool

	if isDigit.MatchString(input) || isOnlyNumber.MatchString(input) {
		return "", ErrInvalidString
	}

	runes := []rune(input)

	if !containsDigit.MatchString(input) || len(runes) == 1 {
		return input, nil
	}

	for i := 0; i < len(runes)-1; i++ {
		repeat = false
		//fmt.Printf("Rune %v is '%c'\n", i, runes[i])
		if unicode.IsDigit(runes[i+1]) {
			repeat = true
			max, _ := strconv.Atoi(string(runes[i+1]))
			for j := 0; j < max; j++ {
				result.WriteRune(runes[i])
			}
			i++
		} else {
			result.WriteRune(runes[i])
		}
	}
	if repeat && !containsDigit.MatchString(string(runes[len(runes)-1])) {
		result.WriteRune(runes[len(runes)-1])
	}

	return result.String(), nil
}
