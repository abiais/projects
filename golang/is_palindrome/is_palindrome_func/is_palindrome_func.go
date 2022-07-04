package is_palindrome_func

import (
	"errors"
	"strings"
)

func Is_palindrome(text string) (bool, error) {
	if strings.TrimSpace(text) == "" {
		return false, errors.New("empty string")
	}
	a := strings.Split(text, "")
	b := make([]string, len(a))
	for i := range a {
		b[len(a)-i-1] = a[i]
	}
	return strings.Join(a, "") == strings.Join(b, ""), nil
}
