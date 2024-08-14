package ezs

type Number interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64
}

type Serializable interface {
	string | bool | Number
}

type Iterable[T any] interface {
	Next() (value T, done bool)
	IterReset()
}
