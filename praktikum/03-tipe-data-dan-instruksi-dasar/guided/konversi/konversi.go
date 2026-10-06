package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu: ")
	fmt.Scanln(&celsius)

	fmt.Println("Kelvin ", celsius + 273)
}
