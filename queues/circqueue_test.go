// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import "testing"

func TestArrayQueue(t *testing.T) {
	testQueueInterface(t, new(arrayQueue), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, minInt, maxInt, 1.5, 1 + 1.5i, "x"})
}

func BenchmarkArrayQueue(b *testing.B) {
	q := new(arrayQueue)
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
