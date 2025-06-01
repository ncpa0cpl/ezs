package ezs_test

import (
	"strconv"
	"testing"

	. "github.com/ncpa0cpl/ezs"
	"github.com/stretchr/testify/assert"
)

func TestArrayPush(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})
	arr.Push(6, 7, 8, 9, 10).Push(11, 12, 13)

	assert.Equal(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		arr.ToSlice(),
	)
}

func TestArrayUnshift(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})
	arr.Unshift(6, 7, 8, 9, 10).Unshift(11, 12, 13)

	assert.Equal(
		[]int{11, 12, 13, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
		arr.ToSlice(),
	)
}

func TestArrayPop(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})
	poped := arr.Pop()

	assert.Equal(5, poped)
	assert.Equal(
		[]int{1, 2, 3, 4},
		arr.ToSlice(),
	)
}

func TestArrayShift(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})
	shifted := arr.Shift()

	assert.Equal(1, shifted)
	assert.Equal(
		[]int{2, 3, 4, 5},
		arr.ToSlice(),
	)
}

func TestArrayAt(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.Equal(1, arr.At(0))
	assert.Equal(2, arr.At(1))
	assert.Equal(3, arr.At(2))
	assert.Equal(4, arr.At(3))
	assert.Equal(5, arr.At(4))
	assert.Equal(5, arr.At(-1))
	assert.Equal(4, arr.At(-2))
	assert.Equal(3, arr.At(-3))
	assert.Equal(2, arr.At(-4))
	assert.Equal(1, arr.At(-5))
}

func TestArrayAtPtr(t *testing.T) {
	assert := assert.New(t)

	type Elem struct {
		value string
	}

	arr := NewArray[Elem]([]Elem{
		Elem{
			value: "first",
		},
		Elem{
			value: "second",
		},
		Elem{
			value: "third",
		},
	})

	assert.Equal("first", arr.AtPtr(0).value)
	assert.Equal("second", arr.AtPtr(1).value)
	assert.Equal("third", arr.AtPtr(2).value)

	assert.Equal("first", arr.AtPtr(-3).value)
	assert.Equal("second", arr.AtPtr(-2).value)
	assert.Equal("third", arr.AtPtr(-1).value)

	// changing the value is refelcted in the original array
	sec := arr.AtPtr(1)
	sec.value = "ABC"

	assert.Equal("ABC", arr.At(1).value)
	assert.Equal("ABC", arr.At(-2).value)
}

func TestArraySet(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})
	arr.Set(0, 10).Set(1, 20).Set(3, 40).Set(4, 50)

	assert.Equal(
		[]int{10, 20, 3, 40, 50},
		arr.ToSlice(),
	)
}

func TestArrayInsert(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	arr.Insert(0, 10)
	assert.Equal(
		[]int{10, 1, 2, 3, 4, 5},
		arr.ToSlice(),
	)

	arr.Insert(1, 20, 30).Insert(2, 11, 22, 33)
	assert.Equal(
		[]int{10, 20, 11, 22, 33, 30, 1, 2, 3, 4, 5},
		arr.ToSlice(),
	)
}

func TestArrLength(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.Equal(5, arr.Length())

	arr.Push(6, 7, 8)

	assert.Equal(8, arr.Length())
}

func TestArrayToSlice(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5, 6, 7, 8})

	assert.Equal(
		[]int{4, 5, 6},
		arr.Slice(3, 6).ToSlice(),
	)
	assert.Equal(
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		arr.ToSlice(),
	)
}

func TestArraySplice(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5, 6, 7, 8})

	assert.Equal(
		[]int{4, 5, 6},
		arr.Splice(3, 3).ToSlice(),
	)
	assert.Equal(
		[]int{1, 2, 3, 7, 8},
		arr.ToSlice(),
	)
}

func TestArrayReplace(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5, 6, 7, 8})

	assert.Equal(
		[]int{3, 4},
		arr.Replace(2, 2, 0).ToSlice(),
	)
	assert.Equal(
		[]int{1, 2, 0, 5, 6, 7, 8},
		arr.ToSlice(),
	)
}

func TestArrayConcat(t *testing.T) {
	assert := assert.New(t)

	arr1 := NewArray[int]([]int{1, 2, 3, 4, 5})
	arr2 := NewArray[int]([]int{6, 7, 8, 9, 10})

	assert.Equal(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		arr1.Concat(arr2).ToSlice(),
	)
	assert.Equal(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		arr1.ToSlice(),
	)
	assert.Equal(
		[]int{6, 7, 8, 9, 10},
		arr2.ToSlice(),
	)

	arr2.Set(1, 100)

	assert.Equal(
		[]int{6, 100, 8, 9, 10},
		arr2.ToSlice(),
	)
	assert.Equal(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		arr1.ToSlice(),
	)
}

func TestArrayReverse(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.Equal(
		[]int{5, 4, 3, 2, 1},
		arr.Reverse().ToSlice(),
	)
	assert.Equal(
		[]int{5, 4, 3, 2, 1},
		arr.ToSlice(),
	)
}

func TestArrayFilter(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.Equal(
		[]int{2, 4},
		arr.Filter(func(v, _ int) bool {
			return v%2 == 0
		}).ToSlice(),
	)
	assert.Equal(
		[]int{1, 2, 3, 4, 5},
		arr.ToSlice(),
	)
}

func TestArrayRemove(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.Equal(
		[]int{2, 4},
		arr.Remove(func(v, _ int) bool {
			return v%2 == 1
		}).ToSlice(),
	)
	assert.Equal(
		[]int{2, 4},
		arr.ToSlice(),
	)
}

func TestArrayFind(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[string]([]string{
		"foo",
		"bar",
		"baz",
		"qux",
		"quux",
	})

	found, bstring := arr.Find(func(v string, _ int) bool {
		return v[0] == 'b'
	})

	assert.True(found)
	assert.Equal("bar", bstring)

	found, qstring := arr.Find(func(v string, _ int) bool {
		return v[0] == 'q'
	})

	assert.True(found)
	assert.Equal("qux", qstring)

	found, _ = arr.Find(func(v string, _ int) bool {
		return v[0] == 'z'
	})

	assert.False(found)
}

func TestArrayFindIndex(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[string]([]string{
		"foo",
		"bar",
		"baz",
		"qux",
		"quux",
	})

	assert.Equal(1, arr.FindIndex(func(v string, _ int) bool {
		return v[0] == 'b'
	}))
	assert.Equal(3, arr.FindIndex(func(v string, _ int) bool {
		return v[0] == 'q'
	}))
	assert.Equal(-1, arr.FindIndex(func(v string, _ int) bool {
		return v[0] == 'z'
	}))
}

func TestArraySome(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[string]([]string{
		"foo",
		"bar",
		"baz",
		"qux",
		"quux",
	})

	assert.True(arr.Some(func(v string, _ int) bool {
		return v == "baz"
	}))
	assert.True(arr.Some(func(_ string, idx int) bool {
		return idx == 4
	}))
	assert.False(arr.Some(func(v string, _ int) bool {
		return v == "coorge"
	}))
}

func TestArrayEvery(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[string]([]string{
		"foo",
		"bar",
		"baz",
		"qux",
	})

	assert.True(arr.Every(func(v string, _ int) bool {
		return len(v) == 3
	}))
	assert.False(arr.Every(func(v string, _ int) bool {
		return v[0] == 'b'
	}))
}

func TestArrayForEach(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	iteratedOver := make([]int, 0)
	arr.ForEach(func(v int, _ int) {
		iteratedOver = append(iteratedOver, v)
	})

	assert.Equal(
		[]int{1, 2, 3, 4, 5},
		iteratedOver,
	)
}

func TestArraySortWith(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	arr.SortWith(func(a, b int) int {
		return b - a
	})

	assert.Equal(
		[]int{5, 4, 3, 2, 1},
		arr.ToSlice(),
	)
}

func TestArraySortWithReverse(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	arr.SortWithReverse(func(a, b int) int {
		return a - b
	})

	assert.Equal(
		[]int{5, 4, 3, 2, 1},
		arr.ToSlice(),
	)
}

func TestArraySort(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{5, 4, 3, 2, 1})

	Sort(arr, func(a int) int {
		return a
	})

	assert.Equal(
		[]int{1, 2, 3, 4, 5},
		arr.ToSlice(),
	)
}

func TestArraySortReverse(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	SortReverse(arr, func(a int) int {
		return a
	})

	assert.Equal(
		[]int{5, 4, 3, 2, 1},
		arr.ToSlice(),
	)
}

func TestArrayMap(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.Equal(
		[]string{"1", "2", "3", "4", "5"},
		MapTo(arr, func(v int) string {
			return strconv.Itoa(v)
		}).ToSlice(),
	)
	assert.Equal(
		[]int{1, 2, 3, 4, 5},
		arr.ToSlice(),
	)
}

func TestArrayContains(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.False(Contains(arr, 0))
	assert.True(Contains(arr, 1))
	assert.True(Contains(arr, 2))
	assert.True(Contains(arr, 3))
	assert.True(Contains(arr, 4))
	assert.True(Contains(arr, 5))
	assert.False(Contains(arr, 6))
}

func TestArrayIndexOf(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5})

	assert.Equal(-1, IndexOf(arr, 0))
	assert.Equal(0, IndexOf(arr, 1))
	assert.Equal(1, IndexOf(arr, 2))
	assert.Equal(2, IndexOf(arr, 3))
	assert.Equal(3, IndexOf(arr, 4))
	assert.Equal(4, IndexOf(arr, 5))
	assert.Equal(-1, IndexOf(arr, 6))
}

func TestArrayLastIndexOf(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[int]([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})

	assert.Equal(-1, LastIndexOf(arr, 0))
	assert.Equal(8, LastIndexOf(arr, 1))
	assert.Equal(7, LastIndexOf(arr, 2))
	assert.Equal(6, LastIndexOf(arr, 3))
	assert.Equal(5, LastIndexOf(arr, 4))
	assert.Equal(4, LastIndexOf(arr, 5))
	assert.Equal(-1, LastIndexOf(arr, 6))
}

func TestArrayJoin(t *testing.T) {
	assert := assert.New(t)

	arr := NewArray[string]([]string{"foo", "bar", "baz", "qux", "quux"})
	assert.Equal(
		"foo,bar,baz,qux,quux",
		Join(arr, ","),
	)

	arr2 := NewArray[int]([]int{1, 2, 3, -5, 4, 5, -100})
	assert.Equal(
		"1,2,3,-5,4,5,-100",
		Join(arr2, ","),
	)

	arr3 := NewArray[bool]([]bool{true, true, false, true})
	assert.Equal(
		"true,true,false,true",
		Join(arr3, ","),
	)

	arr4 := NewArray[uint]([]uint{10, 20, 30})
	assert.Equal(
		"10,20,30",
		Join(arr4, ","),
	)

	arr5 := NewArray[uint8]([]uint8{100, 99})
	assert.Equal(
		"100,99",
		Join(arr5, ","),
	)

	arr6 := NewArray[uint16]([]uint16{69, 420, 69})
	assert.Equal(
		"69,420,69",
		Join(arr6, ","),
	)

	arr7 := NewArray[uint32]([]uint32{1000, 10000, 1000000})
	assert.Equal(
		"1000,10000,1000000",
		Join(arr7, ","),
	)

	arr8 := NewArray[uint64]([]uint64{22, 222, 2222, 222, 22})
	assert.Equal(
		"22,222,2222,222,22",
		Join(arr8, ","),
	)

	arr9 := NewArray[int8]([]int8{0, -1, 1, -2, 2})
	assert.Equal(
		"0,-1,1,-2,2",
		Join(arr9, ","),
	)

	arr10 := NewArray[int16]([]int16{-1, -11, -111, -1111, -11111})
	assert.Equal(
		"-1,-11,-111,-1111,-11111",
		Join(arr10, ","),
	)

	arr11 := NewArray[int32]([]int32{-99, 99, 500, -500, 1000000, -1000000})
	assert.Equal(
		"-99,99,500,-500,1000000,-1000000",
		Join(arr11, ","),
	)

	arr12 := NewArray[int64]([]int64{-999999999999999999, 999999999999999999, 0, 1234})
	assert.Equal(
		"-999999999999999999,999999999999999999,0,1234",
		Join(arr12, ","),
	)

	arr13 := NewArray[float32]([]float32{0.12345568, 12344567, 69.420})
	assert.Equal(
		"0.12345568,12344567,69.42",
		Join(arr13, ","),
	)

	arr14 := NewArray[float64]([]float64{0, 6.9, 98765.56789, -654321.12345, -4.20, -0})
	assert.Equal(
		"0,6.9,98765.56789,-654321.12345,-4.2,0",
		Join(arr14, ","),
	)
}
