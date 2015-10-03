// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

import "testing"

func TestStringToInt(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int64
		err  error
	}{
		{"", 0, ErrSyntax},
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
		{"9223372036854775808", 0, ErrRange},
		{"-9223372036854775808", -1 << 63, nil},
		{"9223372036854775809", 0, ErrRange},
		{"-9223372036854775809", 0, ErrRange},
		{"9223372036854775810", 0, ErrRange},
		{"-9223372036854775810", 0, ErrRange},
		{"a", 0, ErrSyntax},
		{"-a", 0, ErrSyntax},
		{"123a", 0, ErrSyntax},
		{"-123a", 0, ErrSyntax},
	} {
		if got, err := StringToInt(test.in); got != test.want || err != test.err {
			t.Errorf("StringToInt(%q) = %d, %v; want %d, %v", test.in, got, err, test.want, test.err)
		}
	}
}

func BenchmarkStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToInt("9223372036854775807")
	}
}

func TestIntToString(t *testing.T) {
	for _, test := range []struct {
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
	} {
		if got := IntToString(test.in); got != test.want {
			t.Errorf("IntToString(%d) = %q; want %q", test.in, got, test.want)
		}
	}
}

func BenchmarkIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToString(int64(i))
	}
}
