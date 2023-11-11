package tests

import (
	"fmt"
	"strings"
)

// PrintVarTypes gibt die Typen der Variablen aus, die ihr übergeben werden.
func PrintVarTypes(values ...interface{}) {
	types := make([]string, len(values))
	for i, v := range values {
		types[i] = fmt.Sprintf("%T", v)
	}
	fmt.Println(strings.Join(types, " "))
}
