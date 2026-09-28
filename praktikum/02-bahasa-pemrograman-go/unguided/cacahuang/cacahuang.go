package main

import "fmt"

func main() {
	var totalUang int

	//Membaca input
	fmt.Scan(&totalUang)

	//Menghitung lembar uang pecahan 10000, 5000, 1000, dan sisa uang
	lembar10000 := totalUang / 10000
	sisa10000 := totalUang % 10000
	lembar5000 := totalUang / 5000
	sisa5000 := totalUang % 5000
	lembar1000 := totalUang / 1000
	sisa1000 := totalUang % 1000

	//Menampilkan output
	fmt.Println(lembar10000)
	fmt.Println(sisa10000)
	fmt.Println(lembar5000)
	fmt.Println(sisa5000)
	fmt.Println(lembar1000)
	fmt.Println(sisa1000)

}
