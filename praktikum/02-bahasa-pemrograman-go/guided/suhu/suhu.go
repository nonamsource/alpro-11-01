package main

import "fmt"

func main() {
	var celsius float64

	//Membaca input
	fmt.Scan(&celsius)

	//Menghitung suhu dalam Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	//Menghitung suhu dalam reamur
	reamur := celsius * 4 / 5

	//Menghitung suhu dalam Kelvin
	kelvin := celsius + 273.15

	//Menampilkan output
	fmt.Printf("%.2f\n", fahrenheit)
	fmt.Printf("%.2f\n", reamur)
	fmt.Printf("%.2f\n", kelvin)

}
