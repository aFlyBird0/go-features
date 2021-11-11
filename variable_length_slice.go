package main

import "fmt"

func setSubByteSliceHeadAs2(p []byte) {
	p[1:][0] = '2'
}

func main() {
	buf := []byte{'a', 'b', 'c', 'd'}
	/**
	This shows sub-slice like `p[1:]` is also a pointer to the same underlying array as `p`,
	if elements of `p[1:]` is changed, elements of `p` is changed too.
	*/
	setSubByteSliceHeadAs2(buf)
	fmt.Println(string(buf)) // output: a2cd
}
