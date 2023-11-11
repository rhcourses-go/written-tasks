package tests

func ExamplePrintFuncTypes() {
	PrintFuncTypes(Foo1, Foo2, Foo3, Foo4, Foo5, Foo6, Foo7)

	// Output:
	// Foo1: func(string, int) int
	// Foo2: func(int) bool
	// Foo3: func(int, bool) []int
	// Foo4: func(bool, bool) int
	// Foo5: func(bool)
	// Foo6: func() int
	// Foo7: func()
}

func Foo1(string, int) int {
	return 0
}

func Foo2(int) bool {
	return false
}

func Foo3(int, bool) []int {
	return []int{}
}

func Foo4(bool, bool) int {
	return 0
}

func Foo5(bool) {}

func Foo6() int {
	return 0
}

func Foo7() {}
