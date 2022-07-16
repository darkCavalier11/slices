package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testSlices = [][]int{
	{4, 5, 6},
	{},
	{78, 5874854, 56, 39},
}

func TestSlice_IsEmpty(t *testing.T) {
	for _, s := range testSlices {
		if len(s) == 0 {
			require.True(t, IsEmpty(s), "invalid output")
		} else {
			require.False(t, IsEmpty(s), "invalid output")
		}
	}
}

func TestSlice_First(t *testing.T) {
	for _, s := range testSlices {
		if IsEmpty(s) {
			require.Panicsf(t, func() {
				First(s)
			}, "invalid output, should have panicked")
		} else {
			require.Equal(t, s[0], First(s), "invalid output")
		}
	}
}

func TestSlice_Last(t *testing.T) {
	for _, s := range testSlices {
		if IsEmpty(s) {
			require.Panicsf(t, func() {
				Last(s)
			}, "invalid output, should have panicked")
		} else {
			require.Equal(t, s[len(s)-1], Last(s), "invalid output")
		}
	}
}

func TestSlice_Pop(t *testing.T) {
	for _, s := range testSlices {
		if IsEmpty(s) {
			require.Panicsf(t, func() {
				Pop(&s)
			}, "invalid output, should have panicked")
		} else {
			expectedLastElement := s[len(s)-1]
			resLastElement := Pop(&s)
			require.Equal(t, expectedLastElement, resLastElement, "invalid last element")
		}
	}
}

func TestSlice_AppendAll(t *testing.T) {
	for _, s := range testSlices {
		for _, a := range testSlices {
			resSlice := AppendAll(s, a)
			require.Equal(t, len(resSlice), len(a) + len(s))
			i := 0
			for i < len(resSlice) {
				if i >= len(s) {
					require.Equal(t, resSlice[i], a[i-len(s)])
				} else {
					require.Equal(t, resSlice[i], s[i])
				}
				i++
			}
		}
	}
}

func TestSlice_Any(t *testing.T) {
	for _, s := range testSlices {
		// test for any element > 4
		isExist := false
		i := 0
		for i < len(s) {
			if s[i] > 4 {
				isExist = true
			}
			i++
		}
		result := Any(s, func(e int) bool {
			return e > 4
		})
		require.Equal(t, result, isExist)
	}
}

func TestSlice_Clear(t *testing.T) {
	for _, s := range testSlices {
		Clear(&s)
		require.Equal(t, 0, len(s), "invalid result in clearing")
	}
}