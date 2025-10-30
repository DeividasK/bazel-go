package c

import "bazel-go/a"

func Foo() string {
	return "Bar"
}

func Baz() string {
	return a.Hello() + " Baz"
}
