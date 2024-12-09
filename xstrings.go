// sdditional string functions not found in the stdlib such as HasPrefixFold,
// ContainsFold, Underscore, Dashersize, etc
package xstrings

import (
	"strings"
	"unicode"
)

// ContainsFold reports whether substr is within s, ignoring case.
// It returns true if substr is found within s, and false otherwise.
//
// Parameters:
//   - s: The string to search within.
//   - r: The substring to search for.
//
// Returns: true if substr is found within s, false otherwise.
// Example:
//
//	fmt.Println(ContainsFold("GoLang", "golang")) // Output: true
func ContainsFold(s string, r string) bool {
	return IndexFold(s, r) >= 0
}

// HasSuffixFold reports whether the string s ends with suffix r,
// using a case-insensitive comparison. It returns true if r is
// an empty string, or if s ends with r (case-insensitively).
//
// Parameters:
//   - s: The string to be checked.
//   - r: The suffix to be searched for.
//
// Returns:
//   - bool: true if s ends with r (case-insensitively), false otherwise.
func HasSuffixFold(s string, r string) bool {
	sl := len(s)
	rl := len(r)
	if rl == 0 {
		return true
	}

	if sl < rl {
		return false
	}

	return strings.EqualFold(s[sl-rl:], r)
}

// HasPrefixFold reports whether the string s begins with prefix r,
// using a case-insensitive comparison. It returns true if r is an
// empty string or if s starts with r (case-insensitively), and false otherwise.
//
// Parameters:
//   - s: The string to be checked.
//   - r: The prefix to be looked for.
//
// Returns:
//   - bool: true if s starts with r (case-insensitively) or if r is an empty string, false otherwise.
func HasPrefixFold(s string, r string) bool {
	sl := len(s)
	rl := len(r)
	if rl == 0 {
		return true
	}

	if sl < rl {
		return false
	}

	return strings.EqualFold(s[:rl], r)
}

// IndexFold returns the index of the first instance of the substring r in s,
// ignoring case. If r is not present in s, it returns -1. If r is an empty
// string, it returns 0.
//
// Parameters:
// - s: The string to search within.
// - r: The substring to search for.
//
// Returns:
//   - The index of the first occurrence of r in s, ignoring case, or -1 if r is
//     not present in s.
func IndexFold(s string, r string) int {
	sl := len(s)
	rl := len(r)
	if rl == 0 {
		return 0
	}

	runes := []rune(s)

	for i := 0; i < sl; i++ {
		if i+(rl) > sl {
			return -1
		}

		for j, y := range r {
			x := runes[i+j]
			if x == y {
				if j == rl-1 {
					return i
				}

				continue
			}

			if unicode.IsLetter(x) {
				if equalFoldRune(x, y) {
					if j == rl-1 {
						return i
					}

					continue
				}
			}

			break
		}
	}

	return -1
}

// IsSpace checks if all characters in the given string are whitespace characters.
// It returns true if all characters are whitespace, and false otherwise.
//
// Parameters:
//
//	s - the input string to check.
//
// Returns:
//
//	bool - true if all characters in the string are whitespace, false otherwise.
func IsSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

// IsEmpty checks if the given string is empty.
// It returns true if the string has a length of zero, otherwise it returns false.
// Parameters:
//   - s: The string to check.
//
// Returns:
//   - bool: true if the string is empty, false otherwise.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsEmptySpace checks if the given string is either empty or consists solely of whitespace characters.
// It returns true if the string is empty or contains only whitespace, and false otherwise.
// Parameters:
//   - s: The string to check.
//
// Returns:
//   - bool: true if the string is empty or contains only whitespace characters, false otherwise.
func IsEmptySpace(s string) bool {
	return IsEmpty(s) || IsSpace(s)
}

func equalFoldRune(x, y rune) bool {
	xx := unicode.SimpleFold(x)
	if xx == y {
		return true
	}
	yy := unicode.SimpleFold(y)
	return yy == x
}
