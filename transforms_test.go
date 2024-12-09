package xstrings

import (
	"testing"
)

func TestUnderscore(t *testing.T) {
	tests := []struct {
		input    string
		options  []HyphenMinusOption
		expected string
	}{
		{"HelloWorld", nil, "hello_world"},
		{"Hello World", nil, "hello_world"},
		{"Hello-World", nil, "hello_world"},
		{"Hello_World", nil, "hello_world"},
		{"HelloWorld", []HyphenMinusOption{Screaming}, "HELLO_WORLD"},
		{"HelloWorld", []HyphenMinusOption{func(p *HyphenMinusParams) { p.PreserveCase = true }}, "Hello_World"},
	}

	for _, test := range tests {
		result := Underscore(test.input, test.options...)
		if result != test.expected {
			t.Errorf("Underscore(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestDasherize(t *testing.T) {
	tests := []struct {
		input    string
		options  []HyphenMinusOption
		expected string
	}{
		{"HelloWorld", nil, "hello-world"},
		{"Hello World", nil, "hello-world"},
		{"Hello-World", nil, "hello-world"},
		{"Hello_World", nil, "hello-world"},
		{"HelloWorld", []HyphenMinusOption{Screaming}, "HELLO-WORLD"},
		{"HelloWorld", []HyphenMinusOption{func(p *HyphenMinusParams) { p.PreserveCase = true }}, "Hello-World"},
	}

	for _, test := range tests {
		result := Dasherize(test.input, test.options...)
		if result != test.expected {
			t.Errorf("Dasherize(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world_example", "helloWorldExample"},
		{"Hello_World_Example", "helloWorldExample"},
		{"helloWorldExample", "helloWorldExample"},
		{"HelloWorldExample", "helloWorldExample"},
	}

	for _, test := range tests {
		result := CamelCase(test.input)
		if result != test.expected {
			t.Errorf("CamelCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "HelloWorld"},
		{"Hello_World", "HelloWorld"},
		{"helloWorld", "HelloWorld"},
		{"HelloWorld", "HelloWorld"},
	}

	for _, test := range tests {
		result := PascalCase(test.input)
		if result != test.expected {
			t.Errorf("PascalCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
