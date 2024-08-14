package ezs

func Iterator[T any](iterable Iterable[T]) func(func(T) bool) {
	return func(yield func(T) bool) {
		for {
			v, done := iterable.Next()
			if done {
				iterable.IterReset()
				return
			}

			if !yield(v) {
				iterable.IterReset()
				return
			}
		}
	}
}
