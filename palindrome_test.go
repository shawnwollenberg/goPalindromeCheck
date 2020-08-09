package Palindrome

import "testing"

func TestNotPalindrome(t *testing.T) {
	want := false
	if got := Palindrome("shawn"); got != want {
		t.Errorf("Palindrome() = %t, want %t", got, want)
	}
}
func TestPalindrome(t *testing.T) {
	want := true
	if got := Palindrome("hannah"); got != want {
		t.Errorf("Palindrome() = %t, want %t", got, want)
	}
}
