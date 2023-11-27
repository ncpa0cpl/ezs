package convenientstructures_test

import (
	"testing"

	. "github.com/ncpa0cpl/convenient-structures"
	"github.com/stretchr/testify/assert"
)

func TestArrayIterator(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray([]int{1, 2, 3, 4, 5})

	iteratedOver := make([]int, 0)

	iterator := arr.Iterator()
	for !iterator.Done() {
		v, _ := iterator.Next()
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

	iteratedOver := make([]MapEntry[string, int], 0)

	iterator := m.Iterator()
	for !iterator.Done() {
		v, _ := iterator.Next()
		iteratedOver = append(iteratedOver, v)
	}

	assert.Equal(
		[]MapEntry[string, int]{
			{"one", 1},
			{"two", 2},
			{"three", 3},
		},
		iteratedOver,
	)
}
