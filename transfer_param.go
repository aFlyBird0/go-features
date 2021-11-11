package main

// 引用传递 slice（变长数组）, map, chan
// 值传递 数组（定长slice）、等等

import "fmt"

func setByteSliceHeadAs1(p []byte) {
	p[0] = '1'
}

func main() {
	buf := []byte{'a', 'b', 'c', 'd'}

	/**
	[]byte param:
	This sample shows `[]byte` type (variable length arrays) param is just a pointer but not all elements' value.
	Everything in Go is passed by value, slices too. But a slice value is a header,
	describing a contiguous section of a backing array,
	and a slice value only contains a pointer to the array where the elements are actually stored.
	The slice value does not include its elements (unlike arrays).
	So when you pass a slice to a function, a copy will be made from this header, including the pointer,
	which will point to the same backing array.
	Modifying the elements of the slice implies modifying the elements of the backing array,
	and so all slices which share the same backing array will "observe" the change.
	*/
	setByteSliceHeadAs1(buf)
	fmt.Println(string(buf)) // output: 1bcd
}
