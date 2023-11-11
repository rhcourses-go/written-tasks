package output1

import "fmt"

func ExampleFoo1() {
	Foo1()
	// Output:
	// 8 7 20
	// 8-7-20
	// 14
	// 2.5
}

func ExampleFoo2() {
	Foo2(3)
	fmt.Println()
	Foo2(4)
	fmt.Println()
	Foo2(5)

	// Output:
	// *
	// **
	// ****
	//
	// *
	// **
	// ****
	// ********
	//
	// *
	// **
	// ****
	// ********
	// ****************
}

func ExampleFoo3() {
	Foo3(3)
	fmt.Println()
	Foo3(4)
	fmt.Println()
	Foo3(5)

	// Output:
	// *
	// **
	// ***
	//
	// *
	// **
	// ***
	// ****
	//
	// *
	// **
	// ***
	// ****
	// *****
}

func ExampleFoo4() {
	Foo4(3)
	Foo4(4)
	Foo4(5)

	// Output:
	// 342
	// 442
	// 542
}

func ExampleFoo5() {
	Foo5()

	// Output:
	// 2342
	// 23
}

func ExampleFoo6() {
	Foo6()

	// Output:
	// 3
	// 2
	// 1
}
