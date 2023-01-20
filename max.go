package slicefuncs

import (
	"fmt"
)

// These functions are implemented based on their individual types instead of as generics to avoid having to create a type parameter that includes all Numbers except complex64 and complex128.
// Also because implementation details vary by type.

func MaxUint(s []uint) (uint, error) {
	var max uint
	if len(s) < 1 {
		return max, fmt.Errorf("empty slice")
	}
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max, nil
}

func MaxUint64(s []uint64) (uint64, error) {
	var max uint64
	if len(s) < 1 {
		return max, fmt.Errorf("empty slice")
	}
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max, nil
}

func MaxInt(s []int) (int, error) {
	var max int
	if len(s) < 1 {
		return max, fmt.Errorf("empty slice")
	}
	max = s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max, nil
}

func MaxInt64(s []int64) (int64, error) {
	var max int64
	if len(s) < 1 {
		return max, fmt.Errorf("empty slice")
	}
	max = s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max, nil
}

func MaxFloat64(s []float64) (float64, error) {
	var max float64
	if len(s) < 1 {
		return max, fmt.Errorf("empty slice")
	}
	max = s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max, nil
}
