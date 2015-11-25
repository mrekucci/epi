// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bstrees

// BSTree represents a binary search tree.
type BSTree struct {
	Data  interface{}
	left  *BSTree
	right *BSTree
}
