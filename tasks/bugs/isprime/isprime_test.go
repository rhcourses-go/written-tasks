package isprime

import "fmt"

func ExampleIsPrimeBuggy() {
	// Die Funktion liefert für 0,1,2 und alle ungeraden Zahlen >= true.
	fmt.Println(IsPrimeBuggy(0))
	fmt.Println(IsPrimeBuggy(1))
	fmt.Println(IsPrimeBuggy(2))
	fmt.Println(IsPrimeBuggy(3))
	fmt.Println(IsPrimeBuggy(5))
	fmt.Println(IsPrimeBuggy(7))
	fmt.Println(IsPrimeBuggy(9))
	fmt.Println(IsPrimeBuggy(11))
	fmt.Println(IsPrimeBuggy(13))
	fmt.Println(IsPrimeBuggy(15))

	fmt.Println()

	// Die Funktion liefert für alle geraden Zahlen >= 4 false.
	fmt.Println(IsPrimeBuggy(4))
	fmt.Println(IsPrimeBuggy(6))
	fmt.Println(IsPrimeBuggy(8))
	fmt.Println(IsPrimeBuggy(10))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	//
	// false
	// false
	// false
	// false
}

func ExampleIsPrimeCorrect() {
	// Die Funktion liefert korrekte Ergebnisse für Primzahlen.

	fmt.Println(IsPrimeCorrect(2))
	fmt.Println(IsPrimeCorrect(3))
	fmt.Println(IsPrimeCorrect(5))
	fmt.Println(IsPrimeCorrect(7))
	fmt.Println(IsPrimeCorrect(11))
	fmt.Println(IsPrimeCorrect(13))
	fmt.Println(IsPrimeCorrect(17))

	fmt.Println()

	// Die Funktion liefert korrekte Ergebnisse für Nicht-Primzahlen.
	fmt.Println(IsPrimeCorrect(0))
	fmt.Println(IsPrimeCorrect(1))
	fmt.Println(IsPrimeCorrect(4))
	fmt.Println(IsPrimeCorrect(6))
	fmt.Println(IsPrimeCorrect(8))
	fmt.Println(IsPrimeCorrect(9))
	fmt.Println(IsPrimeCorrect(10))
	fmt.Println(IsPrimeCorrect(12))
	fmt.Println(IsPrimeCorrect(14))
	fmt.Println(IsPrimeCorrect(15))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	//
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
}
