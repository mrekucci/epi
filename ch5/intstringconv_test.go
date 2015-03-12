package ch5

import "testing"

var stringToIntTests = []struct {
	in   string
	want int64
	err  error
}{
	{"", 0, errSyntax},
	{"0", 0, nil},
	{"-0", 0, nil},
	{"1", 1, nil},
	{"-1", -1, nil},
	{"+1", 1, nil},
	{"0123456789", 123456789, nil},
	{"-0123456789", -123456789, nil},
	{"123456789", 123456789, nil},
	{"-123456789", -123456789, nil},
	{"9223372036854775807", 1<<63 - 1, nil},
	{"-9223372036854775807", -(1<<63 - 1), nil},
	{"9223372036854775808", 0, errRange},
	{"-9223372036854775808", -1 << 63, nil},
	{"9223372036854775809", 0, errRange},
	{"-9223372036854775809", 0, errRange},
	{"9223372036854775810", 0, errRange},
	{"-9223372036854775810", 0, errRange},
	{"a", 0, errSyntax},
	{"-a", 0, errSyntax},
	{"123a", 0, errSyntax},
	{"-123a", 0, errSyntax},
}

func TestStringToInt(t *testing.T) {
	for _, tt := range stringToIntTests {
		got, err := stringToInt(tt.in)
		if got != tt.want || err != tt.err {
			t.Errorf("stringToInt(%q) = %d, %v; want %d, %v", tt.in, got, err, tt.want, tt.err)
		}
	}
}

func BenchmarkStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringToInt("9223372036854775807")
	}
}

var intToStringTests = []struct {
	in   int64
	want string
}{
	{0, "0"},
	{1, "1"},
	{-1, "-1"},
	{123456789, "123456789"},
	{-123456789, "-123456789"},
	{1<<63 - 1, "9223372036854775807"},
	{-(1<<63 - 1), "-9223372036854775807"},
	{-1 << 63, "-9223372036854775808"},
}

func TestIntToString(t *testing.T) {
	for _, tt := range intToStringTests {
		got := intToString(tt.in)
		if got != tt.want {
			t.Errorf("intToString(%d) = %q; want %q", tt.in, got, tt.want)
		}
	}
}

func BenchmarkIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToString(int64(i))
	}
}
