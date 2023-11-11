package isprime

// IsPrimeBuggy soll true liefern, falls n eine Primzahl ist.
// Erläutern Sie den/die Fehler in der Funktion.
// begin:IsPrimeBuggy
func IsPrimeBuggy(n int) bool {
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		} else {
			return true
		}
	}
	return true
}

// end:IsPrimeBuggy

// IsPrimeCorrect liefert true, falls n eine Primzahl ist.
// begin:IsPrimeCorrect
func IsPrimeCorrect(n int) bool {
	//* Fehler 1: Spezialfall hat gefehlt, 0 und 1 sind keine Primzahlen
	if n <= 1 {
		return false
	}
	// Im
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
		//* Fehler 2: Der else-Zweig war falsch: Wenn n%i != 0,
		//* dann ist n nicht automatisch prim, denn es kann immer
		//*  noch ein anderer Teiler gefunden werden.
	}
	return true
}

// end:IsPrimeCorrect
