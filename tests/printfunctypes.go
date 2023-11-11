package tests

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// PrintFuncTypes prints the names and types of the functions passed to it.
func PrintFuncTypes(functions ...interface{}) {
	types := make([]string, len(functions))
	for i, f := range functions {
		types[i] = fmt.Sprintf("%s: %T", GetFunctionName(f), f)
	}
	fmt.Println(strings.Join(types, "\n"))
}

// GetFunctionName returns the name of the function passed to it.
// Source: https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}
