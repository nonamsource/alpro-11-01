# <h1 align="center">Laporan Praktikum Modul [03] - [Bahasa Pemrograman GO/Runtunan dan Sekuensi]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
[Bahasa pemrograman Go (Golang) merupakan bahasa pemrograman yang dikembangkan oleh tiga insinyur Google yaitu Robert Griesemer, Rob
Pike, dan Ken Thompson pada tahun 2007 dan dirilis secara publik pada tahun 2009. Bahasa pemrograman ini dibuat dengan tujuan menyederhanakan pengembangan perangkat lunak tanpa harus mengorbankan performa dan kecepatan kompilasi yang optimal. Bahasa pemrograman ini dirancang dengan sintaks yang minimalis, sederhana juga adanya fitur _garbage collection_ sehingga membuat bahasa ini ramah untuk pemula.]


#### 1. [Runtunan / Sekuensi]
[Runtunan atau sekuensi adalah satu atau lebih intruksi yang dieksekusi baris demi baris secara berurutan dari atas ke bawah. Runtunan atau sekuensi merupakan salah satu langkah pembentuk konstruksi suatu algoritma sama seperti pemilihan dan pengulangan. Di dalam GO kita tidak perlu menggunakan tanda ; (titik koma) untuk mengganti perintah, karena di setiap baris baru GO sudah secra otomatis menjalankan perintah yang berbeda. Selain itu, pemindahan baris ini juga perlu diperhatikan penggunaan kurung kurawal buka ({). Tanda ({) harus ditaruh di akhir baris perintah, bukan di baris baru supaya tidak terjadi error. Error tersebut dapat terjadi karena GO akan menganggap bahwa program baris sebelumnya sudah selesai dan berujung error.] 






## Guided

### 1. [grade.go]

```go
[package main

import (
	"bufio"
	"os"

	"fmt"
)

func main(){
	var nama string 
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}	
]
```

#### Output
[Screenshoot Output Guided](/praktikum/04-runtunan-sekuensi/guided/grade/grade.png)


#### Deskripsi
[Alur kerja grade.go dimulai dari atas ke bawah secara berurutan (sekuensi). Pertama-tama, program menyiapkan tiga variabel, yaitu nama (string), angka nilai(float64), dan huruf grade(string). Setelah itu, program akan memunculkan perintah bagi pengguna untuk mengetikkan nama terlebih dahulu, lalu dilanjutkan dengan mengetikkan angka nilainya. Begitu kedua data tersebut diterima, komputer langsung melakukan pengecekan lewat sistem percabangan (if-else). Jika nilainya antara 90 sampai 100 maka otomatis mendapat "A", jika 80 sampai 89 mendapat "B", nilai 70 sampai 79 mendapat "C", nilai 60 sampai 69 mendapat "D", dan di bawah itu langsung dianggap gagal dengan nilai "F". Setelah semua proses hitung-menghitung selesai di latar belakang, baris terakhir program akan langsung menampilkan kesimpulan hasilnya ke layar komputer dengan format kalimat seperti "Hong Gil-dong mendapatkan nilai A".]

### 2. [klasifikasi.go]

```go
[package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
]
```

#### Output
[Screenshoot Output Guided](/praktikum/04-runtunan-sekuensi/guided/klasifikasi/klasifikasi.png)

#### Deskripsi
[Dalam klasifikasi.go, komputer mula-mula menyiapkan tiga variabel untuk mencatat angka usia (interger), angka gaji (interger), serta teks keterangan hasil akhir (string). Selanjutnya, program akan menunggu pengguna mengetikkan dua angka secara bergantian, yaitu usia terlebih dahulu kemudian jumlah gajinya. Setelah kedua angka tersebut didapatkan, program langsung mengecek menggunakan sistem evaluasi switch dari atas ke bawah untuk mencari kondisi yang paling cocok. Seseorang yang berusia di bawah 18 tahun akan langsung dilabeli "Masih sekolah", sedangkan untuk usia dewasa, program akan membagi mereka menjadi kelompok "Muda sukses", "Masih belajar hidup", "Pekerja mapan", "Perlu perbaikan karier", "Profesional berpengalaman", atau "Perlu evaluasi finansial" tergantung dari besaran gaji yang mereka punya. Di baris paling akhir, komputer akan langsung mencetak satu teks keterangan hasil klasifikasi tersebut ke layar.]

### 3. [penilaian.go]
```go
[package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
]
```

#### Output
[Screenshoot Output Guided]![alt text](/praktikum/04-runtunan-sekuensi/guided/penilaian/penilaian.png)

#### Deskripsi
[Kode program dimulai dengan menyiapkan empat variabel untuk menampung nomor pilihan menu (interger), nama siswa (string), angka nilai(float64), dan huruf grade (string). Selanjutnya, komputer akan menampilkan teks pilihan menu di layar dan menunggu pengguna mengetikkan angka. Jika pengguna memilih angka 1, program akan masuk ke menu penilaian, lalu meminta pengguna mengetikkan nama lengkap serta angka nilai siswa secara berurutan. Setelah data nama dan nilai diterima, komputer melakukan evaluasi bertingkat (if-else) untuk menentukan grade dari "A" hingga "F" sebelum akhirnya mencetak kalimat kesimpulan seperti "Hong Gil-dong mendapatkan nilai A". Namun, jika pengguna mengetikkan angka 0 sejak awal, program akan langsung menampilkan pesan penutup dan menyudahi tugasnya.]


<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [grade2.go]
```go
[package main

import (
	"bufio"
	"os"

	"fmt"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "E"
	}
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
]
```


##### Output
[Screenshoot Ouput Unguided](/praktikum/04-runtunan-sekuensi/unguided/grade2/grade2.png)



#### Deskripsi
[Sama seperti grade.go di tugas guided, program dimulai dengan menyiapkan tiga variabel untuk nama (string), angka nilai(float64), dan grade (string). Selanjutnya, program akan menunggu pengguna mengetikkan nama siswa terlebih dahulu, lalu dilanjutkan dengan mengetikkan angka nilainya pada baris berikutnya. Begitu kedua data tersebut siap, komputer akan melakukan pengecekan lewat sistem percabangan bertingkat (if-else). Jika nilai berada di antara 90 hingga 100 maka siswa mendapat "A", nilai 80 hingga 89 mendapat "B", nilai 70 hingga 79 mendapat "C", nilai 60 hingga 69 mendapat "D", dan angka di bawah itu otomatis langsung dianggap gagal dan mendapatkan nilai "E". Setelah proses evaluasi selesai, baris terakhir program akan langsung menampilkan kalimat kesimpulan hasil akhirnya ke layar dengan format seperti "Lee Haechan mendapatkan nilai B".]

### 2. [Pajak.go]


```go
[package main

import "fmt"

func main() {
	var penghasilan, pajak float64

	fmt.Println("Masukkan penghasilan (dalam juta): ")
	fmt.Scan(&penghasilan)

	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan > 50 && penghasilan <= 100 {
		pajak = (0.05 * 50) + (0.1 * (penghasilan - 50))
	} else if penghasilan > 100 && penghasilan <= 200 {
		pajak = (0.05 * 50) + (0.1 * 50) + (0.15 * (penghasilan - 100))
	} else {
		pajak = (0.05 * 50) + (0.1 * 50) + (0.15 * 100) + (0.20 * (penghasilan - 200))
	}

	fmt.Printf("Pajak yang harus dibayarkan: %.2f juta\n", pajak)
}
]
```

##### Output
[Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/pajak/pajak.png)

#### Deskripsi
[Komputer mula-mula menyiapkan dua tempat variabel berupa angka desimal untuk mencatat jumlah penghasilan dan hasil hitung pajak. Selanjutnya, program akan memunculkan teks perintah di layar dan menunggu pengguna mengetikkan nominal penghasilannya dalam satuan juta. Setelah angka tersebut diterima, program langsung melakukan perhitungan berlapis menggunakan sistem percabangan (if-else). Aturan pajaknya dibagi menjadi beberapa tingkatan: penghasilan sampai 50 juta dikenakan tarif 5%, potongan 50 juta pertama ditambah 10% dari sisanya untuk penghasilan hingga 100 juta, potongan bertingkat ditambah 15% dari sisanya untuk penghasilan hingga 200 juta, serta tambahan tarif 20% bagi penghasilan di atas 200 juta. Di baris paling akhir, komputer akan menghitung total kumulatifnya dan mencetak hasil potongan pajak tersebut ke layar dengan dua angka di belakang koma.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Kesimpulan dari materi dan latihan soal di atas adalah bahwa runtunan atau sekuensi itu merupakan aturan antrean di dalam program Go, di mana komputer akan membaca dan menjalankan perintah kita baris demi baris, dari paling atas sampai paling bawah.
Aturan ini membuat cara kerja program jadi sangat teratur dan tidak bisa melompati perintah sembarangan. Penulisan runtunan di GO juga simpel karena kita tidak perlu repot mengetik tanda titik koma (;) setiap kali ganti perintah. Melalui ketiga program yang sudah dibuat—mulai dari sistem cek nilai, pengelompok status hidup, sampai menu interaktif—kita bisa melihat kalau aturan jalan lurus (sekuensi) ini akan selalu jadi fondasi utama yang bekerja sama dengan logika pilih-pilih jalan (if-else atau switch).]

## Referensi
1. Sari, Indah Purnama, "Algoritma dan Pemrograman." Medan: UMSU Press, 2024. Tersedia: https://books.google.co.id/books?id=UGI3EQAAQBAJ. Diakses 13 Oktober 2026
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
