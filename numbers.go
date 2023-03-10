package slicefuncs

type Number interface {
	uint | uintptr | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 |
		complex64 | complex128
}

type RealNumber interface {
	uint | uintptr | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64
}
