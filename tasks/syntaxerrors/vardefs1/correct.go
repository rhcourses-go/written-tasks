package vardefs1

import "fmt"

func VarDefs() {
	// begin:task
	x := 42
	var y int = 55 //* var int y 55
	var z int = 77 //* int z = 42
	s := string([]byte{'a', 'b', 'c'})
	b := []byte{'a', 'b', 'c'}
	l := make([]int, 0) //* var l1 []int := make([]int, 0)
	string := "hallo"   //* string := hallo
	// end:task

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n", x, y, z, s, b, l, string)
}
