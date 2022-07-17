package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testSlices = [][]int{
	{4, 5, 6},
	{},
	{78, 5874854, 56, 39},
	{4, 4, 4},
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
			require.Equal(t, len(resSlice), len(a)+len(s))
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

func TestSlice_Every(t *testing.T) {
	for _, s := range testSlices {
		// test for any element > 4
		elementGreaterThan4 := 0
		i := 0
		for i < len(s) {
			if s[i] > 4 {
				elementGreaterThan4++
			}
			i++
		}
		result := Every(s, func(e int) bool {
			return e > 4
		})
		if elementGreaterThan4 == len(s) {
			require.True(t, result, "invalid result for every")
		} else {
			require.False(t, result, "invalid result for every")
		}
	}
}

func TestSlice_Clear(t *testing.T) {
	for _, s := range testSlices {
		Clear(&s)
		require.Equal(t, 0, len(s), "invalid result in clearing")
	}
}

func TestSlice_Contains(t *testing.T) {
	for _, s := range testSlices {
		result := Contains(s, -1)
		require.False(t, result)

		if !IsEmpty(s) {
			result = Contains(s, s[0])
			require.True(t, result)
		}
	}
}

func TestSlice_FirstWhere(t *testing.T) {
	for _, s := range testSlices {
		result := FirstWhere(s, func(x int) bool {
			return x == 4
		})
		indexWhereFirst4 := -1
		for i := range s {
			if s[i] == 4 {
				indexWhereFirst4 = i
				break
			}
		}
		require.Equal(t, indexWhereFirst4, result)
	}
}

func TestSlice_ForEach(t *testing.T) {
	for _, s := range testSlices {
		var modifiedSlice = []int{}
		ForEach(s, func(t int) {
			modifiedSlice = append(modifiedSlice, t/2)
		})
		for i := range s {
			require.Equal(t, s[i]/2, modifiedSlice[i])
		}
	}
}

func TestSlice_GetRange(t *testing.T) {
	for _, s := range testSlices {
		require.Panicsf(t, func() {
			GetRange(s, -1, len(s))
		}, "invalid output for begin -1")
		require.Panicsf(t, func() {
			GetRange(s, 0, len(s)+1)
		}, "invalid output for end len(s) + 1")
		require.Panicsf(t, func() {
			GetRange(s, len(s), len(s))
		}, "invalid output for begin == end")
		require.Panicsf(t, func() {
			GetRange(s, 0, 0)
		}, "invalid output for begin == end")
		if len(s) != 0 {
			rangeSlice := GetRange(s, 0, len(s))
			require.Equal(t, *rangeSlice, s, "invalid output")

			rangeSlice = GetRange(s, 0, 1)
			require.Equal(t, *rangeSlice, []int{First(s)}, "invalid output")

			rangeSlice = GetRange(s, len(s)-1, len(s))
			require.Equal(t, *rangeSlice, []int{Last(s)}, "invalid output")
		}
	}
}
