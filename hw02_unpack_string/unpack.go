package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	stack := make([]rune, 0)
	runes := []rune(str)
	for idx, r := range runes {
		if unicode.IsDigit(r) {
			if len(stack) == 0 || idx < len(runes)-1 && unicode.IsDigit(runes[idx+1]) {
				return "", ErrInvalidString
			}
			if r == '0' {
				stack = stack[:len(stack)-1]
			} else {
				for i := 1; i < int(r)-'0'; i++ {
					stack = append(stack, stack[len(stack)-1])
				}
			}
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack), nil
}
