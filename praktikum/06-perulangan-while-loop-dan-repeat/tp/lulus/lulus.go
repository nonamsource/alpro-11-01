package main

import "fmt"

func main() {
	var nilai int
	for {

		fmt.Println("Masukkan nilai anda: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			break
		}
		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}
	}
}
