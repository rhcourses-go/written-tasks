package signatures2

import "github.com/rhcourses-go/written-tasks/tests"

func Example_fooTypes() {

	tests.PrintFuncTypes(Foo1, Foo2, Foo3, Foo4, Foo5)

	// Output:
	// Foo1: func(int) int
	// Foo2: func(int)
	// Foo3: func([]string) bool
	// Foo4: func(bool, int) int
	// Foo5: func(string) (int, int)
}
