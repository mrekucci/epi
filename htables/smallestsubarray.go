// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package htables

// SubarrayCoveringSet returns the indices of a shortest subarray of
// the given paragraph that covers all the words in the keywords.
// The time complexity is O(n). The space complexity is O(k), where k
// is the size of the keywords set.
func SubarrayCoveringSet(paragraph []string, keywords map[string]struct{}) (start, end int) {
	start, end = -1, -1
	if len(keywords) == 0 {
		return start, end
	}

	var l, r int
	kwCnt := make(map[string]int)
	for r < len(paragraph) {
		// Iterate r index through paragraph until it reaches end
		// or every word in keywords haven't been found.
		for r < len(paragraph) && len(kwCnt) < len(keywords) {
			w := paragraph[r]
			if _, ok := keywords[w]; ok {
				kwCnt[w]++
			}
			r++
		}
		// Iterate l index through paragraph until it reaches right
		// index position or keywordsCnt does not have all keywords.
		for l < r && len(kwCnt) == len(keywords) {
			w := paragraph[l]
			if cnt, ok := kwCnt[w]; ok {
				if cnt == 1 {
					delete(kwCnt, w)
				} else {
					kwCnt[w]--
				}
				if start == -1 && end == -1 || r-1-l < end-start {
					start, end = l, r-1
				}
			}
			l++
		}
	}

	return start, end
}

// SubarrayCoveringSetNaive returns the indices of a shortest subarray of
// the given paragraph that covers all the words in the keywords.
// The time complexity is O(n*n). The space complexity is O(k), where k
// is the size of the keywords set.
func SubarrayCoveringSetNaive(paragraph []string, keywords map[string]struct{}) (start, end int) {
	start, end = -1, -1
	if len(keywords) == 0 {
		return start, end
	}

	for l := 0; l < len(paragraph); l++ {
		kwCnt := make(map[string]int)
		for r := l; r < len(paragraph) && (start == -1 && end == -1 || r-l < end-start); r++ {
			w := paragraph[r]
			if _, ok := keywords[w]; ok {
				kwCnt[w]++
			}
			if len(kwCnt) == len(keywords) {
				if start == -1 && end == -1 || r-l < end-start {
					start, end = l, r
				}
				break
			}
		}
	}
	return start, end
}
