package slices

import (
	"github.com/adagit94/t"
	"slices"
)

// Map maps S passed as argument with elements of type E to different type EE returned by the passed function. len(E) == len(EE) and new slice with underlaying array is allocated without mutation of original one.
func Map[S ~[]E, E any, EE any](s S, f func(e E, i int) EE) []EE {
	ss := make([]EE, len(s))

	for i, e := range s {
		ss[i] = f(e, i)
	}

	return ss
}

// Difference returns new slice with elements from s1 that aren't found in s2. It uses slices.Contains behind the scenes to compare element values.
func Difference[S ~[]E, E comparable](s1 S, s2 S) S {
	s := make(S, 0)

	for _, v := range s1 {
		if !slices.Contains(s2, v) {
			s = append(s, v)
		}
	}

	return s
}

// Flat flattens elements of nested slices of 2-dim slice into new single dimensional one (1 level deep).
func Flat[S ~[][]E, E any](s S) []E {
	ss := make([]E, 0)

	for _, e := range s {
		for _, ee := range e {
			ss = append(ss, ee)
		}
	}

	return ss
}

// ExpandRange expands passed range (from, to) and returns slice with intermediary values included. It's inclusive of destination (to) value. E.g. [1, 4] returns [1, 2, 3, 4].
func ExpandRange[N t.Int](from N, to N) []N {
	rr := make([]N, 0, to-from+1)

	for n := from; n <= to; n++ {
		rr = append(rr, n)
	}

	return rr
}

// ExpandRanges expands based on passed range bounds in form [from, to] and returns slices with intermediary values included for each range. It's inclusive of destination (to) values. It's just a plural form of ExpandRange with support for multiple of them. E.g. [2, 4], [6, 8] returns [[2, 3, 4], [6, 7, 8]].
func ExpandRanges[R [2]N, N t.Int](rs ...R) [][]N {
	rrs := make([][]N, 0, len(rs))

	for _, r := range rs {
		rrs = append(rrs, ExpandRange(r[0], r[1]))
	}

	return rrs
}

// ExpandRangesFlat expands based on passed range bounds in form [from, to] with inclusion of intermediary values and then flattens the slice, so expanded slices form just single 1D resulting slice. It's inclusive of destination (to) values. E.g. [[2, 4], [6, 8]] returns [2, 3, 4, 6, 7, 8].
func ExpandRangesFlat[R [2]N, N t.Int](rs ...R) []N {
	return Flat(ExpandRanges(rs...))
}

// Replace elements from a to b (exclusive) with passed compatible value e and return modified slice.
func ReplaceRepeated[S ~[]E, E any](s S, e E, a, b int) S {
	for i := a; i < b; i++ {
		s[i] = e
	}

	return s
}

// Replace every element in s with passed compatible value e and return modified slice.
func ReplaceRepeatedEvery[S ~[]E, E any](s S, e E) S {
	for i := range len(s) {
		s[i] = e
	}

	return s
}

// Creates and fills new slice based on passed capacity with element value.
func NewFilled[E any](element E, capacity int) []E {
	s := make([]E, 0, capacity)

	for range cap(s) {
		s = append(s, element)
	}

	return s
}
