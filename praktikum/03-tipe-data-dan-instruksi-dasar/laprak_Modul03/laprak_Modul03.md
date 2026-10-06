# <h1 align="center">Laporan Praktikum Modul [02] - [Bahasa Pemrograman GO/Variabel dan Operator]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
[Bahasa pemrograman Go (Golang) merupakan bahasa pemrograman yang dikembangkan oleh tiga insinyur Google yaitu Robert Griesemer, Rob
Pike, dan Ken Thompson pada tahun 2007 dan dirilis secara publik pada tahun 2009. Bahasa pemrograman ini dibuat dengan tujuan menyederhanakan pengembangan perangkat lunak tanpa harus mengorbankan performa dan kecepatan kompilasi yang optimal. Bahasa pemrograman ini dirancang dengan sintaks yang minimalis, sederhana juga adanya fitur _garbage collection_ sehingga membuat bahasa ini ramah untuk pemula.]


#### 1. [Variabel]
[Variabel merupakan suatu lokasi untuk menyimpan sebuah nilai. Sedangkan deklarasi variabel merupakan proses pengenalan sebuah nilai variabel dan tipe datanya kepada program komputer sebelum data tersebut digunakan. Dalam bahasa pemrograman Go, penulisan deklarasi variabel cukup ketat dikarenakan kita perlu menulis tipe data pasti untuk sebuah variabel, selain itu dalam bahasa pemrograman GO kita bisa mendeklarasikan variabel dengan cara yang berbeda (menggunakan var atau :=) tergantung kebutuhan struktur kode kita. Dalam bahasa pemrograman GO, variabel dikelompokkan berdasarkan cara pendeklarasiannya, ruang lingkupnya (scope), serta tipe data yang disimpan. 

Jenis-jenis variabel berdasarkan cara pendeklarasiannya dapat dilakukan dengan kata kunci var (_manifest typing_) yang dimana pendeklarasian variabel dilakukan dengan memasukkan tipe data variabel tersebut. Selain itu, ada juga deklarasi tanpa tipe data (_type inference_) yang dimana tipe data ini ditentukan secara otomatis ole GO berdasarkan nilai yang diberikan. Lalu terakhir, ada cara pendeklarasian variabel menggunakan operator := (_short operator declaration_) cara ini hanya bisa dilakukan di dalam blok fungsi.

Selanjutnya, jenis-jenis variabel berdasarkan ruang lingkup (scope) dapat dilakukan dengan variael tingkat paket yang dimana variabel ini didekalarasikan di luar fungsi dan ditempatkan di atas file setelah package, variabel ini dapat diakses oleh semua fungsi dalam satu paket dan harus menggunakan kata kunci var. Lalu, ada juga variabel tingkat blok yang dimana variabel ini dideklarasikan dalam blok kode, seperti dalam fungsi, perulangan for atau kondisi if dan biasanya menggunakan operator :=.

Lalu terakhir jenis-jenis variabel berdasarkan tipe data yang disimpan dapat digunakan berdasarkan beberapa jenis tipe data, contoh seperti jenis tipe data numerik (int = untuk bilangan bulat, floating point = untuk bilangan desimal), tipe data boolean yang menyimpan nilai kebenaran _true_ atau _false_, tipe data string untuk menyimpan teks atau kumpulan karakter dalam kutip gnda ("") atau backtick (```), dan tipe data kompleks atau koleksi (array & slice, map, channel, pointer & interface).]

#### 2. [Operator]
[Operator adalah simbol atau tanda khusus yang digunakan untuk melakukan operasi tertentu terhadap suatu nilai atau variabel. Terdapat lima operator dalam bahasa pemrograman GO, pertama ada opeartor aritmatika yang imana operator ini berfungsi untuk melakukan sebuah perhitungan matematika dasar (penjumlahan (+), pengurangan (-), perkalian (*), pembagian (/), hasil sisa %). 

kedua, yaitu operator perbandingan yang digunakan untuk membandingkan dua nilai yang akan menghasilkan boolean (sama dengan (==), tidak sama dengan (!=), lebih kecil dari (<), lebih besar dari (>), lebih kecil atau sama dengan (<=), lebih besar atau sama dengan (>=)).

Ketiga, operator logika untuk menggabungkan beberapa ekspresi boolean (logika AND (&&), logika OR (||), logika NOT (!)).

Keempat, operator penugasan untuk memberikan nilai ke dalam suatu variabel (pengisian nilai standar (=), deklarasi variabel pendek dan pengisian nilai otomatis (=:), Gabungan operator aritmatika/bit dengan assignment seperti (+=, _=, *=, /:))

Terakhir, operator bitwise digunakan untuk melakukan operasi pada level bit (biner) dari data numerik (&, |, ^, <<, >>).]




## Guided

### 1. [konversi.go]

```go
[package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu: ")
	fmt.Scanln(&celsius)

	fmt.Println("Kelvin ", celsius + 273)
}
]
```

#### Output
[Screenshoot Output Guided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/konversi/konversi.png)


#### Deskripsi
[Dalam konversi.go, soal mmerintahkan untuk menkoversi suhu celcius menjadi suhu kelvin. Dalam penginputannya, baris program ini menggunakan variabel celsius dengan tipe data float64 untuk memasukkan nilai desimal, lalu penggunaan fmt.Print untuk meminta  pengguna menginput nilai suhu, lalu (celsius + 273) untuk melakukan perhitungan dengan rumus konversi dari Celsius ke Kelvin secara matematis, lalu terakhir fmt.Println untuk mencetak hasil akhir kode yang sudah dibuat. Hasil akhir dari baris program ini adalah program dapat menghtiung dan mencetak hasil konversi satuan Celsius dan satuan Kevin dengan akurat.]

### 2. [tukarnilai.go]

```go
[package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}
]
```

#### Output
[Screenshoot Output Guided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/konversi/tukarnilai/tukarnilai.png)

#### Deskripsi
[Dalam tukarnilai.go, soal memerintahkan untuk menukar suatu nilai bilangan bulat x, y, z dengan ketentuan nilai y berisi nilai x, nilai x berisi nilai z, dan nilai z berisi nilai y. Dalam penginputannya, baris program ini memnggunaka variabl x, y, z dengan tipe data interger, fmt.Scan untuk membaca input secara berurutan, lalu melakukan pertukaran nilai dengancara nilai x ditaruh sementara kedalam variabel temp agar tidak hilang, lalu menukar nilai x menjadi nilai z (x = z), menukar nilai z menjadi y (z = y), dan yang terakhir menukar nilai y dengan x yang tadi disimpan di temp (y = temp), setelah itu, fmt.println untuk memcetak ouput pertukaran nilai tersebut. Hasil akhir dari baris program ini adalah program dapat menukar nilai variabel x, y,x dengan urut dan akurat.]

### 3. [cacahuang.go]
```go
[package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var sepuluhribuan int = x / 10000
	var sisa int = x % 10000

	var limaribuan int = sisa / 5000
	sisa = sisa % 5000

	var seribuan int = sisa / 1000

	fmt.Println("Sisa ", sepuluhribuan, limaribuan, seribuan)
}
]
```

#### Output
[Screenshoot Output Guided]![alt text](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/konversi/cacahuang/cacahuang.png)

#### Deskripsi
[Dalam cacahuang.go, soal memerintahkan untuk mencacah uang kembalian berupa lembar uang sepuluh ribu, lima ribuan, dan seribuan. Dalam penginputannya, baris program ini menggunakan variabel x dengan tipe data interger, lalu melakukan operasi hitung dengan operator pembagian (/) dan operator modulus (%), setelah itu fmt.Println untuk mencetak output pertukaran nilai uang tersebut. Hasil akhir dari baris program ini adalah program ini dapat mencacah uang kembalian dalam lembar uang sepuluh ribu, lima ribu dan seribu beserta sisanya dengan akurat.]


<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [reamur.go]
```go
[package  main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)
	reamur := celsius * 4 / 5
	fmt.Printf("%.2f derajat Celsius sama dengan %.2f derajat Reamur\n", celsius, reamur)

}
]
```


##### Output
[Screenshoot Ouput Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/reamur/reamur.png)



#### Deskripsi
[Dalam reamur.go, soal memerintahkan untuk mengkonversi nilai Celsius menjadi nilai Reamur. Dalam penginputannya baris program ini menggunakan variabel Celsius dengan tipe data float64 untuk memasukkan nilai desimal, fmt.Scanln untuk membaca input yang dimasukkan pengguna, lalu melakukan operasi hitung konversi nilai satuan Celsius ke nilai satuan Reamur dengan rumus (Celsius * 4 / 5), lalu yang terakhir fmt.Printf untuk mencetak output konversi satuan Celsius ke satuan Reamur dengan menggunakan format %2f  untuk menampilkan angka desimal dengan membulatkannya menjadi 2 angka di belakang koma agar terlihat rapi. Hasil akhir dari baris program ini adalah program ini dapat mengkonversi nilai satuan Clesius ke satuan Reamur dengan akurat. ]

### 2. [jumlahhari.go]

```go
[package main

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
]
```

##### Output
[Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/jumlahhari/jumlahhari.png)

#### Deskripsi
[Dalam soal jumlahhari.go, soal memerintahkan untuk mengkonversi jumlah hari ke dalam satuan tahun, bulan, minggu, dan hari. Dalam penginputannya, baris program ini menggunakan variabel totalHari dengan tipe data interger, fmt.Scan untuk membaca input dari pengguna, lalu melakukan operasi hitung konversi dan sisa totalHari ke tahun dengan rumus (totalHari / 360, totalHari % 360), konversi dan sisa sisalHari ke bulan (sisaHari / 30. sisaHari % 30), konversi sisaHari ke minggu (sisaHari / 7) dan perhitungan sisa hari dengan rumus (sisaHari % 7), setelah itu fmt.Println untuk menampilkan output konversi total hari ke tahun, bulan, minggu dan hari. Hasil akhir dari baris program ini adalah program dapat menjalankan dan menkonversi total hari ke tahun, bulan, minggu dan hasi beserta sisa harinya dengan akurat.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Praktikum modul ini membantu memahami dasar-dasar penggunaan bahasa pemrograman Go, mulai dari mendeklarasikan variabel, menggunakan tipe data yang tepat, hingga memanfaatkan operator aritmatika dan logika. Melalui latihan membuat program konversi suhu, pertukaran nilai, pencacahan uang, hingga konversi satuan waktu, dapat disimpulkan bahwa struktur kode yang rapi serta pemilihan fungsi input-output seperti fmt.Scan dan fmt.Println sangat penting agar program dapat menerima data, memproses perhitungan matematika dengan akurat, dan menampilkan hasil akhir yang sesuai dengan kebutuhan. Selain itu, Praktikum modul ini juga membantu memahami dasar-dasar penggunaan bahasa pemrograman Go, mulai dari pentingnya mendeklarasikan variabel dengan tipe data yang tepat—baik menggunakan kata kunci var maupun operator := hingga memanfaatkan berbagai operator seperti aritmatika dan penugasan untuk mengolah data. Melalui latihan membuat program konversi suhu.]

## Referensi
1. Kontributor Komunitas Go. “Spesifikasi Bahasa Pemrograman Go: Variables.” Golang-ID, 2026, https://golang-id.org/ref/spec/#Variables. Diakses 6 Oktober 2026.
2. BuildWithAngga. “Dasar-Dasar Bahasa Pemrograman Go: Variabel, Tipe Data, dan Operasi Dasar.”  17 Feb. 2024, https://buildwithangga.com/tips/dasar-dasar-bahasa-pemrograman-go-variabel-tipe-data-dan-operasi-dasar. Diakses 6 Oktober 2026.
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
