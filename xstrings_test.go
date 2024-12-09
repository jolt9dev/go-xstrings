package xstrings

import (
	"testing"
)

func TestContainsFold(t *testing.T) {
	tests := []struct {
		s, r string
		want bool
	}{
		{"Hello, World", "world", true},
		{"Hello, World", "WORLD", true},
		{"Hello, World", "earth", false},
	}

	for _, tt := range tests {
		if got := ContainsFold(tt.s, tt.r); got != tt.want {
			t.Errorf("ContainsFold(%q, %q) = %v; want %v", tt.s, tt.r, got, tt.want)
		}
	}
}

func TestHasSuffixFold(t *testing.T) {
	tests := []struct {
		s, r string
		want bool
	}{
		{"Hello, World", "world", true},
		{"Hello, World", "WORLD", true},
		{"Hello, World", "earth", false},
	}

	for _, tt := range tests {
		if got := HasSuffixFold(tt.s, tt.r); got != tt.want {
			t.Errorf("HasSuffixFold(%q, %q) = %v; want %v", tt.s, tt.r, got, tt.want)
		}
	}
}

func TestHasPrefixFold(t *testing.T) {
	tests := []struct {
		s, r string
		want bool
	}{
		{"Hello, World", "hello", true},
		{"Hello, World", "HELLO", true},
		{"Hello, World", "earth", false},
	}

	for _, tt := range tests {
		if got := HasPrefixFold(tt.s, tt.r); got != tt.want {
			t.Errorf("HasPrefixFold(%q, %q) = %v; want %v", tt.s, tt.r, got, tt.want)
		}
	}
}

func TestIndexFold(t *testing.T) {
	tests := []struct {
		s, r string
		want int
	}{
		{"Hello, World", "world", 7},
		{"Hello, World", "WORLD", 7},
		{"Hello, World", "earth", -1},
	}

	for _, tt := range tests {
		if got := IndexFold(tt.s, tt.r); got != tt.want {
			t.Errorf("IndexFold(%q, %q) = %v; want %v", tt.s, tt.r, got, tt.want)
		}
	}
}

func TestIsSpace(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{" ", true},
		{"\t", true},
		{"\n", true},
		{"Hello", false},
	}

	for _, tt := range tests {
		if got := IsSpace(tt.s); got != tt.want {
			t.Errorf("IsSpace(%q) = %v; want %v", tt.s, got, tt.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"Hello", false},
	}

	for _, tt := range tests {
		if got := IsEmpty(tt.s); got != tt.want {
			t.Errorf("IsEmpty(%q) = %v; want %v", tt.s, got, tt.want)
		}
	}
}

func TestIsEmptySpace(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", true},
		{" ", true},
		{"Hello", false},
	}

	for _, tt := range tests {
		if got := IsEmptySpace(tt.s); got != tt.want {
			t.Errorf("IsEmptySpace(%q) = %v; want %v", tt.s, got, tt.want)
		}
	}
}
