// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package btrees

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	A := &PBTree{Data: "A"}
	B := &PBTree{Data: "B"}
	C := &PBTree{Data: "C"}
	D := &PBTree{Data: "D"}
	E := &PBTree{Data: "E"}
	F := &PBTree{Data: "F"}
	G := &PBTree{Data: "G"}
	H := &PBTree{Data: "H"}
	I := &PBTree{Data: "I"}
	J := &PBTree{Data: "J"}
	K := &PBTree{Data: "K"}
	L := &PBTree{Data: "L"}
	M := &PBTree{Data: "M"}
	N := &PBTree{Data: "N"}
	O := &PBTree{Data: "O"}
	P := &PBTree{Data: "P"}

	// A Binary Tree of height 5 with root in "A".
	A.parent, A.left, A.right = nil, B, I
	// A left subtree:
	B.parent, B.left, B.right = A, C, F
	C.parent, C.left, C.right = B, D, E
	D.parent, E.parent = C, C
	F.parent, F.left, F.right = B, nil, G
	G.parent, G.left, G.right = F, H, nil
	H.parent = G
	// A right subtree:
	I.parent, I.left, I.right = A, J, O
	J.parent, J.left, J.right = I, nil, K
	K.parent, K.left, K.right = J, L, N
	N.parent = K
	L.parent, L.left, L.right = K, nil, M
	M.parent = L
	O.parent, O.left, O.right = I, nil, P
	P.parent = O

	want := []interface{}{"D", "C", "E", "B", "F", "H", "G", "A", "J", "L", "M", "K", "N", "I", "O", "P"}
	got := InorderTraversal(A)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InorderTraversal got %v; want %v", got, want)
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	A := &PBTree{Data: "A"}
	B := &PBTree{Data: "B"}
	C := &PBTree{Data: "C"}
	D := &PBTree{Data: "D"}
	E := &PBTree{Data: "E"}
	F := &PBTree{Data: "F"}
	G := &PBTree{Data: "G"}
	H := &PBTree{Data: "H"}
	I := &PBTree{Data: "I"}
	J := &PBTree{Data: "J"}
	K := &PBTree{Data: "K"}
	L := &PBTree{Data: "L"}
	M := &PBTree{Data: "M"}
	N := &PBTree{Data: "N"}
	O := &PBTree{Data: "O"}
	P := &PBTree{Data: "P"}

	// A Binary Tree of height 5 with root in "A".
	A.parent, A.left, A.right = nil, B, I
	// A left subtree:
	B.parent, B.left, B.right = A, C, F
	C.parent, C.left, C.right = B, D, E
	D.parent, E.parent = C, C
	F.parent, F.left, F.right = B, nil, G
	G.parent, G.left, G.right = F, H, nil
	H.parent = G
	// A right subtree:
	I.parent, I.left, I.right = A, J, O
	J.parent, J.left, J.right = I, nil, K
	K.parent, K.left, K.right = J, L, N
	N.parent = K
	L.parent, L.left, L.right = K, nil, M
	M.parent = L
	O.parent, O.left, O.right = I, nil, P
	P.parent = O

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InorderTraversal(A)
	}
}
