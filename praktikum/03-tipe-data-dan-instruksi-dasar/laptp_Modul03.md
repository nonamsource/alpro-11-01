# <h1 align="center">Tugas Pendahuluan Modul [Nomor Modul] - [Judul Modul/Topik]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>

### 1. Sisa Kue

```go
[package main

import "fmt"

func main() {
	var y, x int
	fmt.Scan(&y, &x)

	sisa := y % x
	fmt.Println("Sisa ", sisa)

}
]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisakue.png)


#### Deskripsi
[Di praktikum kali ini, saya belajar tentang tipe data dasar dan cara membuat instruksi sederhana pakai bahasa Go. Di sini, saya mengerjakan bagian guided lalu lanjut bikin program unguided berjudul "Sisa Kue". Program ini dibuat untuk menghitung sisa kue dengan cara memasukkan dua angka (jumlah kue dan jumlah orang atau bagiannya), lalu dibagi menggunakan operasi sisa bagi (modulus). Hasilnya, program berhasil berjalan lancar dan menampilkan sisa kue yang benar sesuai dengan angka yang saya input di terminal, seperti yang kelihatan di gambar output-nya.]

### 2. [bool.go]

```go
[package main

import "fmt"

func main () {
	var nilai bool
	fmt.Scan(&nilai)
	fmt.Println(nilai)
}
]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/boolean.png)


#### Deskripsi
[Selain itu, saya juga mengerjakan program bool.go untuk memahami penggunaan tipe data boolean (true atau false) di bahasa Go. Program ini dibuat sederhana untuk menerima masukan nilai kebenaran dari terminal, lalu langsung mencetaknya kembali ke layar. Dari hasil pengujian, program berhasil berjalan dengan lancar dan menampilkan nilai boolean yang sesuai dengan input yang diberikan.]


### 3. [kilometer.go]
```go
[package main

import "fmt"

func main() {
	var mil float64
	fmt.Print("Masukkan jarak dalam mil: ")
	fmt.Scanln(&mil)

	kilometer := mil * 1.6
	fmt.Printf("%.1f\n", kilometer)
}
]
```

#### Output
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/kilometer.png)

#### Deskripsi
[Selanjutnya, saya juga membuat program untuk mengonversi jarak dari satuan mil ke kilometer menggunakan tipe data pecahan (float64). Di sini, program meminta pengguna memasukkan nilai jarak dalam bentuk mil, lalu menghitung dan mengubahnya ke dalam bentuk kilometer, serta mencetak hasilnya dengan format satu angka di belakang koma. Program ini juga berhasil berjalan dengan lancar saat diuji.]

## Kesimpulan
[Secara keseluruhan, rangkaian tugas pendahuluan ini membantu saya memahami dasar-dasar pemrograman menggunakan bahasa Go, mulai dari pengenalan struktur kode dasar, penggunaan berbagai tipe data seperti integer, boolean, hingga bilangan pecahan (float64), serta cara mengoperasikan instruksi input dan output. Melalui latihan membuat program seperti perhitungan sisa pembagian (modulus), pemrosesan nilai kebenaran, hingga konversi satuan jarak, saya bisa melatih logika pemrograman secara bertahap. Seluruh program yang diuji juga berhasil berjalan dengan baik dan sesuai dengan hasil yang diharapkan.]