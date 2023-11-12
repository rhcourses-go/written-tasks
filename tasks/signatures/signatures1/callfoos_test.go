package signatures1

import "github.com/rhcourses-go/written-tasks/tests"

func Example_fooTypes() {

	tests.PrintFuncTypes(Foo1, Foo2, Foo3, Foo4, Foo5)

	// begin:solution
	// Output:
	// Foo1: func(string, int) int
	// Foo2: func(int) bool
	// Foo3: func(int, bool) []int
	// Foo4: func(bool, bool) int
	// Foo5: func(bool)

	// end:solutions
}
