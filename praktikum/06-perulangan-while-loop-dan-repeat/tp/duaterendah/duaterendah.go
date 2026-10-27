package main

import "fmt"

func main() {
	var n, skor int
	var nama, nama1, nama2 string

	min1, min2 := 999, 999

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nama, &skor)
		if skor < min1 {
			min2, nama2 = min1, nama1
			min1, nama1 = skor, nama
		} else if skor < min2 {
			min2, nama2 = skor, nama
		}
	}
	fmt.Println(nama1)
	fmt.Println(nama2)
}
