package ezs_test

import (
	"testing"

	. "github.com/ncpa0cpl/ezs"
	"github.com/stretchr/testify/assert"
)

func TestArrayIterator(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray([]int{1, 2, 3, 4, 5})

	iteratedOver := make([]int, 0)

	for v := range arr.Iter() {
		iteratedOver = append(iteratedOver, v)
	}

	assert.Equal(
		[]int{1, 2, 3, 4, 5},
		iteratedOver,
	)
}

func TestMapIterator(t *testing.T) {
	assert := assert.New(t)

	m := NewMap(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	iteratedOver := make([]*MapEntry[string, int], 0)

	for entry := range m.Iter() {
		iteratedOver = append(iteratedOver, entry)
	}

	assert.Contains(
		iteratedOver,
		&MapEntry[string, int]{"one", 1},
	)
	assert.Contains(
		iteratedOver,
		&MapEntry[string, int]{"two", 2},
	)
	assert.Contains(
		iteratedOver,
		&MapEntry[string, int]{"three", 3},
	)
}

func TestIteratorBreak(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray([]string{})

	arr.Push("foo")
	arr.Push("bar")
	arr.Push("baz")
	arr.Push("qux")

	acc := ""

	for v := range arr.Iter() {
		acc = acc + v
		if v == "bar" {
			break
		}
	}

	assert.Equal(
		"foobar",
		acc,
	)
}

func TestIterateMultipleTimesOverTheSameIterator(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray([]string{})

	arr.Push("foo")
	arr.Push("bar")
	arr.Push("baz")
	iterator := arr.Iter()

	acc := ""
	iterCount := 0

	for v := range iterator {
		iterCount++
		acc = acc + v
	}

	assert.Equal(
		3,
		iterCount,
	)

	for v := range iterator {
		iterCount++
		acc = acc + v
	}

	assert.Equal(
		6,
		iterCount,
	)

	assert.Equal(
		"foobarbazfoobarbaz",
		acc,
	)
}

func TestIterateMultipleTimesOverTheSameIteratorAfterBreaking(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray([]string{})

	arr.Push("foo")
	arr.Push("bar")
	arr.Push("baz")
	iterator := arr.Iter()

	acc := ""
	iterCount := 0

	for v := range iterator {
		iterCount++
		acc = acc + v
		if v == "foo" {
			break
		}
	}

	assert.Equal(
		1,
		iterCount,
	)

	for v := range iterator {
		iterCount++
		acc = acc + v
	}

	assert.Equal(
		4,
		iterCount,
	)

	assert.Equal(
		"foofoobarbaz",
		acc,
	)
}
