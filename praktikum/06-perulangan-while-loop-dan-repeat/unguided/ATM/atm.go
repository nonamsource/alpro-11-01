package main

import "fmt"

func main() {
	saldo := 10000000
	kesempatan := 3

	// Loop Login
	for kesempatan > 0 {
		var pwd string
		fmt.Print("Masukkan password ATM: ")
		fmt.Scanln(&pwd)

		if pwd == "1234" {
			fmt.Println("Login berhasil!")
			break
		}

		kesempatan--
		fmt.Println("Password salah. Silakan coba lagi.")

		if kesempatan == 0 {
			fmt.Println("ATM terblokir. Program selesai.")
			return
		}
	}

	// Loop Menu ATM
	for {
		fmt.Println("Menu ATM:\n1. Lihat Saldo\n2. Tarik Uang\n3. Keluar")
		var pilih int
		fmt.Print("Pilih: ")
		fmt.Scanln(&pilih)

		if pilih == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilih == 2 {
			var tarik int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: ")
			fmt.Scanln(&tarik)

			if tarik <= saldo {
				saldo -= tarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\nSaldo tersisa: Rp %d\n", tarik, saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilih == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
			return
		}
	}
}
