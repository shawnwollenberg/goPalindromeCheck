package palindrome

import "testing"

func TestNotPalindrome(t *testing.T) {
	want := false
	if got := palindrome("shawn"); got != want {
		t.Errorf("Palindrome() = %t, want %t", got, want)
	}
}
func TestPalindrome(t *testing.T) {
	want := true
	if got := palindrome("hannah"); got != want {
		t.Errorf("Palindrome() = %t, want %t", got, want)
	}
}
