package tests

func ExamplePrintVarTypes() {
	PrintVarTypes(42, 3.14, "Hello")
	PrintVarTypes(true, false, 42, float32(2), "Hello", 'x')

	// Output:
	// int float64 string
	// bool bool int float32 string int32
}
