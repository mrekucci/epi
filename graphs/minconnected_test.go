// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

import "testing"

func TestIsMinimallyConnected(t *testing.T) {
	a, b, c, d, e, f, g, h, i, j, k, l, m := &Vertex{label: "a"}, &Vertex{label: "b"}, &Vertex{label: "c"}, &Vertex{label: "d"}, &Vertex{label: "e"}, &Vertex{label: "f"}, &Vertex{label: "g"}, &Vertex{label: "h"}, &Vertex{label: "i"}, &Vertex{label: "j"}, &Vertex{label: "k"}, &Vertex{label: "l"}, &Vertex{label: "m"}
	a.edges = []*Vertex{b}
	b.edges = []*Vertex{a, e}
	c.edges = []*Vertex{d}
	d.edges = []*Vertex{c, e}
	e.edges = []*Vertex{b, d, h, f}
	f.edges = []*Vertex{e, g, i}
	g.edges = []*Vertex{f}
	h.edges = []*Vertex{e}
	i.edges = []*Vertex{f, j, m}
	m.edges = []*Vertex{i, k}
	k.edges = []*Vertex{m, l}
	l.edges = []*Vertex{k}
	x, y, z := &Vertex{label: "x"}, &Vertex{label: "y"}, &Vertex{label: "z"}
	x.edges = []*Vertex{y, z}
	y.edges = []*Vertex{x, z}
	z.edges = []*Vertex{x, y}
	for _, test := range []struct {
		in   []*Vertex
		want bool
	}{
		{[]*Vertex(nil), true},
		{[]*Vertex{x, y, z}, false},
		{[]*Vertex{a, b, c, d, e, f, g, h, i, j, k, l, m}, true},
	} {
		if got := IsMinimallyConnected(test.in); got != test.want {
			t.Errorf("IsMinimallyConnected(%v) = %t; want %t", test.in, got, test.want)
		}
	}
}

func BenchmarkIsMinimallyConnected(b *testing.B) {
	a, c, d, e, f, g, h, i, j, k, l, m := &Vertex{label: "a"}, &Vertex{label: "c"}, &Vertex{label: "d"}, &Vertex{label: "e"}, &Vertex{label: "f"}, &Vertex{label: "g"}, &Vertex{label: "h"}, &Vertex{label: "i"}, &Vertex{label: "j"}, &Vertex{label: "k"}, &Vertex{label: "l"}, &Vertex{label: "m"}
	a.edges = []*Vertex{e}
	c.edges = []*Vertex{d}
	d.edges = []*Vertex{c, e}
	e.edges = []*Vertex{a, d, h, f}
	f.edges = []*Vertex{e, g, i}
	g.edges = []*Vertex{f}
	h.edges = []*Vertex{e}
	i.edges = []*Vertex{f, j, m}
	m.edges = []*Vertex{i, k}
	k.edges = []*Vertex{m, l}
	l.edges = []*Vertex{k}
	graph := []*Vertex{a, c, d, e, f, g, h, i, j, k, l, m}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsMinimallyConnected(graph)
	}
}
