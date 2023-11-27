package convenientstructures

import (
	"cmp"
	"slices"
	"strconv"
)

type Array[T any] struct {
	data []T
}

func NewArray[T any](data []T) *Array[T] {
	return &Array[T]{data}
}

// Adds new elements to the end of the array
func (a *Array[T]) Push(data ...T) *Array[T] {
	a.data = append(a.data, data...)
	return a
}

// Adds new elements to the beginning of the array
func (a *Array[T]) Unshift(data ...T) *Array[T] {
	a.data = append(data, a.data...)
	return a
}

// Removes the last element from an array and returns that element
func (a *Array[T]) Pop() T {
	lastIdx := len(a.data) - 1
	data := a.data[lastIdx]
	a.data = a.data[:lastIdx]
	return data
}

// Removes the first element from an array and returns that element
func (a *Array[T]) Shift() T {
	data := a.data[0]
	a.data = a.data[1:]
	return data
}

// Returns the element at the specified index
func (a *Array[T]) At(idx int) T {
	if idx < 0 {
		idx = len(a.data) + idx
	}

	return a.data[idx]
}

// Changes the value at the specified index
func (a *Array[T]) Set(idx int, data T) *Array[T] {
	a.data[idx] = data
	return a
}

// Inserts new elements at the specified index, shifting the
// elements after the index
func (a *Array[T]) Insert(after int, data ...T) *Array[T] {
	a.data = append(a.data[:after+1], append(data, a.data[after+1:]...)...)
	return a
}

// Returns the length of the array
func (a *Array[T]) Length() int {
	return len(a.data)
}

// Returns a shallow copy of a portion of an array
func (a *Array[T]) Slice(start, end int) *Array[T] {
	return NewArray[T](a.data[start:end])
}

// Removes elements from an array from the given index range
// and returns the removed elements
func (a *Array[T]) Splice(start, end int) *Array[T] {
	s := slices.Clone(a.data[start:end])
	a.data = append(a.data[:start], a.data[end:]...)
	return NewArray[T](s)
}

// Removes elements from an array from the given index range,
// inserts new elements in their place and returns the removed elements
func (a *Array[T]) Replace(start, end int, data ...T) *Array[T] {
	s := slices.Clone(a.data[start:end])
	a.data = append(a.data[:start], append(data, a.data[end:]...)...)
	return NewArray[T](s)
}

// Concatenates in place the elements of a provided array
func (a *Array[T]) Concat(arr *Array[T]) *Array[T] {
	a.data = append(a.data, arr.data...)
	return a
}

// Reverses the array in place
func (a *Array[T]) Reverse() *Array[T] {
	slices.Reverse[[]T](a.data)
	return a
}

// Removes elements from the array that do not satisfy the
// predicate
func (a *Array[T]) Filter(predicate func(T, int) bool) *Array[T] {
	var arr []T
	for idx, v := range a.data {
		if predicate(v, idx) {
			arr = append(arr, v)
		}
	}
	a.data = arr
	return a
}

// Returns the first element in the array that satisfies the
// predicate
func (a *Array[T]) Find(predicate func(T, int) bool) (bool, T) {
	for idx, v := range a.data {
		if predicate(v, idx) {
			return true, v
		}
	}
	var zero T
	return false, zero
}

// Returns the index of the first element in the array that
// satisfies the predicate
func (a *Array[T]) FindIndex(predicate func(T, int) bool) int {
	for idx, v := range a.data {
		if predicate(v, idx) {
			return idx
		}
	}
	return -1
}

// Returns true if at least one element in the array satisfies
// the predicate
func (a *Array[T]) Some(predicate func(T, int) bool) bool {
	for idx, v := range a.data {
		if predicate(v, idx) {
			return true
		}
	}
	return false
}

// Returns true if all elements in the array satisfy the predicate
func (a *Array[T]) Every(predicate func(T, int) bool) bool {
	for idx, v := range a.data {
		if !predicate(v, idx) {
			return false
		}
	}
	return true
}

func (a *Array[T]) ForEach(callback func(T, int)) {
	for idx, v := range a.data {
		callback(v, idx)
	}
}

func (a *Array[T]) Iterator() Iterator[T] {
	return NewArrayIterator[T](a)
}

// Creates a shallow copy of the array
func (a *Array[T]) Copy() *Array[T] {
	arr := make([]T, len(a.data))
	copy(arr, a.data)
	return NewArray[T](arr)
}

// Creates a new slice with the same elements as the array and returns it
func (a *Array[T]) ToSlice() []T {
	return a.Copy().data
}

// Sorts the array in place
func (a *Array[T]) SortWith(compare func(T, T) int) {
	slices.SortFunc[[]T, T](a.data, compare)
}

// Sorts the array in place in reverse order
func (a *Array[T]) SortWithReverse(compare func(T, T) int) {
	slices.SortFunc[[]T, T](a.data, func(a, b T) int {
		return compare(b, a)
	})
}

// Sorts the array in place using the provided function to get the
// comparable value
func Sort[T any, C cmp.Ordered](array *Array[T], getComparable func(T) C) {
	slices.SortFunc[[]T, T](array.data, func(a, b T) int {
		return cmp.Compare(getComparable(a), getComparable(b))
	})
}

// Sorts the array in place in reverse order using the provided
// function to get the comparable value
func SortReverse[T any, C cmp.Ordered](array *Array[T], getComparable func(T) C) {
	slices.SortFunc[[]T, T](array.data, func(a, b T) int {
		return cmp.Compare(getComparable(b), getComparable(a))
	})
}

// Calls a defined callback function on each element of an array,
// and returns an array that contains the results.
func MapTo[T any, U any](array *Array[T], mapper func(T) U) *Array[U] {
	arr := make([]U, len(array.data))
	for idx, v := range array.data {
		arr[idx] = mapper(v)
	}
	return NewArray[U](arr)
}

// Returns true if any element in the array is equal to the given
// element, -1 if none is equal
func Contains[T comparable](array *Array[T], elem T) bool {
	for _, v := range array.data {
		if v == elem {
			return true
		}
	}
	return false
}

// Returns the index of the first element in the array that is
// equal to the given element, -1 if none is equal
func IndexOf[T comparable](array *Array[T], elem T) int {
	for idx, v := range array.data {
		if v == elem {
			return idx
		}
	}
	return -1
}

// Returns the index of the last element in the array that is equal
// to the given element, -1 if none is equal
func LastIndexOf[T comparable](array *Array[T], elem T) int {
	for idx := len(array.data) - 1; idx >= 0; idx-- {
		if array.data[idx] == elem {
			return idx
		}
	}
	return -1
}

// Removes consecutive duplicates from the array in place
func Compact[T comparable](array *Array[T]) *Array[T] {
	array.data = slices.Compact[[]T, T](array.data)
	return array
}

// Compares all elements of the arrays. The result is 0 if a == b,
// -1 if a < b, and +1 if a > b.
func Compare[T cmp.Ordered](a, b *Array[T]) int {
	return slices.Compare[[]T, T](a.data, b.data)
}

// Joins all elements of the array into a string separated by the
// given separator
func Join[T Serializable](array *Array[T], sep string) string {
	str := ""
	for idx, v := range array.data {
		if idx > 0 {
			str += sep
		}
		switch x := any(v).(type) {
		case string:
			str += x
		case bool:
			str += strconv.FormatBool(x)
		case uint:
			str += strconv.FormatUint(uint64(x), 10)
		case uint8:
			str += strconv.FormatUint(uint64(x), 10)
		case uint16:
			str += strconv.FormatUint(uint64(x), 10)
		case uint32:
			str += strconv.FormatUint(uint64(x), 10)
		case uint64:
			str += strconv.FormatUint(x, 10)
		case int:
			str += strconv.FormatInt(int64(x), 10)
		case int8:
			str += strconv.FormatInt(int64(x), 10)
		case int16:
			str += strconv.FormatInt(int64(x), 10)
		case int32:
			str += strconv.FormatInt(int64(x), 10)
		case int64:
			str += strconv.FormatInt(x, 10)
		case float32:
			str += strconv.FormatFloat(float64(x), 'f', -1, 32)
		case float64:
			str += strconv.FormatFloat(x, 'f', -1, 64)
		}
	}
	return str
}
