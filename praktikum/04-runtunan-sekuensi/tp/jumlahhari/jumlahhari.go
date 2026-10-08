package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun)
	fmt.Scan(&bulan)

	switch bulan {
		case "Januari", "Maret", "Mei", "Juli", "Agustus", "Oktober", "Desember":
			fmt.Println("Jumlah hari:", 31)
		case "April", "Juni", "September", "November":
			fmt.Println("Jumlah hari:", 30)
		case "Februari":
			if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
				fmt.Println("Jumlah hari:", 29)
			} else {
				fmt.Println("Jumlah hari:", 28)
			}
		default:
			fmt.Println("Bulan tidak valid")
	}
}