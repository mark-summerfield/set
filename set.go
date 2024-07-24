// Copyright © 2024 Mark Summerfield. All rights reserved.
package set

import (
	"fmt"
	"iter"
	"maps"
	"strings"
)

type Set[E comparable] struct{ set map[E]struct{} }

// New returns a new set containing the given elements (if any).
// If no elements are given, the type must be specified since it can't be
// inferred.
func New[E comparable](elements ...E) Set[E] {
	set := Set[E]{make(map[E]struct{}, len(elements))}
	if len(elements) > 0 {
		set.Add(elements...)
	}
	return set
}

// Add adds the given element(s) to the set.
func (me *Set[E]) Add(elements ...E) {
	for _, element := range elements {
		me.set[element] = struct{}{}
	}
}

// Delete deletes the given element(s) from the set.
func (me *Set[E]) Delete(elements ...E) {
	for _, element := range elements {
		delete(me.set, element)
	}
}

// Clear deletes all the elements to make this an empty set.
func (me *Set[E]) Clear() { clear(me.set) }

// Len returns the number of elements in the set.
func (me *Set[E]) Len() int { return len(me.set) }

// IsEmpty returns true if there are no elements in the set; otherwise
// returns false.
func (me *Set[E]) IsEmpty() bool { return len(me.set) == 0 }

// Contains returns true if element is in the set; otherwise returns false.
// Alternatively, use map syntax.
func (me *Set[E]) Contains(element E) bool {
	_, ok := me.set[element]
	return ok
}

// Difference returns a new set that contains the elements which are in this
// set that are not in the other set.
func (me *Set[E]) Difference(other Set[E]) Set[E] {
	diff := New[E]()
	for element := range me.set {
		if _, ok := other.set[element]; !ok {
			diff.set[element] = struct{}{}
		}
	}
	return diff
}

// SymmetricDifference returns a new set that contains the elements which
// are in this set or the other set—but not in both sets.
func (me *Set[E]) SymmetricDifference(other Set[E]) Set[E] {
	diff := New[E]()
	for element := range me.set {
		if _, ok := other.set[element]; !ok {
			diff.set[element] = struct{}{}
		}
	}
	for element := range other.set {
		if _, ok := me.set[element]; !ok {
			diff.set[element] = struct{}{}
		}
	}
	return diff
}

// Intersection returns a new set that contains the elements this set has in
// common with the other set.
func (me *Set[E]) Intersection(other Set[E]) Set[E] {
	intersection := New[E]()
	for element := range me.set {
		if _, ok := other.set[element]; ok {
			intersection.set[element] = struct{}{}
		}
	}
	return intersection
}

// Union returns a new set that contains the elements from this set and from
// the other set (with no duplicates of course).
// See also [Set.Unite].
func (me *Set[E]) Union(other Set[E]) Set[E] {
	union := me.Clone()
	union.Unite(other)
	return union
}

// Unite adds all the elements from other that aren't already in this set to
// this set.
// See also [Set.Union].
func (me *Set[E]) Unite(other Set[E]) {
	for element := range other.set {
		me.set[element] = struct{}{}
	}
}

// Clone returns a copy of this set.
func (me *Set[E]) Clone() Set[E] {
	return Set[E]{maps.Clone(me.set)}
}

// Equal returns true if this set has the same elements as the other set;
// otherwise returns false.
func (me *Set[E]) Equal(other Set[E]) bool {
	return maps.Equal(me.set, other.set)
}

// IsDisjoint returns true if this set has no elements in common with the
// other set; otherwise returns false.
func (me *Set[E]) IsDisjoint(other Set[E]) bool {
	for element := range me.set {
		if _, ok := other.set[element]; ok {
			return false
		}
	}
	return true
}

// IsSubsetOf returns true if this set is a subset of the other set, i.e.,
// if every member of this set is in the other set; otherwise returns false.
func (me *Set[E]) IsSubsetOf(other Set[E]) bool {
	for element := range me.set {
		if _, ok := other.set[element]; !ok {
			return false
		}
	}
	return true
}

// IsSupersetOf returns true if this set is a superset of the other set,
// i.e., if every member of the other set is in this set; otherwise returns
// false.
func (me Set[E]) IsSupersetOf(other Set[E]) bool {
	return other.IsSubsetOf(me)
}

// All returns an iterator, e.g., for element := range set.All() ...
func (me *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for key := range me.set {
			if !yield(key) {
				return
			}
		}
	}
}

// AllX returns an iterator, e.g.,
// for count, element := range set.AllX(1) ...
func (me *Set[E]) AllX(start ...int) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		i := 0
		if len(start) > 0 {
			i = start[0]
		}
		for key := range me.set {
			if !yield(i, key) {
				return
			}
			i++
		}
	}
}

// ToSlice returns this set's elements as an unsorted slice.
// For iteration either use this, or if you only need one value at a time,
// use [All] or [AllX]. To sort, use slices.Sorted (if E is cmp.Orderable).
func (me *Set[E]) ToSlice() []E {
	slice := make([]E, 0, len(me.set))
	for element := range me.set {
		slice = append(slice, element)
	}
	return slice
}

// String returns a human readable string representation of the set.
func (me *Set[E]) String() string {
	format := "%s%v"
	if me.hasStringElements() {
		format = "%s%q"
	}
	var out strings.Builder
	out.WriteByte('{')
	sep := ""
	for _, element := range me.ToSlice() {
		fmt.Fprintf(&out, format, sep, element)
		sep = " "
	}
	out.WriteByte('}')
	return out.String()
}

func (me *Set[E]) hasStringElements() bool {
	for key := range me.set {
		_, ok := any(key).(string)
		return ok
	}
	return false
}
