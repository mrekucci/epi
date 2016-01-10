// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import "testing"

func TestIntStackQueue(t *testing.T) {
	testQueueInterface(t, NewIntStackQueue(), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, minInt, maxInt})
}

func BenchmarkIntStackQueue(b *testing.B) {
	q := new(IntListQueue)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
		if i%64 == 0 {
			for i := 0; i < 16; i++ {
				q.Dequeue()
			}
		}
	}
}
