# <h1 align="center">Laporan Praktikum Modul [02] - [Bahasa Pemrograman GO/Koding, Kompilasi, dan Eksekusi Golang]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
[Bahasa pemrograman Go (Golang) merupakan bahasa pemrograman yang dikembangkan oleh tiga insinyur Google yaitu Robert Griesemer, Rob
Pike, dan Ken Thompson pada tahun 2007 dan dirilis secara publik pada tahun 2009. Bahasa pemrograman ini dibuat dengan tujuan menyederhanakan pengembangan perangkat lunak tanpa harus mengorbankan performa dan kecepatan kompilasi yang optimal. Bahasa pemrograman ini dirancang dengan sintaks yang minimalis, sederhana juga adanya fitur _garbage collection_ sehingga membuat bahasa ini ramah untuk pemula.]

### B. [Proses Koding, Kompilasi, dan Eksekusi Golang]
[Bahasa pemrograman Go bersifat _compiled_, yang di mana artinya seluruh kode perlu diterjemahkan ke dalam bahasa mesin terdahulu sebelum komputer menjalankan perintah. Hal ini juga menjadi salah satu pembeda bahasa pemrograman Go dengan beberapa bahasa pemrograman yang populer lainnya seperti Python dan Javascript. Struktur dasar program Go biasanya diawali dengan deklarasi package main, dan fungsi utama main () untuk memulai eksekusi program. Bahasa pemrograman Go juga menyediakan beberapa tipe data seperti tipe data numerik, string, dan boolean.]

#### 1. [Proses Koding Bahasa Pemrograman Go]
[Proses koding dalam bahasa pemrograman Go memiliki beberapa karakteristik dan aturan tersendiri. Mulai dari ekstensi file, dalam bahasa pemrograman Go ekstensi file wajib diakhiri dengan .go. Untuk memulai koding, bahasa pemrograman go harus berada dalam sebuah package, menggunakan deklarasi package main di baris awal kode. Setelah itu bahasa pemrograman Go memerlukan deklarasi func main () sebagai titik awal sebuah kode bisa disimpan dan dijalankan dalam komputer. Bahasa pemrograman Go juga memiliki karakteristik unik dan ketat yaitu mengenai penggunaan tool fmtgo, tool tersebut digunakan supaya baris kode yang sudah tersimpan akan otomatis terformat dengan rapih sehingga kode-kode yang tersimpan terlihat seragam.]

#### 2. [Kompilasi dan Eksekusi Bahasa Pemrograman Go]
[Seperti yang dijelaskan dengan singkat sebelumnya, bahasa pemrograman Go perlu diterjemahkan ke dalam bahasa mesin terlebih dahulu. Di dalam bahasa pemrograman Go, terdapat dua cara utama untuk proses kompilasi dan eksekusi. Cara pertama bisa dilakukan dengan mengetikkan perintah go run nama_file.go di terminal. Hal tersebut merupakan cara kompilasi sementara yang di mana perintah tersebut akan menerjemahkan kode ke biner (bahasa komputer) lalu menyimpannya sementara di memori dan dapat menjalankan kodenya di terminal.

Cara kedua bisa dilakukan dengan mengetikkan perintah go build nama_file.go di terminal. Perintah tersebut akan menerjemahkan kode ke biner secara menyeluruh dan membentuk sebuah file aplikasi, seperti file .exe di windows. Dengan cara ini, kode bisa dijalankan tanpa harus menginstall Go jika ingin menjalankan kode di _device_ lainnya]



## Guided

### 1. [skor.go]

```go
[package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaInggris int

	//Membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	//Menghitung total & rata-rata (Pembagian bilangan bulat)
	total := skorMatematika + skorBahasaInggris
	rataRata := total / 2

	//menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)

}
]
```

#### Output
[Screenshoot Output Guided]![alt text](<img width="1642" height="606" alt="output skor" src="https://github.com/user-attachments/assets/5fbc6f0f-ad8e-42f7-949a-1adfb7388c8a" />)


#### Deskripsi
[Dalam tugas skor.go, soal memerintahkan untuk membuat baris program tentang rata-rata nilai siswa di mata pelajaran bahasa Inggris dan juga matematika. Dalam penginputannya, barisan program ini menggunakan tipe data string untuk nama siswa, lalu int untuk skor mata pelajaran bahasa Inggris dan matematika, fmt.Scan untuk membaca masukan secara berurutan, setelah itu melakukan kalkulasi aritmatika (penjumlahan dan pembagian skor mata pelajaran), serta yang terakhir fmt.Println untuk menampilkan output dari kode yang sudah dibuat. Hasil akhir dari baris program ini adalah program dapat membaca input kode dengan benar dan dapat melakukan perhitungan skor dengan akurat.]

### 2. [tukar.go]

```go
[package main

import "fmt"

func main() {
	var a, b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar a dan b
	a, b = b, a

	//Menampilkan output
	fmt.Println(a)
	fmt.Println(b)

}
]
```

#### Output
[Screenshoot Output Guided]![alt text](<img width="1626" height="621" alt="output tukar" src="https://github.com/user-attachments/assets/3aed94dc-39a0-4444-b76d-ffbfc41821b5" />)


#### Deskripsi
[Dalam tugas tukar.go, soal memerintahkan untuk menukar dua bilangan bulat yang disimbolkan dengan variabel yaitu dari variabel a menjadi variabel b, dan variabel b menjadi variabel a. Dalam penginputannya, baris program ini menggunakan tipe data int untuk dua bilangan bulat acak, lalu fmt.Scan untuk membaca masukan variabel a dan variabel b secara berurutan, setelah itu adanya pross pertukaran nilai, dan yang terakhir fmt.Println untuk menampilkan output dari kode yang sudah dibuat. Hasil akhir dari program barisan ini adalah program dapat membaca input kode dengan baik dan dapat menukar dua bilangan variabel a dan b dengan akurat.]

### 3. [lingkaran.go]
```go
[package main

import "fmt"

func main() {
	var r float64

	//Membaca input
	fmt.Scan(&r)

	//Menghitung luas lingkaran
	luas := 3.14 * r * r

	//Menampilkan output
	fmt.Printf("%.2f\n", luas)
}
]
```

#### Output
[Screenshoot Output Guided]![alt text](<img width="1643" height="722" alt="output lingkaran" src="https://github.com/user-attachments/assets/1a55b134-f1fc-4f2b-aced-94f7318b5865" />)


#### Deskripsi
[Dalam tugas lingkaran.go, soal memerintahkan untuk menghitung luas suatu lingkaran. Dalam penginputannya, baris program ini menggunakan tipe data float64 untuk variabel nilai jari-jari dan luas lingkaran, lalu fmt.Scan untuk membaca masukkan nilai jari-jari, setelah itu kalkulasi aritmatika menggunakan rumus luas lingkaran (pi * r * r), dan yang terakhir fmt.Println untuk menampilkan output luas lingkaran yang menggunakan format %.2f untuk menampilkan hasil perhitungan dengan dua angka di belakang koma supaya terlihat lebih rapih. Hasil akhir dari baris program ini adalah program dapat membaca input kode dengan baik dan dapat menghitung nilai luas lingkaran dengan baik dan akurat

### 4. [suhu.go]
```go
[import "fmt"

func main() {
	var celsius float64

	//Membaca input
	fmt.Scan(&celsius)

	//Menghitung suhu dalam Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	//Menghitung suhu dalam Reamur
	Reamur := celsius * 4 / 5

	//Menghitung suhu dalam Kelvin
	kelvin := celsius + 273.15

	//Menampilkan output
	fmt.Printf("%.2f\n", fahrenheit)
	fmt.Printf("%.2f\n", reamur)
	fmt.Printf("%.2f\n", kelvin)

}
]
```

#### Output
[Screenshoot Output Guided]![alt text](<img width="1628" height="592" alt="output suhu" src="https://github.com/user-attachments/assets/1fc92e41-e792-4dca-be89-6208634843cb" />)


#### Deskripsi
[Dalam suhu.go, soal memerintahkan untuk menghitung dan mengubah nilai suhu Celsius menjadi suhu Fahrenheit, Reamur, dan Kelvin. Dalam penginputannya, baris program ini menggunakan tipe data float64 untuk variabel Celsius dan hasil konversi (perubahan suhu) ke Fahrenheit, Reamur, dan Kelvin agar nilai desimalnya tetap terjaga, lalu menggunakan fmt.Scan untuk membaca masukan nilai Celsius, setelah itu dilanjut dengan kalkulasi rumus konversi suhu, dan yang terakhir fmt.Printf juga format %.2f untuk menampilkan output nilai suhu yang rapih dengan pembulatan dua angka di belakang koma. Hasil akhir dari baris program ini adalah program dapat membaca input kode dan dapat menukar nilai suhu Celsius menjadi suhu Fahrenheit, Reamur, dan Kelvin dengan baik dan akurat.]

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [kalkulator.go]
```go
[package main

import "fmt"

func main() {
	var a, b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menampilkan output
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

}
]
```


##### Output
[Screenshoot Ouput Unguided]![alt text](<img width="1632" height="612" alt="output kalkulator go" src="https://github.com/user-attachments/assets/c5ac6c08-b1bc-4240-bc19-460712c53ab4" />)



#### Deskripsi
[Dalam tugas kalkulator.go, soal memerintahkan untuk melakukan operasi hitung dua bilangan bulat. Dalam penginputannya, baris program ini menggunakan tipe data int untuk nilai variabel bilangan bulat, lalu menggunakan fmt.Scan untuk membaca masukan dari nilai variabel bilangan bulat, setelah itu dilanjut dengan kalkulasi aritmatika (penjumlahan, pengurangan, perkalian, pembagian dan hasil kali) sekaligus menggunakan fmt.Println untuk menampilkan output operasi hitung dua bilangan bulat. Hasil akhir dari baris program ini adalah program ini dapat membaca input kode dan dapat menghitung operasi bilangan bulat dengan baik dan akurat.]

### 2. [cacahuang.go]

```go
[package main

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
]
```

##### Output
[Screenshot Output Unguided]![alt text](<img width="1632" height="600" alt="output cacahuang go" src="https://github.com/user-attachments/assets/06e53aba-d8d3-4dbe-b403-9e04655f87ba" />)


#### Deskripsi
[Dalam tugas cacahuang.go, soal memerintahkan untuk menyatakan sebuah nilai mata uang rupiah ke dalam pecahan uang sepuluh ribu, lima ribu, seribu dengan jumlah lembar se-sedikit mungkin. Dalam penginputannya, barisan program ini menggunakan tipe data int untuk variabel total uang, setelah itu dilanjut dengan kalkulasi aritmatika untuk pembagian total uang dan perhitungan sisa uang, yang terakhir program ini menggunakan fmt.Pintln untuk menampilkan output hasil pembagian dan sisa uang. Hasil akhir dari baris program ini adalahprogram ini dapat membaca input kode dan menghitung hasil pmbagian juga sisa total uang dengan baik dan akurat.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Praktikum bahasa pemrograman Go mengenai proses koding, kompilasi, dan eksekusi ini memberikan pemahaman mengenai dasar-dasar pengembangan bahasa pemrograman Go, perbedaannya dengan beberapa bahasa pemrograman lain, serta cara proses koding, kompilasi, dan eksekusi yang benar. Proses koding, kompilasi, dan eksekusi dalam praktikum ini juga didukung dengan adanya penugasan mandiri dalam tugas guided dan unguided, dan dapat disimpulkan bahwa bahasa pemrograman Go memiliki karakteristik sintaks yang minimalis tetapi tetap akurat dan optimal, selain itu wajib memuat deklarasi package main dan func main() sebagai awal ekskusi suatu program, juga penggunaan ekstensi .go di setiap file program merupakan aturan penting untuk menjalankan sebuah program di bahasa pemrograman Go.

Penggunaan tipe data yang tepat juga menjadi salah satu aturan yang penting untuk dapat menjalankan suatu program. Sepeerti tipe data string yang cocok dipakai untuk suatu data teks, int dipakai untuk suatu bilangan bulat, float (contoh float64) dipakai untuk memuat angka desimal namun dengan akhir penghitungan yang rapih dan akurat.

Dengan adanya penugasan guided dan unguided ini dapat membantu praktikkan untuk mempraktikkan secara nyata bagaimana suatu kode program dasar dibuat dan bekerja, selain itu praktikan juga dapat menerapkan logika dasar, operasi matematika, dan teknik pertukaran nilai (manipulasi variabel). Sehingga secara keseluruhan, baris program dapat bekerja dengan akurat dari proses koding, kompilasi, dan eksekusi dasar yang dilakukan berdasarkan aturan.]

## Referensi
1. [Andaria, A. C., dkk.]. ([2025]). *[Pengenalan Bahasa Pemrograman untuk Pemula]*. [Bekasi]: [PT. Hadla Media Informasi]. Diakses pada [25 September 2026] melalui [https://media.hadlacorp.com/wp-content/uploads/2025/02/PENGENALAN-BAHASA-PEMROGRAMAN-UNTUK-PEMULA.pdf]
2. [Thawaddu Shalam, Kharansyah]. ([2024]). *[Penerapan Algoritma Cosine Similiarity dan TF-IDF pada Auto-Scoring Essay Question]*. : [Knowledge Center UMN]. Diakses pada [26 September 2026] melalui [https://kc.umn.ac.id/id/eprint/32253/3/BAB_II.pdf]
3. [Prayogo, N. A.]. ([2015]). *[Dasar Pemrograman Golang]*. :  Diakses pada [26 September 2026] melalui [https://dasarpemrogramangolang.novalagung.com/]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
