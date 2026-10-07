package main

import "fmt"

func main() {
	var mil float64
	fmt.Print("Masukkan jarak dalam mil: ")
	fmt.Scanln(&mil)

	kilometer := mil * 1.6
	fmt.Printf("%.1f\n", kilometer)
}
