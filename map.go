package convenientstructures

type Map[K comparable, V any] struct {
	inner map[K]V
}

func NewMap[K comparable, V any](inner map[K]V) *Map[K, V] {
	return &Map[K, V]{
		inner: inner,
	}
}

func (m *Map[K, V]) Has(key K) bool {
	_, ok := m.inner[key]
	return ok
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	v, ok := m.inner[key]
	return v, ok
}

func (m *Map[K, V]) Set(key K, value V) *Map[K, V] {
	m.inner[key] = value
	return m
}

func (m *Map[K, V]) Delete(key K) *Map[K, V] {
	delete(m.inner, key)
	return m
}

func (m *Map[K, V]) Count() int {
	return len(m.inner)
}

func (m *Map[K, V]) Keys() *Array[K] {
	var keys = make([]K, 0)
	for k := range m.inner {
		keys = append(keys, k)
	}
	return NewArray[K](keys)
}

func (m *Map[K, V]) Values() *Array[V] {
	var values = make([]V, 0)
	for _, v := range m.inner {
		values = append(values, v)
	}
	return NewArray[V](values)
}

type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

func (m *Map[K, V]) Entries() *Array[MapEntry[K, V]] {
	var entries = make([]MapEntry[K, V], len(m.inner))
	i := 0
	for k, v := range m.inner {
		entries[i] = MapEntry[K, V]{
			Key:   k,
			Value: v,
		}
		i++
	}
	return NewArray[MapEntry[K, V]](entries)
}

func (m *Map[K, V]) ForEach(fn func(key K, value V)) {
	for k, v := range m.inner {
		fn(k, v)
	}
}

func (m *Map[K, V]) Iterator() *MapIterator[K, V] {
	return NewMapIterator[K, V](m)
}

func (m *Map[K, V]) ToMap() map[K]V {
	var newMap = make(map[K]V)
	for k, v := range m.inner {
		newMap[k] = v
	}
	return newMap
}

func (m *Map[K, V]) Find(fn func(key K, value V) bool) (V, bool) {
	for k, v := range m.inner {
		if fn(k, v) {
			return v, true
		}
	}

	var zeroV V
	return zeroV, false
}

func (m *Map[K, V]) FindKey(fn func(key K, value V) bool) (K, bool) {
	for k, v := range m.inner {
		if fn(k, v) {
			return k, true
		}
	}

	var zeroK K
	return zeroK, false
}

func (m *Map[K, V]) Copy() *Map[K, V] {
	newMap := make(map[K]V)
	for k, v := range m.inner {
		newMap[k] = v
	}
	return NewMap[K, V](newMap)
}
