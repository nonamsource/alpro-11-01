package main

import "fmt"

func main() {
	var a int //bilangan yang dikalikan
	var b int //batas perkalian

	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}
}
