package slicefuncs

import (
	"fmt"
)

// These functions are implemented based on their individual types instead of as generics to avoid having to create a type parameter that includes all Numbers except complex64 and complex128.
func MinUint(s []uint) (uint, error) {
	var min uint
	if len(s) < 1 {
		return min, fmt.Errorf("empty slice")
	}
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min, nil
}

func MinUint64(s []uint64) (uint64, error) {
	var min uint64
	if len(s) < 1 {
		return min, fmt.Errorf("empty slice")
	}
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min, nil
}

func MinInt(s []int) (int, error) {
	var min int
	if len(s) < 1 {
		return min, fmt.Errorf("empty slice")
	}
	min = s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min, nil
}

func MinInt64(s []int64) (int64, error) {
	var min int64
	if len(s) < 1 {
		return min, fmt.Errorf("empty slice")
	}
	min = s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min, nil
}

func MinFloat64(s []float64) (float64, error) {
	var min float64
	if len(s) < 1 {
		return min, fmt.Errorf("empty slice")
	}
	min = s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min, nil
}
