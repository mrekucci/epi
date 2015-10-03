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

// move transfers n rings from p0 to p1 as follows:
// - Recursively move n-1 rings from p0 to p2 using p1 (this leaves disc n alone on p0).
// - Move the ring numbered n-1 from p0 to p1.
// - Recursively move the n-1 rings on p2 to p1, using p0 (so they sit on disc n).
// E.g.:
//
//	step    cnt    p0        p1        p2
//	---------------------------------------
// 	   -    0     [2,1,0] | []      | []
// 	p0->p1  1     [2,1]   | [0]     | []
// 	p0->p2  2     [2]     | [0]     | [1]
// 	p1->p2  3     [2]     | []      | [1,0]
// 	p0->p1  4     []      | [2]     | [1,0]
// 	p2->p0  5     [0]     | [2]     | [1]
// 	p2->p1  6     [0]     | [2,1]   | []
// 	p0->p1  7     []      | [2,1,0] | []
//
// Every step is written together with step count and error to r.
func move(p [3]*stacks.IntStack, n, s, d, m int, r *recorder) {
	if n > 0 {
		move(p, n-1, s, m, d, r)
		p[d].Push(p[s].Pop())
		fmt.Fprintf(&r.buf, "{p%d->p%d}", s, d) // We don't check for error 'cause writing to bytes.Buffer has always nil err.
		move(p, n-1, m, d, s, r)
		r.cnt++
	}
}

// HanoiSteps records and returns number of steps needed for moving n rings from
// one peg to another peg with using one help peg. The rings on every peg must
// be in ordered ascending. The number of steps is defined by function: cnt = 2**n - 1.
// If the n is bigger then size of int, then zero values and false is returned.
func HanoiSteps(n int) (cnt int, steps string, ok bool) {
	if n > intSize {
		return 0, "", false
	}
	p := [...]*stacks.IntStack{new(stacks.IntStack), new(stacks.IntStack), new(stacks.IntStack)}
	for i := n - 1; i >= 0; i-- {
		p[0].Push(interface{}(i))
	}

	var r recorder
	move(p, n, 0, 1, 2, &r)
	return r.cnt, r.buf.String(), true
}
