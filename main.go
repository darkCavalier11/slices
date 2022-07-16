package slices

import "log"

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

// returns true if any of the array elements pass the
// predicate function and returns true
func Any[T any](S []T, PredicateFunc func(T) bool) bool {
	for _, e := range S {
		if PredicateFunc(e) {
			return true
		}
	}
	return false
}

func Clear[T any](S *[]T) {
	*S = []T{}
}