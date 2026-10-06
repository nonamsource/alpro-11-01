package  main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)
	reamur := celsius * 4 / 5
	fmt.Printf("%.2f derajat Celsius sama dengan %.2f derajat Reamur\n", celsius, reamur)

}