// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

import "sort"

type key []rune

func (a key) Less(i, j int) bool { return a[i] < a[j] }
func (a key) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a key) Len() int           { return len(a) }

// GroupAnagrams divide words to anagram groups and returns result.
// The words that are not in any group are omitted from result.
// The time complexity is O(n*m*log(m)) where n is the number of
// words and m is the maximum string length. The O(n) additional
// space is needed (beyond the space needed to write the final result).
func GroupAnagrams(words []string) (ag [][]string) {
	if len(words) == 0 {
		return ag
	}

	a := make(map[string][]string)
	for _, w := range words {
		k := key(w)
		sort.Sort(k)
		h := string(k)
		a[h] = append(a[h], w)
	}

	for _, g := range a {
		if len(g) > 1 {
			ag = append(ag, g)
		}
	}
	return ag
}
