# EZS
*Ezy Structures*

Arrays and Maps that are easier to use.

## Array Example

```go
package main

import (
		. "github.com/ncpa0cpl/ezs"
)

func main() {
	myArray := NewArray([]string{})

	myArray.Push("")
	myArray.Push("Hello")
	myArray.Push("")
	myArray.Push("World")

	myArray = myArray.Filter(func(value string) bool {
		return value != ""
	})

	fmt.Println(myArray.At(0)) // "hello"
	fmt.Println(myArray.At(1)) // "world"
}
```

## Map Example

```go
package main

import (
		. "github.com/ncpa0cpl/ezs"
)

func main() {
	myMap := NewMap(map[string]string{})

	myMap.Set("foo", "1")
	myMap.Set("bar", "2")


	fmt.Println(myMap.Has("foo")) // true

	if value, ok := myMap.Get("bar"); ok {
		fmt.Println(value) // "2"
	}

	fmt.Println(myMap.Keys()) // Array{"foo", "bar"}
	fmt.Println(myMap.Values()) // Array{"1", "2"}
}
```

## Iterators Example

```go
package main

import (
		. "github.com/ncpa0cpl/ezs"
)

func main() {
	myArray := NewArray([]string{})

	myArray.Push("foo")
	myArray.Push("bar")
	myArray.Push("baz")

	for value := range myArray.Iter() {
		fmt.Println(value) // "foo", "bar", "baz"
	}
}
```

## Custom Iterators


```go
type Iterable[T any] interface {
	Next() (value T, done bool)
	IterReset()
}
```

Any struct implementing the above interface can be used as Iterators, and be iterated over with the for range loop:

```go
package main

import (
		. "github.com/ncpa0cpl/ezs"
)

type MyIterableStruct {
	value1       string
	value2       string
	value3       string
	iteratorNext int
}

func (myStruct *MyIterableStruct) IterReset() {
	myStruct.iteratorNext = 0
}

func (myStruct *MyIterableStruct) Next() (string, bool) {
	switch myStruct.iteratorNext {
	case 0:
		myStruct.iteratorNext++
		return myStruct.value1, false
	case 1:
		myStruct.iteratorNext++
		return myStruct.value2, false
	case 2:
		myStruct.iteratorNext++
		return myStruct.value3, false
	default:
		return "", true
	}
}

func main() {
	myStruct := MyIterableStruct{"value 1", "value 2", "value 3", 0}

	for value := range Iterator(myStruct) {
		fmt.Println(value) // "value 1", "value 2", "value 3"
	}
}
```
