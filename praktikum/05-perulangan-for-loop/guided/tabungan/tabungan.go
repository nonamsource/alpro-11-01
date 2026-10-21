package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)

	total := 0;
	tabunganHariIini := a;

	for i := 1; i <= n; i++{
		total = total + tabunganHariIini
		tabunganHariIini = tabunganHariIini + b
	}

	fmt.Println(total)	
	
}

