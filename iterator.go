package convenientstructures

type ArrayIterator[T any] struct {
	slice   []T
	nextIDX int
}

func NewArrayIterator[T any](array *Array[T]) *ArrayIterator[T] {
	return &ArrayIterator[T]{
		slice:   array.data,
		nextIDX: 0,
	}
}

func (a *ArrayIterator[T]) Next() (T, bool) {
	idx := a.nextIDX
	a.nextIDX = idx + 1

	return a.slice[idx], a.Done()
}

func (a *ArrayIterator[T]) Done() bool {
	return a.nextIDX >= len(a.slice)
}

type MapIterator[K comparable, V any] struct {
	entries []MapEntry[K, V]
	nextIDX int
}

func NewMapIterator[K comparable, V any](m *Map[K, V]) *MapIterator[K, V] {
	entries := make([]MapEntry[K, V], m.Count())

	i := 0
	for k, v := range m.inner {
		entries[i] = MapEntry[K, V]{
			Key:   k,
			Value: v,
		}
		i++
	}

	return &MapIterator[K, V]{
		entries: entries,
	}
}

func (m *MapIterator[K, V]) Next() (MapEntry[K, V], bool) {
	idx := m.nextIDX
	m.nextIDX = idx + 1

	return m.entries[idx], m.Done()
}

func (m *MapIterator[K, V]) Done() bool {
	return m.nextIDX >= len(m.entries)
}
