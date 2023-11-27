package convenientstructures_test

import (
	"cmp"
	"slices"
	"testing"

	. "github.com/ncpa0cpl/convenient-structures"
	"github.com/stretchr/testify/assert"
)

func TestMapSetAndGet(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	m.Set("four", 4)
	m.Set("five", 5)
	m.Set("six", 6)

	assert.Equal(
		6,
		m.Count(),
	)

	v1, ok1 := m.Get("one")
	v2, ok2 := m.Get("two")
	v3, ok3 := m.Get("three")
	v4, ok4 := m.Get("four")
	v5, ok5 := m.Get("five")
	v6, ok6 := m.Get("six")
	_, ok7 := m.Get("seven")

	assert.True(ok1)
	assert.Equal(
		1,
		v1,
	)
	assert.True(ok2)
	assert.Equal(
		2,
		v2,
	)
	assert.True(ok3)
	assert.Equal(
		3,
		v3,
	)
	assert.True(ok4)
	assert.Equal(
		4,
		v4,
	)
	assert.True(ok5)
	assert.Equal(
		5,
		v5,
	)
	assert.True(ok6)
	assert.Equal(
		6,
		v6,
	)
	assert.False(ok7)
}

func TestMapDelete(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	m.Delete("two")

	assert.Equal(
		2,
		m.Count(),
	)

	_, ok1 := m.Get("one")
	_, ok2 := m.Get("two")
	_, ok3 := m.Get("three")

	assert.True(ok1)
	assert.False(ok2)
	assert.True(ok3)
}

func TestMapHas(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	assert.True(m.Has("one"))
	assert.True(m.Has("two"))
	assert.True(m.Has("three"))
	assert.False(m.Has("four"))
}

func TestMapCount(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	assert.Equal(
		3,
		m.Count(),
	)

	m.Set("four", 4)

	assert.Equal(
		4,
		m.Count(),
	)
}

func TestMapKeys(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	keys := m.Keys()

	assert.Equal(
		3,
		keys.Length(),
	)

	assert.True(Contains(keys, "one"))
	assert.True(Contains(keys, "two"))
	assert.True(Contains(keys, "three"))
	assert.False(Contains(keys, "four"))
}

func TestMapValues(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	values := m.Values()

	assert.Equal(
		3,
		values.Length(),
	)

	assert.True(Contains(values, 1))
	assert.True(Contains(values, 2))
	assert.True(Contains(values, 3))
	assert.False(Contains(values, 4))
}

func TestMapEntries(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	entries := m.Entries()

	assert.Equal(
		3,
		entries.Length(),
	)

	assert.True(entries.Some(func(entry MapEntry[string, int], _ int) bool {
		return entry.Key == "one" && entry.Value == 1
	}))
	assert.True(entries.Some(func(entry MapEntry[string, int], _ int) bool {
		return entry.Key == "two" && entry.Value == 2
	}))
	assert.True(entries.Some(func(entry MapEntry[string, int], _ int) bool {
		return entry.Key == "three" && entry.Value == 3
	}))
	assert.False(entries.Some(func(entry MapEntry[string, int], _ int) bool {
		return entry.Key == "four"
	}))
}

func TestMapForEach(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	entries := []MapEntry[string, int]{}

	m.ForEach(func(key string, value int) {
		entries = append(entries, MapEntry[string, int]{
			Key:   key,
			Value: value,
		})
	})

	slices.SortFunc(entries, func(a, b MapEntry[string, int]) int {
		return cmp.Compare(a.Key, b.Key)
	})

	assert.Equal(
		3,
		len(entries),
	)

	assert.Equal(
		"one",
		entries[0].Key,
	)
	assert.Equal(
		1,
		entries[0].Value,
	)

	assert.Equal(
		"three",
		entries[1].Key,
	)
	assert.Equal(
		3,
		entries[1].Value,
	)

	assert.Equal(
		"two",
		entries[2].Key,
	)
	assert.Equal(
		2,
		entries[2].Value,
	)
}

func TestMapToMap(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	m2 := m.ToMap()

	assert.Equal(
		map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
		},
		m2,
	)

	m2["two"] = 5

	assert.Equal(
		5,
		m2["two"],
	)

	two, ok := m.Get("two")
	assert.True(ok)
	assert.Equal(
		2,
		two,
	)
}

func TestMapFind(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	value, ok := m.Find(func(key string, value int) bool {
		return key == "two"
	})

	assert.True(ok)
	assert.Equal(
		2,
		value,
	)

	_, ok = m.Find(func(key string, value int) bool {
		return key == "four"
	})

	assert.False(ok)
}

func TestMapFindKey(t *testing.T) {
	assert := assert.New(t)

	m := NewMap[string, int](map[string]int{
		"one":   10,
		"two":   20,
		"three": 30,
	})

	key, ok := m.FindKey(func(key string, value int) bool {
		return value == 20
	})

	assert.True(ok)
	assert.Equal(
		"two",
		key,
	)

	_, ok = m.FindKey(func(key string, value int) bool {
		return value == 0
	})

	assert.False(ok)
}

func TestMapCopy(t *testing.T) {
	assert := assert.New(t)

	m1 := NewMap[string, int](map[string]int{
		"one":   10,
		"two":   20,
		"three": 30,
	})

	m2 := m1.Copy()

	assert.Equal(
		3,
		m2.Count(),
	)

	m2.Set("four", 40)

	assert.Equal(
		4,
		m2.Count(),
	)

	assert.Equal(
		3,
		m1.Count(),
	)
}
