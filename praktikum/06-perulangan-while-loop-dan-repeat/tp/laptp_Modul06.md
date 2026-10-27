# <h1 align="center">Tugas Pendahuluan Modul [06] - [Bahasa Pemrograman GO/Perulangan while loop dan repeat]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>



### 1. [Biaya.go]

```go
[package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	biaya := 10000000.0

	for i := 2; i <= n; i++ {
		biaya *= 1.06
	}
	fmt.Printf("%.0f\n", biaya)
}
]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/biaya/output%20biaya.png)


#### Deskripsi
[Program Go di atas berfungsi untuk menghitung perkiraan suatu biaya atau nilai awal yang mengalami kenaikan sebesar 6% setiap periodenya berdasarkan input dari pengguna, dengan memanfaatkan beberapa variabel dan tipe data penting. Variabel n menggunakan tipe data int (bilangan bulat) untuk menyimpan jumlah periode atau tahun dari input pengguna, variabel biaya menggunakan tipe data float64 (bilangan desimal) untuk menampung nilai awal sebesar 10.000.000,0 sekaligus menyimpan akumulasi hasil perhitungannya, serta variabel i dengan tipe data int yang bertindak sebagai penghitung iterasi di dalam perulangan. Cara kerja program ini dimulai dengan membaca nilai n, kemudian menjalankan perulangan dari periode ke-2 hingga n di mana nilai biaya akan dikalikan dengan 1,06 secara berulang, dan diakhiri dengan mencetak hasil akhir akumulasi biaya tersebut ke layar dalam format angka bulat tanpa nilai desimal.]


### 2. [deret.go]
```go
[package main

import "fmt"

func main() {
	var n, total int
	fmt.Scan(&n)

	i := 1
	for i <= n {
		total += i
		i++
	}
	fmt.Println(total)
}
]
```

#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/deret/output%20deret.png)

#### Deskripsi
[Program Go di atas berfungsi untuk menghitung jumlah total dari deret bilangan bulat dari 1 hingga n (penjumlahan beruntun) berdasarkan input dari pengguna, dengan memanfaatkan beberapa variabel berjenis data int (bilangan bulat). Variabel n digunakan untuk menyimpan batas angka input dari pengguna, variabel total (yang secara default bernilai awal 0) berfungsi untuk menampung akumulasi hasil penjumlahan, serta variabel i bertindak sebagai penghitung (counter) iterasi yang dimulai dari angka 1. Cara kerja program ini dimulai dengan membaca nilai n dari pengguna, kemudian menjalankan perulangan (for) selama i kurang dari atau sama dengan n, di mana pada setiap perulangan nilai i akan ditambahkan ke dalam total dan nilai i akan terus bertambah satu hingga mencapai batas n, dan diakhiri dengan mencetak hasil akhir penjumlahan tersebut ke layar menggunakan fungsi fmt.Println.]


### 3. [duaterendah.go]
```go
[package main

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
]
```
#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/duaterendah/output%20duaterendah.png)

#### Deskripsi
[Program Go di atas berfungsi untuk mencari dan menampilkan dua nama dengan skor terkecil (terendah) dari sekumpulan data peserta berdasarkan jumlah input n yang dimasukkan oleh pengguna, dengan memanfaatkan beberapa variabel dan tipe data yang sesuai. Dalam program ini, terdapat variabel n dan skor dengan tipe data int (bilangan bulat) untuk menyimpan jumlah total data serta nilai skor peserta, variabel nama, nama1, dan nama2 dengan tipe data string (teks) untuk menyimpan nama peserta, serta variabel min1 dan min2 berjenis int yang diinisialisasi dengan nilai awal besar (999) untuk melacak skor terkecil pertama dan kedua. Cara kerja program ini dimulai dengan membaca jumlah data n, lalu melakukan perulangan sebanyak n kali untuk membaca pasangan nama dan skor dari pengguna; di dalam perulangan tersebut, program akan menyeleksi dan memperbarui posisi nama1 dan nama2 secara dinamis jika ditemukan skor yang lebih kecil dari minimum yang sedang tercatat, dan di bagian akhir program akan mencetak dua nama dengan skor terendah tersebut secara berurutan ke layar menggunakan fungsi fmt.Println.]


### 4. [lulus.go]
```go
[package main

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
]
```
#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/lulus/output%20lulus.png)

#### Deskripsi
[Program Go di atas berfungsi untuk membuat sistem pengecekan kelulusan ujian secara terus-menerus (berulang) hingga pengguna memasukkan angka tertentu untuk menghentikannya, dengan memanfaatkan variabel berjenis data int (bilangan bulat). Variabel nilai digunakan untuk menyimpan input angka atau skor yang dimasukkan oleh pengguna pada setiap iterasi. Cara kerja program ini dimulai dengan menjalankan perulangan tak terbatas (for tanpa kondisi), di mana program akan mencetak teks permintaan input dan membaca nilai dari pengguna; jika pengguna memasukkan angka -1, perulangan akan langsung berhenti karena adanya perintah break, namun jika nilainya bukan -1, program akan melakukan percabangan (if-else) untuk mengecek apakah nilai lebih besar dari atau sama dengan 60 (maka akan mencetak pesan "Anda lulus ujian"), atau kurang dari 60 (maka akan mencetak pesan "Anda tidak lulus ujian"), dan proses ini akan terus berulang sampai pengguna memasukkan angka -1.]




## Kesimpulan
[Dari keempat kode program Go yang telah dibahas, dapat disimpulkan bahwa seluruh program tersebut memperlihatkan bagaimana fondasi dasar pemrograman digunakan untuk menyelesaikan berbagai jenis masalah secara terstruktur. Keempatnya mengombinasikan penggunaan variabel dan tipe data seperti int untuk bilangan bulat, float64 untuk perhitungan desimal, serta string untuk data berbentuk teks demi menyimpan informasi dengan tepat. Selain itu, konsep perulangan (for) diterapkan secara fleksibel, mulai dari melakukan perhitungan berkala dan akumulasi penjumlahan, memproses data masukan secara berulang, hingga menciptakan sistem interaktif tak terbatas yang terus berjalan. Fondasi ini diperkuat dengan logika percabangan (if-else) yang memungkinkan program untuk mengambil keputusan secara otomatis, seperti menentukan kelulusan, menyeleksi dua nilai paling minimum, atau menghentikan jalannya program saat menerima masukan tertentu. Secara keseluruhan, keempat kodingan tersebut menjadi gambaran sederhana bagaimana sebuah program menerima masukan dari pengguna, mengolahnya melalui proses logika dan perhitungan, lalu menghasilkan keluaran yang berguna.]