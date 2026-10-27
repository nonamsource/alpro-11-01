# <h1 align="center">Laporan Praktikum Modul [06] - [Bahasa Pemrograman GO/Perulangan while loop dan repeat]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>

## Dasar Teori

### A. [Bahasa Pemrograman Go]
[Bahasa pemrograman Go (Golang) merupakan bahasa pemrograman yang dikembangkan oleh tiga insinyur Google yaitu Robert Griesemer, Rob
Pike, dan Ken Thompson pada tahun 2007 dan dirilis secara publik pada tahun 2009. Bahasa pemrograman ini dibuat dengan tujuan menyederhanakan pengembangan perangkat lunak tanpa harus mengorbankan performa dan kecepatan kompilasi yang optimal. Bahasa pemrograman ini dirancang dengan sintaks yang minimalis, sederhana juga adanya fitur _garbage collection_ sehingga membuat bahasa ini ramah untuk pemula.]


#### 1. [Perulangan while loop dan repeat]
[perulangan while loop dan repeat merupakan dua jenis perulangan yang digunakan untuk menjalankan sekumpulan blok kode secara berulang. Dalam bahasa pemrograman Go, penggunaan while loop dan repeat ini disederhanakan menggunakan perintah perulangan for. Untuk membuat konsep while loop di Go, kita menggunakan bentuk for dengan hanya menyertakan satu kondisi saja di depannya (tanpa inisialisasi awal atau penambahan nilai di dalam format standar). Perulangan ini akan terus berjalan selama kondisi tersebut bernilai benar (true), contohnya seperti for i <= 10 { ... }.

Sementara itu, untuk meniru konsep repeat, Go biasanya menggunakan perulangan tak terbatas (for { ... }) di mana blok kode di dalamnya dipastikan berjalan lebih dulu minimal satu kali, lalu di bagian akhir atau tengah blok diberikan pengecekan kondisi menggunakan if beserta perintah break untuk menghentikan perulangan jika syarat tertentu sudah terpenuhi.] 




## Guided

### 1. [terkecil.go]

```go
[package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}
]
```

#### Output
[Screenshoot Output Guided](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-1/output%20terkecil.png)


#### Deskripsi
[Dalam soal terkecil.go program diperintah untuk mencari bilangan bulat positif terkecil yang pangkat tiga-nya melebihi batas angka tertentu yang dimasukkan oleh pengguna, dengan memanfaatkan variabel berjenis data int (bilangan bulat). Variabel batas digunakan untuk menyimpan nilai maksimal dari input pengguna, sedangkan variabel n (yang diinisialisasi mulai dari angka 1) bertindak sebagai pencari atau penghitung angka yang akan dipangkatkan tiga (n x n x n). Cara kerja program ini dimulai dengan membaca nilai batas, kemudian menjalankan perulangan (for) selama hasil pangkat tiga dari n masih kurang dari atau sama dengan batas, di mana pada setiap iterasi nilai n akan terus bertambah satu secara berkala; begitu hasil pangkat tiganya melampaui nilai batas, perulangan berhenti dan program mencetak nilai n terakhir tersebut ke layar.]

### 2. [terbesar.go]

```go
[package main

import "fmt"

func main() {
	var batas int
	fmt.Scan(&batas)

	n := 1
	for n*n < batas {
		n++
	}

	fmt.Println(n - 1)
}	
]
```

#### Output
[Screenshoot Output Guided](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-2/output%20terbesar%206.png)

#### Deskripsi
[Dalam soal terbesar.go, program diperintah untuk mencari nilai bilangan bulat terbesar yang jika dikuadratkan (n x n) hasilnya masih kurang dari batas angka tertentu yang dimasukkan oleh pengguna, dengan memanfaatkan variabel berjenis data int (bilangan bulat). Variabel batas digunakan untuk menyimpan nilai batas maksimal dari input pengguna, sedangkan variabel n (yang dimulai dari angka 1) bertindak sebagai pencari atau kandidat angka yang akan dikuadratkan secara berurutan. Cara kerja program ini dimulai dengan membaca nilai batas, lalu menjalankan perulangan (for) selama hasil kuadrat dari n masih kurang dari batas, di mana nilai n akan terus bertambah satu pada setiap iterasi; begitu perulangan berhenti karena kuadratnya sudah tidak lagi kurang dari batas, program akan mencetak hasil n - 1 ke layar, yang merepresentasikan angka bulat terbesar yang kuadratnya tidak melebihi batasan tersebut.]

### 3. [vokal.go]
```go
[package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}
]
```

#### Output
[Screenshoot Output Guided]![alt text](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-3/output%20vokal.png)

#### Deskripsi
[Dalam soal vokal.go, program diperintah untuk menghitung jumlah huruf vokal (a, i, u, e, o) dalam kalimat input pengguna. Program menggunakan variabel kalimat (string) untuk teks, jumlahVokal dan i (int) untuk akumulasi hitungan dan indeks posisi karakter, serta ch (string) untuk memeriksa tiap huruf. Yang dimana cara kerjanya program membaca kalimat, lalu melakukan perulangan untuk mengecek setiap karakter satu per satu (diubah menjadi huruf kecil), menyeleksinya dengan percabangan if untuk mencocokkannya dengan huruf vokal, dan terakhir mencetak total huruf vokal yang ditemukan.]


<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [atm.go]
```go
[package main

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
]

```


##### Output
[Screenshoot Ouput Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/unguided/ATM/output%20atm.png)



#### Deskripsi
[Dalam soal atm.go, program diperintah untuk mensimulasikan sistem mesin ATM sederhana yang mencakup fitur login dengan batasan percobaan password dan menu utama untuk bertransaksi.

Program ini menggunakan variabel saldo dan kesempatan (int) untuk menyimpan jumlah uang awal sebesar 10.000.000 dan sisa jatah salah password sebanyak 3 kali, pwd (string) untuk input password, serta pilih dan tarik (int) untuk navigasi menu serta nominal penarikan uang.

Cara kerjanya terbagi menjadi dua bagian: pertama, perulangan login di mana pengguna diminta memasukkan password "1234"—jika benar, program lanjut; jika salah, kesempatan berkurang hingga habis dan ATM terblokir. Kedua, jika login berhasil, perulangan menu tak terbatas akan aktif untuk menampilkan pilihan cek saldo (menu 1), tarik tunai dengan pengecekan ketersediaan saldo (menu 2), atau keluar dari sistem (menu 3).]



## Kesimpulan
[Dari keempat kode program Go yang telah dibahas (pencari skor terkecil, kuadrat terbesar, penghitung huruf vokal, dan simulasi ATM), dapat disimpulkan bahwa seluruh program tersebut memperlihatkan penerapan konsep-konsep dasar pemrograman untuk menyelesaikan berbagai kebutuhan secara terstruktur. Keempatnya mengombinasikan penggunaan variabel dan tipe data yang sesuai seperti int untuk angka bulat, string untuk teks atau password, serta fungsi pemrosesan teks khusus (bufio dan strings) untuk menyimpan dan mengolah informasi dengan tepat. Konsep perulangan (for) diterapkan secara fleksibel untuk berbagai keperluan, seperti menyeleksi data masukan secara berulang, mencari nilai batas, memeriksa karakter teks satu per satu, hingga menjaga menu interaktif tetap aktif. Logika ini diperkuat dengan percabangan (if-else) yang memungkinkan program mengambil keputusan secara otomatis, seperti membandingkan skor untuk menemukan nilai paling minimum, memvalidasi huruf vokal, memverifikasi keamanan password, hingga memproses transaksi tarik tunai berdasarkan ketersediaan saldo. Secara keseluruhan, keempat kodingan tersebut menjadi gambaran bagaimana logika dasar pemrograman digunakan untuk menerima masukan pengguna, mengolahnya melalui perhitungan dan seleksi kondisi, lalu menghasilkan keluaran yang berguna dan akurat.]

## Referensi
1. Agung, N. (n.d.). "Dasar Pemrograman Golang". Diakses dari https://dasarpemrogramangolang.novalagung.com/A-perulangan.html
2. Abdulmalik, H. I. "Apa itu Looping? Arti, Fungsi, dan Contohnya". Diakses dari https://www.dicoding.com/blog/apa-itu-looping-arti-fungsi-dan-contohnya/
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
