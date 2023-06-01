package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(s, ok)

	// panic: interface conversion: interface {} is string, not float64
	f = i.(float64)
	fmt.Println(f)
}
