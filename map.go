package ezs

type Map[K comparable, V any] struct {
	inner   map[K]V
	keys    []K
	iterIdx int
}

func NewMap[K comparable, V any](inner map[K]V) *Map[K, V] {
	keys := make([]K, len(inner))
	i := 0
	for k := range inner {
		keys[i] = k
		i++
	}

	return &Map[K, V]{
		inner:   inner,
		keys:    keys,
		iterIdx: 0,
	}
}

func (m *Map[K, V]) addkey(key K) {
	_, ok := m.inner[key]
	if !ok {
		m.keys = append(m.keys, key)
	}
}

func (m *Map[K, V]) removekey(key K) {
	idx := -1
	for i, k := range m.keys {
		if k == key {
			idx = i
			break
		}
	}
	if idx != -1 {
		m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
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
	m.addkey(key)
	return m
}

func (m *Map[K, V]) Delete(key K) *Map[K, V] {
	delete(m.inner, key)
	m.removekey(key)
	return m
}

func (m *Map[K, V]) Count() int {
	return len(m.inner)
}

func (m *Map[K, V]) Keys() *Array[K] {
	keyListCopy := make([]K, len(m.keys))
	copy(keyListCopy, m.keys)
	return NewArray[K](keyListCopy)
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

func (m *Map[K, V]) Entries() *Array[*MapEntry[K, V]] {
	var entries = make([]*MapEntry[K, V], len(m.inner))
	i := 0
	for k, v := range m.inner {
		entries[i] = &MapEntry[K, V]{
			Key:   k,
			Value: v,
		}
		i++
	}
	return NewArray[*MapEntry[K, V]](entries)
}

func (m *Map[K, V]) ForEach(fn func(key K, value V)) {
	for k, v := range m.inner {
		fn(k, v)
	}
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

func (m *Map[K, V]) Next() (*MapEntry[K, V], bool) {
	len := len(m.inner)
	if m.iterIdx < len {
		retKey := m.keys[m.iterIdx]
		retVal := &MapEntry[K, V]{
			Key:   retKey,
			Value: m.inner[retKey],
		}
		m.iterIdx++
		return retVal, false
	}
	return nil, true
}

func (m *Map[K, V]) IterReset() {
	m.iterIdx = 0
}

func (m *Map[K, V]) Iter() func(func(*MapEntry[K, V]) bool) {
	return Iterator(m)
}
