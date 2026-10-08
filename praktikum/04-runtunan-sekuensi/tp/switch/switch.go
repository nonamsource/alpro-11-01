package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Bukan 1, 2, atau 3")
	}
}
