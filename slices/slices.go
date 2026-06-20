package slices

import (
	"github.com/adagit94/t"
	"slices"
)

func Map[S ~[]E, E any, EE any](s S, f func(e E, i int) EE) []EE {
	ss := make([]EE, len(s))

	for i, e := range s {
		ss[i] = f(e, i)
	}

	return ss
}

func Difference[S ~[]E, E comparable](s1 S, s2 S) S {
	s := make(S, 0)

	for _, v := range s1 {
		if !slices.Contains(s2, v) {
			s = append(s, v)
		}
	}

	return s
}

func Flat[S ~[][]E, E any](s S) []E {
	ss := make([]E, 0)

	for _, e := range s {
		for _, ee := range e {
			ss = append(ss, ee)
		}
	}

	return ss
}

func ExpandRange[R [2]N, N t.Int](r R) []N {
	rr := make([]N, 0, r[1]-r[0]+1)

	for n := r[0]; n <= r[1]; n++ {
		rr = append(rr, n)
	}

	return rr
}

func ExpandRanges[Rs []R, R [2]N, N t.Int](rs Rs) [][]N {
	rrs := make([][]N, 0, len(rs))

	for _, r := range rs {
		rrs = append(rrs, ExpandRange(r))
	}

	return rrs
}

func ExpandRangesFlat[Rs []R, R [2]N, N t.Int](rs Rs) []N {
	return Flat(ExpandRanges(rs))
}
