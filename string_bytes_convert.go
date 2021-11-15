package main

import "fmt"

func main()  {
	// when convert string to byte array, string terminator \0 is not included
	s1 := "str1"
	b1 := []byte(s1)
	fmt.Println(len(b1)) // output 4

	// when convert byte array to string, string terminator \0 will be automatically added.
	b2 := []byte{'s', 't', 'r', '2'}
	s2 := string(b2)
	fmt.Println(len(s2)) // output 4
	s2 = string(b2[:3])
	fmt.Println(len(s2)) // output 3
}