package main

import "fmt"

func main() {
	var totalHari int

	fmt.Scan(&totalHari)

	tahun := totalHari / 360
	sisaHari := totalHari % 360

	bulan := sisaHari / 30
	sisaHari = sisaHari % 30

	minggu := sisaHari / 7

	hari := sisaHari % 7

	fmt.Println(tahun, bulan, minggu, hari)
}