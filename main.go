package slices

import (
	"fmt"
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

// returns index where element satisfy the predicate function. if
// non of the element pass the predicate returns -1
func IndexWhere[T any](S []T, PredicateFunc func(T) bool) int {
	for i, e := range S {
		if PredicateFunc(e) {
			return i
		}
	}
	return -1
}

// returns index from behind where element satisfy the predicate function. if
// non of the element pass the predicate returns -1
func LastIndexWhere[T any](S []T, PredicateFunc func(T) bool) int {
	i := len(S)-1
	for i >= 0 {
		if PredicateFunc(S[i]) {
			return i
		}
		i--
	}
	return -1
}

// ForEach performs certain action on individual elements
func ForEach[T any](S []T, Action func(T)) {
	for _, e := range S {
		Action(e)
	}
}

// get a slice pointer that points to range of elements
// having index i, begin ≤ i < end. required: begin < end
func GetRange[T any](S []T, begin int, end int) *[]T {
	if begin < 0 {
		panic(fmt.Sprintf("begin should be >= 0, here %v", begin))
	} else if end > len(S) {
		panic(fmt.Sprintf("end should be less than length of slice: len(S): %v, end: %v", len(S), end))
	} else if begin >= end {
		panic(fmt.Sprintf("begin should be < end, here begin: %v, end: %v", begin, end))
	}
	rangeSlice := []T{}
	for begin < end {
		rangeSlice = append(rangeSlice, S[begin])
		begin++
	}
	return &rangeSlice
}

// Insert an element at particular index. If the index > len(s)
// it will panic
func Insert[T any](S *[]T, index int, element T) {
	if index > len(*S) || index < 0 {
		panic(fmt.Sprintf("invalid index to insert, len(S): %v, index: %v", len(*S), index))
	}
	if index == len(*S) {
		(*S) = append((*S), element)
		return
	}
	nextElement := (*S)[index]
	(*S)[index] = element
	index++
	for index < len(*S) {
		temp := (*S)[index]
		(*S)[index] = nextElement
		nextElement = temp
		index++
	}
	*S = append(*S, nextElement)
}


