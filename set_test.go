// Copyright © 2024 Mark Summerfield. All rights reserved.
package set

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func check(act string, actSize int, exp string, expSize int, t *testing.T) {
	if actSize != expSize {
		t.Errorf("expected %d elements, got %d", expSize, actSize)
	}
	if exp != act {
		t.Errorf("expected %s, got %s", exp, act)
	}
}

func TestNew(t *testing.T) {
	s1 := New[int]()
	check(s1.String(), s1.Len(), "{}", 0, t)
	s2 := New(5)
	check(s2.String(), s2.Len(), "{5}", 1, t)
	s3 := New(50, 35, 78)
	check(s3.String(), s3.Len(), "{35 50 78}", 3, t)
	s4 := New("one", "two")
	check(s4.String(), s4.Len(), "{\"one\" \"two\"}", 2, t)
	a := New[int]()
	check(a.String(), a.Len(), "{}", 0, t)
	b := New("a string")
	check(b.String(), b.Len(), "{\"a string\"}", 1, t)
	c := New(19, 21, 1, 2, 4, 8)
	check(c.String(), c.Len(), "{1 2 4 8 19 21}", 6, t)
	s := []string{"A", "B", "C", "De", "Fgh"}
	d := New(s...)
	check(d.String(), d.Len(), "{\"A\" \"B\" \"C\" \"De\" \"Fgh\"}", len(s),
		t)
}

func TestToSlice(t *testing.T) {
	s := New(19, 21, 1, 2, 4, 8)
	u := s.ToSlice()
	sort.Ints(u)
	check(fmt.Sprintf("%v", u), len(u), "[1 2 4 8 19 21]", s.Len(), t)
}

func TestToSortedSlice(t *testing.T) {
	s := New(19, 21, 1, 7, 2, 4, 8, 0)
	u := s.ToSlice()
	slices.Sort(u)
	check(fmt.Sprintf("%v", u), len(u), "[0 1 2 4 7 8 19 21]", s.Len(), t)
}

func TestAdd(t *testing.T) {
	s := New(19, 21, 1, 2, 4, 8)
	s.Add(5, 7, 1, 19)
	check(s.String(), s.Len(), "{1 2 4 5 7 8 19 21}", 8, t)
}

func TestDelete(t *testing.T) {
	s := New(19, 21, 1, 2, 5, 4, 8, 9, 11, 13, 7)
	s.Delete(5, 7, 1, 19)
	check(s.String(), s.Len(), "{2 4 8 9 11 13 21}", 7, t)
}

func TestClear(t *testing.T) {
	s := New(19, 21, 1, 2, 5, 4, 8, 9, 11, 13, 7)
	s.Clear()
	check(s.String(), s.Len(), "{}", 0, t)
	s.Add(1, 2, 3)
	check(s.String(), s.Len(), "{1 2 3}", 3, t)
}

func TestContains(t *testing.T) {
	s := New(19, 21, 1, 2, 5, 4, 8, 9, 11, 13, 7)
	if !s.Contains(11) {
		t.Error("expected set to contain 11")
	}
	if s.Contains(23) {
		t.Error("expected set not to contain 23")
	}
}

func TestDifference(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8)
	d := s.Difference(u)
	check(d.String(), d.Len(), "{0 1 3 5 7 9}", 6, t)
	d = u.Difference(s)
	check(d.String(), d.Len(), "{}", 0, t)
}

func TestSymmetricDifference(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8)
	d := s.SymmetricDifference(u)
	check(d.String(), d.Len(), "{0 1 3 5 7 9}", 6, t)
	d = u.SymmetricDifference(s)
	check(d.String(), d.Len(), "{0 1 3 5 7 9}", 6, t)
}

func TestIntersection(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8)
	x := s.Intersection(u)
	check(x.String(), x.Len(), "{2 4 6 8}", 4, t)
	v := New(1, 3, 5)
	y := u.Intersection(v)
	check(y.String(), y.Len(), "{}", 0, t)
	z := v.Intersection(u)
	check(z.String(), z.Len(), "{}", 0, t)
}

func TestUnion(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := New(2, 4, 6, 8, 10, 12)
	x := s.Union(u)
	check(x.String(), x.Len(), "{0 1 2 3 4 5 6 7 8 9 10 12}", 12, t)
}

func TestUnite(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	s.Unite(New(2, 4, 6, 8, 10, 12))
	check(s.String(), s.Len(), "{0 1 2 3 4 5 6 7 8 9 10 12}", 12, t)
}

func TestClone(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 6, 7, 8, 9)
	u := s.Clone()
	u.Add(5)
	s.Add(5)
	check(s.String(), s.Len(), u.String(), u.Len(), t)
}

func TestEqual(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := s.Clone()
	if !s.Equal(u) {
		t.Errorf("%v != %v", s, u)
	}
	u.Add(-3)
	if s.Equal(u) {
		t.Errorf("%v == %v", s, u)
	}
}

func TestIsDisjoint(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := s.Clone()
	if s.IsDisjoint(u) {
		t.Error("unexpectedly disjoint")
	}
	w := New(10, 11, 12)
	if !u.IsDisjoint(w) {
		t.Error("unexpectedly not disjoint")
	}
}

func TestIsSubsetOf(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := s.Clone()
	if !s.IsSubsetOf(u) {
		t.Error("unexpectedly not subset")
	}
	w := New(10, 11, 12)
	if w.IsSubsetOf(s) {
		t.Error("unexpectedly a subset")
	}
	x := New(4, 6, 2)
	if !x.IsSubsetOf(s) {
		t.Error("unexpectedly not subset")
	}
}

func TestIsSupersetOf(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	u := s.Clone()
	if !s.IsSupersetOf(u) {
		t.Error("unexpectedly not superset")
	}
	w := New(10, 11, 12)
	if w.IsSupersetOf(s) {
		t.Error("unexpectedly a superset")
	}
	x := New(4, 6, 2)
	if x.IsSupersetOf(s) {
		t.Error("unexpectedly a superset")
	}
	if !s.IsSupersetOf(x) {
		t.Error("unexpectedly not a superset")
	}
}

func TestMap(t *testing.T) {
	s := New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	if !s.Contains(7) {
		t.Error("expected to contain 7")
	}
	v := s.ToSlice()
	slices.Sort(v)
	check(fmt.Sprintf("%v", v), len(v), "[0 1 2 3 4 5 6 7 8 9]", s.Len(), t)
	w := make([]int, 0, s.Len())
	for x := range v {
		w = append(w, x)
	}
	sort.Ints(w)
	check(fmt.Sprintf("%v", w), len(w), "[0 1 2 3 4 5 6 7 8 9]", s.Len(), t)
	if len(w) != 10 {
		t.Error("expected 10 elements")
	}
	s.Clear()
	if s.Len() != 0 {
		t.Error("expected no elements")
	}
}

func TestStringMaxElements(t *testing.T) {
	s := New[int]()
	for i := 0; i < 111; i++ {
		s.Add(i)
	}
	check(s.String(), s.Len(), "{…111 elements…}", s.Len(), t)
}

func TestAll(t *testing.T) {
	s := New(10, 20, 30, 40, 50, 60, 70, 80, 90)
	n := 0
	for v := range s.All() {
		n += v
	}
	if n != 450 {
		t.Errorf("expected 450, got %d", n)
	}
}

func TestAllX(t *testing.T) {
	s := New(10, 20, 30, 40, 50, 60, 70, 80, 90)
	n := 0
	for i, v := range s.AllX() {
		n += v + i
	}
	if n != 486 {
		t.Errorf("expected 486, got %d", n)
	}
	n = 0
	for i, v := range s.AllX(1) {
		n += v + i
	}
	if n != 495 {
		t.Errorf("expected 495, got %d", n)
	}
}

func TestEg(t *testing.T) {
	s := New(1, 2, 3, 4, 5, 6)
	d := s.Difference(New(2, 4))
	v := d.ToSlice()
	slices.Sort(v)
	check(fmt.Sprintf("%v", v), len(v), "[1 3 5 6]", d.Len(), t)
}
