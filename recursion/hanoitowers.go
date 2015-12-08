// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

import (
	"bytes"
	"fmt"

	"github.com/mrekucci/epi/stacks"
)

// recorder records the steps and the number of steps during the rings movement.
type recorder struct {
	buf bytes.Buffer
	cnt int
}

// move transfers n rings from src to dst using aux as follows:
// - Recursively move n-1 rings from src to aux using dst (this leaves disc n alone on src).
// - Move the ring numbered n from src to dst.
// - Recursively move the n-1 rings from aux to dst, using src (so they sit on disc n).
// E.g.:
//
//	step      cnt   src      dst       aux
//	---------------------------------------
// 	   -      0     [2,1,0] | []      | []
// 	src->dst  1     [2,1]   | [0]     | []
// 	src->aux  2     [2]     | [0]     | [1]
// 	dst->aux  3     [2]     | []      | [1,0]
// 	src->dst  4     []      | [2]     | [1,0]
// 	aux->src  5     [0]     | [2]     | [1]
// 	aux->dst  6     [0]     | [2,1]   | []
// 	src->dst  7     []      | [2,1,0] | []
//
// Every step is written together with step count and error to r.
func move(p [3]*stacks.IntStack, n, src, dst, aux int, r *recorder) {
	if n > 0 {
		move(p, n-1, src, aux, dst, r)
		p[dst].Push(p[src].Pop())
		fmt.Fprintf(&r.buf, "{p%d->p%d}", src, dst) // We don't check for error 'cause writing to bytes.Buffer has always nil err.
		move(p, n-1, aux, dst, src, r)
		r.cnt++
	}
}

// HanoiSteps records and returns number of steps needed for moving n rings from
// one peg to another peg with using one help peg. The rings on every peg must
// be in ordered ascending. The number of steps is defined by function: cnt = 2**n - 1.
// If the n is bigger then size of int, then zero values and false is returned.
// The time complexity is O(2**n) and O(1) additional space is needed.
func HanoiSteps(n int) (cnt int, steps string, ok bool) {
	if n > intSize {
		return 0, "", false
	}
	p := [...]*stacks.IntStack{new(stacks.IntStack), new(stacks.IntStack), new(stacks.IntStack)}
	for i := n - 1; i >= 0; i-- {
		p[0].Push(interface{}(i))
	}

	r := new(recorder)
	move(p, n, 0, 1, 2, r)
	return r.cnt, r.buf.String(), true
}
