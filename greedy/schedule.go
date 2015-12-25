// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package greedy

import "sort"

// MinWaitingTime returns a time that represents a sum of all the minimum
// waiting times ordered in a way in which to process queries minimizes time.
func MinWaitingTime(times []int) (waiting int) {
	sort.Ints(times)
	for i, t := range times { // Sum of all the waiting times.
		waiting += t * (len(times) - 1 - i) // Add a waiting time for the ith client.
	}
	return waiting
}
