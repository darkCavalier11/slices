package slices

import (
	"log"
)

const (
	EMPTY_SLICE_ERROR = "The slice is empty"
)

// Checks if the slice is empty
func IsEmpty[T any](S []T) bool {
	return len(S) == 0
}

// returns the first element of the slice. if the slice has 0 Elements, panic.
func First[T any](S []T) T {
	if IsEmpty(S) {
		panic(EMPTY_SLICE_ERROR)
	}
	return S[0]
}

// returns the last element of the slice. If the slice has 0 Elements, panic.
func Last[T any](S []T) T {
	if IsEmpty(S) {
		panic(EMPTY_SLICE_ERROR)
	}
	return S[len(S)-1]
}

// Pops and returns the last element of the slice. if the slice has no Elements panic.
func Pop[T any](S *[]T) T {
	if IsEmpty(*S) {
		panic(EMPTY_SLICE_ERROR)
	}
	lastElement := Last(*S)
	newSlice := (*S)[:len(*S)-1]
	*S = newSlice
	log.Println(S)
	return lastElement
}

// appends all the Elements at the end of the slice
func AppendAll[T any](S []T, Elements []T) []T {
	for _, e := range Elements {
		S = append(S, e)
	}
	return S
}

// returns true if any one of the array elements pass the
// predicate function and returns true
func Any[T any](S []T, PredicateFunc func(T) bool) bool {
	for _, e := range S {
		if PredicateFunc(e) {
			return true
		}
	}
	return false
}

// returns true if all of the array elements pass the
// predicate function and returns true
func Every[T any](S []T, PredicateFunc func(T) bool) bool {
	for _, e := range S {
		if !PredicateFunc(e) {
			return false
		}
	}
	return true
}

// clears the entire slice
func Clear[T any](S *[]T) {
	*S = []T{}
}


// returns true if the slice contains the element
func Contains[T comparable](S []T, element T) bool {
	for i := range S {
		if S[i] == element {
			return true
		}
	}
	return false
}

// returns index of first element where element satisfy the predicate function. if 
// non of the element pass the predicate returns -1
func FirstWhere[T any](S []T, PredicateFunc func (T) bool) int {
	for i, e := range S {
		if PredicateFunc(e) {
			return i
		}
	}
	return -1
}