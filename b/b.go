package b

import "bazel-go/b/c"

func World() string {
	return c.Foo() + " World"
}
