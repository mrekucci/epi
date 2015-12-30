// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package strings

// keypad maps keypad numbers to possible characters.
var keypad = [...][]byte{
	[]byte("0"),    // 0
	[]byte("1"),    // 1
	[]byte("ABC"),  // 2
	[]byte("DEF"),  // 3
	[]byte("GHI"),  // 4
	[]byte("JKL"),  // 5
	[]byte("MNO"),  // 6
	[]byte("PQRS"), // 7
	[]byte("TUV"),  // 8
	[]byte("WXYZ")} // 9

// enumMnemonics recursively enumerates and write into mnemonics all
// possible character sequences that corresponds to the number.
func enumMnemonics(number string, i int, m []byte, mnemonics []string) ([]string, bool) {
	if i == len(number) { // Base case. Since all the digits are processed, add copy of mnemonic m to mnemonics.
		return append(mnemonics, string(m)), true
	}

	// Check if character is a digit.
	d := number[i]
	if d < '0' || d > '9' {
		return nil, false
	}

	// Try all possible characters for digit d.
	for _, c := range keypad[int(d-'0')] {
		m[i] = c
		ms, ok := enumMnemonics(number, i+1, m, mnemonics)
		if !ok {
			return nil, false
		}
		mnemonics = ms
	}
	return mnemonics, true
}

// PhoneMnemonics returns all possible character
// sequences that corresponds to the number.
// The time complexity is O(4**n). The space complexity is O(4**n).
func PhoneMnemonics(number string) (mnemonics []string, ok bool) {
	if len(number) == 0 {
		return nil, false
	}
	return enumMnemonics(number, 0, make([]byte, len(number)), mnemonics)
}
