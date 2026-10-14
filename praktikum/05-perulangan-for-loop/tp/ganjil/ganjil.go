package main

import "fmt"

func main () {
	var n int

	fmt.Scan(&n)

	jumlah := 0
	first := true

	for i := 1; i <= n; i += 2 {
		if !first {
			fmt.Print(" ")
		}
		fmt.Print(i)
		jumlah += i
		first = false
	
	}

	fmt.Println()
	fmt.Println(jumlah)

}
