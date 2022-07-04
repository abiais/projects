package is_palindrome_func

import (
	"testing"
)

func TestIs_palindrom_empty(t *testing.T) {
	input := ""
	want := false
	got, _ := Is_palindrome(input)
	if got != want {
		t.Errorf("is_palindrome('') = %t; want %t", got, want)
	}
}

func TestIs_palindrom_palindrom(t *testing.T) {
	input := "101"
	want := true
	got, _ := Is_palindrome(input)
	if got != want {
		t.Errorf("is_palindrome('101') = %t; want %t", got, want)
	}
}

func TestIs_palindrom_not_palindrom(t *testing.T) {
	input := "123"
	want := false
	got, _ := Is_palindrome(input)
	if got != want {
		t.Errorf("is_palindrome('123') = %t; want %t", got, want)
	}
}
