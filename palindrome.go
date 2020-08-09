package Palindrome

import (
	rev "github.com/shawnwollenberg/goReverseString"
)

func Palindrome(str string) bool {
	return str == rev.Reversestring(str)
}
