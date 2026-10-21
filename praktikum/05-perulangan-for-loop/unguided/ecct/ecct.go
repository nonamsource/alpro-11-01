package main

import "fmt"

func main() {
	var m, lulus int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		var p, v, ga, fs, ci, tc float64
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		if (p+v+ga+fs+ci+tc)/6 >= 3.33 {
			fmt.Println("Ya")
			lulus++
		} else {
			fmt.Println("Tidak")
		}
	}
	fmt.Println(lulus)
}
