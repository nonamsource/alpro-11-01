# <h1 align="center">Laporan Praktikum Modul [05] - [Bahasa Pemrograman GO/Perulangan-for-Loop]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
[Bahasa pemrograman Go (Golang) merupakan bahasa pemrograman yang dikembangkan oleh tiga insinyur Google yaitu Robert Griesemer, Rob
Pike, dan Ken Thompson pada tahun 2007 dan dirilis secara publik pada tahun 2009. Bahasa pemrograman ini dibuat dengan tujuan menyederhanakan pengembangan perangkat lunak tanpa harus mengorbankan performa dan kecepatan kompilasi yang optimal. Bahasa pemrograman ini dirancang dengan sintaks yang minimalis, sederhana juga adanya fitur _garbage collection_ sehingga membuat bahasa ini ramah untuk pemula.]


#### 1. [Perulangan-for-loop]
[Perulangan adalah proses mengulang-ulang suatu blok kode tanpa henti selama kondisi yang dijadikan acuan terpenuhi. Dalam bahasa pemrograman Go, for merupakan satu-satunya kata kunci untuk melakukan perulangan (loop) dalam kode. Go tidak memiliki keyword while atau do-while terpisah, tetapi fungsi ketiganya dapat digabungkan langsung menggunakan for. Perulangan dibagi menjadi beberapa tipe, yaitu for standard, for sebagai while loop, for infinet loop, for each dengan range.] 






## Guided

### 1. [tabungan.go]

```go
[package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)

	total := 0;
	tabunganHariIini := a;

	for i := 1; i <= n; i++{
		total = total + tabunganHariIini
		tabunganHariIini = tabunganHariIini + b
	}

	fmt.Println(total)	
	
}	
]
```

#### Output
[Screenshoot Output Guided](/praktikum/05-perulangan-for-loop/guided/tabungan/output%20tabungan.png)


#### Deskripsi
[Dalam soal tabungan.go, program ini pertama-tama meminta pengguna untuk memasukkan tiga variabel, yaitu nilai tabungan awal (a), besaran penambahan tabungan harian (b), dan jumlah hari atau lama menabung (n). Selanjutnya, program akan menjalankan sebuah perulangan (looping) sebanyak n hari. Di setiap harinya, jumlah tabungan pada hari tersebut akan ditambahkan ke dalam variabel total, lalu nilai tabungannya akan dinaikkan sebesar b untuk hari berikutnya. Setelah perulangan selesai, program akan mencetak total keseluruhan uang yang berhasil ditabung.]

### 2. [terbesar.go]

```go
[package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var angka, max int
	fmt.Scan(&angka)
	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}

	}

	fmt.Println(max)
}
]
```

#### Output
[Screenshoot Output Guided](/praktikum/05-perulangan-for-loop/guided/terbesar/output%20terbesar.png)

#### Deskripsi
[Dalam soal terbesar.go, program meminta pengguna memasukkan variabel jumlah total data (n) terlebih dahulu, lalu membaca angka pertama yang langsung diasumsikan sebagai nilai maksimum sementara. Selanjutnya, melalui sebuah perulangan (looping) dari data kedua hingga data ke-n, program akan terus menerima input angka berikutnya dan membandingkannya dengan nilai maksimum saat ini. Jika ditemukan angka yang nilainya lebih besar dari penyimpanan maksimum sebelumnya, variabel tersebut akan diperbarui dengan angka yang baru. Setelah seluruh angka selesai dicek, program akan mencetak nilai maksimum tertinggi yang berhasil ditemukan.]

### 3. [bioskop.go]
```go
[package main

import "fmt"

func main() {
	var t int
	fmt.Print("Masukkan jumlah transaksi: ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var film, hari string
		var tiket, harga, total int

		fmt.Print("Masukkan film: ")
		fmt.Scan(&film)
		fmt.Print("Masukkan hari: ")
		fmt.Scan(&hari)
		fmt.Print("Masukkan jumlah tiket: ")
		fmt.Scan(&tiket)

		switch film {
		case "Reguler":
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				harga = 40000
			} else {
				harga = 50000
			}
		case "3D":
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				harga = 60000
			} else {
				harga = 70000
			}
		case "IMAX":
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				harga = 80000
			} else {
				harga = 90000
			}
		}

		total = harga * tiket

		if tiket >= 5 {
			total = total - total*10/100
		}

		fmt.Println("Total: ", total)
		totalPendapatan += total
	}

	fmt.Println("Total Pendapatan: ", totalPendapatan)
}
]
```

#### Output
[Screenshoot Output Guided]![alt text](/praktikum/05-perulangan-for-loop/guided/bioskop/output%20bioskop.png)

#### Deskripsi
[Dalam soal bioskop.go, program ini meminta pengguna memasukkan jumlah transaksi yang ingin diproses. Melalui perulangan sebanyak transaksi tersebut, program akan meminta detail tiap pembelian yang meliputi jenis film (Reguler, 3D, atau IMAX), hari kunjungan, dan jumlah tiket yang dibeli. Berdasarkan jenis film dan hari, program menggunakan struktur switch dan percabangan kondisi untuk menentukan harga satuan tiket di mana hari kerja (weekday) memiliki harga lebih murah dibandingkan akhir pekan (weekend). Harga satuan ini kemudian dikalikan dengan jumlah tiket untuk mendapatkan subtotal transaksi. Jika pembelian tiket mencapai 5 buah atau lebih, pelanggan otomatis mendapatkan potongan diskon sebesar 10%. Setelah menampilkan total harga untuk masing-masing transaksi, seluruh subtotal tersebut dijumlahkan, dan di akhir program dicetaklah total keseluruhan pendapatan yang diperoleh.]


<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [Prima.go]
```go
[package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlahFaktor := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			if jumlahFaktor > 0 {
				fmt.Print(" ")

			}
			fmt.Print(i)
			jumlahFaktor++
		}
	}
	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
]

```


##### Output
[Screenshoot Ouput Unguided](/praktikum/05-perulangan-for-loop/unguided/prima/output%20prima.png)



#### Deskripsi
[Dalam prima.go, program ini meminta pengguna memasukkan variabel sebuah bilangan bulat n. Melalui sebuah perulangan dari 1 hingga n, program mengecek setiap angka untuk melihat apakah habis membagi n. Jika ditemukan angka yang menjadi faktor, program akan mencetaknya ke layar dengan format dipisahkan spasi, sambil terus menghitung total faktor yang ditemukan. Di akhir program, dilakukan pengecekan apakah jumlah total faktornya tepat 2 (yakni angka 1 dan bilangan itu sendiri). Jika ya, program mencetak "Ya" yang menandakan bahwa angka tersebut merupakan bilangan prima, dan jika tidak, program akan mencetak "Tidak".]

### 2. [ecct.go]


```go
[package main

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
]
```

##### Output
[Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/unguided/ecct/output%20ecct.png)

#### Deskripsi
[Dalam soal ecct.go, program ini meminta pengguna memasukkan variabel jumlah total data peserta (m) terlebih dahulu. Melalui sebuah perulangan sebanyak m kali, program membaca enam nilai desimal berbeda untuk setiap peserta. Setiap rata-rata dari keenam komponen tersebut dihitung, lalu dibandingkan dengan batas minimum kelulusan yaitu 3.33. Jika nilai rata-ratanya mencapai atau melebihi 3.33, program akan mencetak "Ya" dan menambah jumlah penghitung peserta yang lulus, sedangkan jika di bawah itu, program mencetak "Tidak". Di akhir perulangan, program mencetak total keseluruhan peserta yang berhasil lulus.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Secara keseluruhan, kelima soal unguided program Go tersebut merupakan latihan dasar logika pemrograman yang menggabungkan perulangan dan percabangan kondisi untuk menyelesaikan berbagai masalah, mulai dari perhitungan matematika, pengelolaan transaksi keuangan, sampai evaluasi nilai akademik. Melalui struktur kontrol tersebut, program-program ini mampu memproses data secara dinamis serta melakukan akumulasi dan penyaringan hasil akhir secara otomatis dan sistematis, sehingga sangat efektif untuk mengasah ketepatan alur berpikir komputasional.]

## Referensi
1. Agung, N. (n.d.). "Dasar Pemrograman Golang". Diakses dari https://dasarpemrogramangolang.novalagung.com/A-perulangan.html
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
