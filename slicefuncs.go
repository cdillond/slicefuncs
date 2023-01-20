package slicefuncs

import (
	"fmt"
)

// Returns the index of the first occurence of `v` in `s` for built-in types that satisfy the `comparable` interface.
func IndexOf[C comparable](v C, s []C) (int, error) {
	for i := 0; i < len(s); i++ {
		if v == s[i] {
			return i, nil
		}
	}
	return -1, fmt.Errorf("not found")
}

// Returns the index of the first occurence of `v` in `q`. `equals` must return `true` when `a==b` and `false` when `a!=b`.
// This function is intended to accomodate types that cannot be used with the built-in comparison operators (and therefore do not work with `IndexOf`).
func IndexOfAny[T any](v T, s []T, equals func(a, b T) bool) (int, error) {
	for i := 0; i < len(s); i++ {
		if equals(v, s[i]) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("not found")
}

// Returns the result of `fn(s[i])` for all integer values of `i` between `0` and `len(s)`.
func MapTo[T, R any](s []T, fn func(t T) R) []R {
	res := make([]R, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = fn(s[i])
	}
	return res
}

// Returns a slice of all elements of `s` that satisfy the condtion determined by `fn`.
func Filter[T any](s []T, fn func(t T) bool) []T {
	// pre-allocate res
	res := make([]T, 0, len(s))
	for i := 0; i < len(s); i++ {
		if fn(s[i]) {
			res = append(res, s[i])
		}
	}
	// reduce cap(res) back to len(res) to save memory
	return res[:len(res):len(res)]
}

// Returns a slice of the accumulated results of `fn` applied to `s`.
// When `len(s) > 0`, `s[0]` will always be equal to the value in the initial position of the returned slice.
func Accumulate[T any](s []T, fn func(a, b T) T) []T {
	res := make([]T, len(s))
	if len(s) == 0 {
		return res
	}
	res[0] = s[0]
	for i := 1; i < len(s); i++ {
		res[i] = fn(res[i-1], s[i])
	}
	return res
}

// Returns a reversed copy of `s`.
func Reverse[T any](s []T) []T {
	res := make([]T, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[len(s)-1-i]
	}
	return res
}

// Returns the sum of all elements of `s`. Elements of `s` must instantiate the `Number` interface, which permits all built-in numeric types supported by Go. `Sum` overflows silently and returns a zero value when `s` is an empty slice.
func Sum[N Number](s []N) N {
	var res N
	for i := 0; i < len(s); i++ {
		res += s[i]
	}
	return res
}

// Returns the product of all elements of `s`. Elements of `s` must instantiate the `Number` interface, which permits all built-in numeric types supported by Go. `Prod` overflows silently and returns a zero value when `s` is an empty slice.
func Prod[N Number](s []N) N {
	var res N
	if len(s) < 1 {
		return res
	}
	res = s[0]
	for i := 1; i < len(s); i++ {
		res *= s[i]
	}
	return res
}

// Returns a slice of type `[]T` with a length of `n` where all elements are set to `v`.
func Repeat[T any](v T, n int) ([]T, error) {
	if n < 0 {
		return []T{}, fmt.Errorf("n cannot be less than 0")
	}
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = v
	}
	return res, nil
}

func Flatten[T any](s [][]T) []T {
	var l int
	for i := 0; i < len(s); i++ {
		l += len(s[i])
	}
	res := make([]T, 0, l)

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			res = append(res, s[i][j])
		}
	}
	return res
}
