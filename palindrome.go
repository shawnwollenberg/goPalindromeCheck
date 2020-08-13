package palindrome

import (
	rev "github.com/shawnwollenberg/goReverseString"
)

func palindrome(str string) bool {
	return str == rev.Reversestring(str)
}
