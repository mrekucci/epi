// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package graphs

import "testing"

func TestIsMinimallyConnected(t *testing.T) {
	a, b, c, d, e, f, g, h, i, j, k, l, m := &vertex{label: "a"}, &vertex{label: "b"}, &vertex{label: "c"}, &vertex{label: "d"}, &vertex{label: "e"}, &vertex{label: "f"}, &vertex{label: "g"}, &vertex{label: "h"}, &vertex{label: "i"}, &vertex{label: "j"}, &vertex{label: "k"}, &vertex{label: "l"}, &vertex{label: "m"}
	a.edges = []*vertex{b}
	b.edges = []*vertex{a, e}
	c.edges = []*vertex{d}
	d.edges = []*vertex{c, e}
	e.edges = []*vertex{b, d, h, f}
	f.edges = []*vertex{e, g, i}
	g.edges = []*vertex{f}
	h.edges = []*vertex{e}
	i.edges = []*vertex{f, j, m}
	m.edges = []*vertex{i, k}
	k.edges = []*vertex{m, l}
	l.edges = []*vertex{k}
	x, y, z := &vertex{label: "x"}, &vertex{label: "y"}, &vertex{label: "z"}
	x.edges = []*vertex{y, z}
	y.edges = []*vertex{x, z}
	z.edges = []*vertex{x, y}
	for _, test := range []struct {
		in   []*vertex
		want bool
	}{
		{[]*vertex(nil), true},
		{[]*vertex{x, y, z}, false},
		{[]*vertex{a, b, c, d, e, f, g, h, i, j, k, l, m}, true},
	} {
		if got := IsMinimallyConnected(test.in); got != test.want {
			t.Errorf("IsMinimallyConnected(%v) = %t; want %t", test.in, got, test.want)
		}
	}
}

func BenchmarkIsMinimallyConnected(b *testing.B) {
	a, c, d, e, f, g, h, i, j, k, l, m := &vertex{label: "a"}, &vertex{label: "c"}, &vertex{label: "d"}, &vertex{label: "e"}, &vertex{label: "f"}, &vertex{label: "g"}, &vertex{label: "h"}, &vertex{label: "i"}, &vertex{label: "j"}, &vertex{label: "k"}, &vertex{label: "l"}, &vertex{label: "m"}
	a.edges = []*vertex{e}
	c.edges = []*vertex{d}
	d.edges = []*vertex{c, e}
	e.edges = []*vertex{a, d, h, f}
	f.edges = []*vertex{e, g, i}
	g.edges = []*vertex{f}
	h.edges = []*vertex{e}
	i.edges = []*vertex{f, j, m}
	m.edges = []*vertex{i, k}
	k.edges = []*vertex{m, l}
	l.edges = []*vertex{k}
	graph := []*vertex{a, c, d, e, f, g, h, i, j, k, l, m}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsMinimallyConnected(graph)
	}
}
