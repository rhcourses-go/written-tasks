package output1

import "fmt"

func ExampleFoo1() {
	Foo1()
	// begin:Foo1
	// Output:
	// 8 7 20
	// 8-7-20
	// 14
	// 2.5

	// end:Foo1
}

func ExampleFoo2() {
	Foo2(3)
	fmt.Println()
	Foo2(4)
	fmt.Println()
	Foo2(5)

	// begin:Foo2
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

	// end:Foo2
}

func ExampleFoo3() {
	Foo3(3)
	fmt.Println()
	Foo3(4)
	fmt.Println()
	Foo3(5)

	// begin:Foo3
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

	// end:Foo3
}

func ExampleFoo4() {
	Foo4(3)
	Foo4(4)
	Foo4(5)

	// begin:Foo4
	// Output:
	// 342
	// 442
	// 542

	// end:Foo4
}

func ExampleFoo5() {
	Foo5()

	// begin:Foo5
	// Output:
	// 2342
	// 23

	// end:Foo5
}

func ExampleFoo6() {
	Foo6()

	// begin:Foo6
	// Output:
	// 3
	// 2
	// 1

	// end:Foo6
}
