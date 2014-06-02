package main

import "github.com/nanoeru/go-swig/foo"
import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(foo.Foo(12, 17))
}
