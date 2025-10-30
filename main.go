package main

import (
	"bazel-go/a"
	"bazel-go/b"
	"fmt"
)

func main() {
	fmt.Println(a.Hello() + " " + b.World())
}
