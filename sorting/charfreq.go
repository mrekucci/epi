// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import (
	"bytes"
	"fmt"
	"sort"
)

type sentence []rune

func (s sentence) Less(i, j int) bool { return s[i] < s[j] }
func (s sentence) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sentence) Len() int           { return len(s) }

// CountOccurrences counts and prints the occurrences
// of individual characters that appears in s.
// The time complexity is O(n*long(n)), and O(1) additional space is needed.
func CountOccurrences(s string) string {
	if len(s) == 0 {
		return ""
	}

	chars := sentence(s)
	sort.Sort(chars)
	cnt := 1
	buf := new(bytes.Buffer)
	for i := 1; i < len(chars); i++ {
		if chars[i-1] != chars[i] {
			fmt.Fprintf(buf, "(%c, %d), ", chars[i-1], cnt)
			cnt = 0
		}
		cnt++
	}
	fmt.Fprintf(buf, "(%c, %d)", chars[len(chars)-1], cnt)
	return buf.String()
}
