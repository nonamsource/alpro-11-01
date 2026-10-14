# <h1 align="center">Tugas Pendahuluan Modul [05] - [Bahasa Pemrograman GO/Perulangan for Loop]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>


### 1.	Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)!

Fungsi utama for adalah mengeksekusi blok kode secara berulang berdasarkan kondisi tertentu pada elemen kumpulan data menggunakan klausa range.
Berikut tiga komponen utama perulangan for:
1.	Inisialisasi : Adalah suatu komponen yang berfungsi untuk mendeklarasikan suatu variable yang akan digunakan dalam proses perulangan berlangsung dan pernyataan ini hanya dieksekusi sekali di awal sebelum proses iterasi berlangsung.
2.	Kondisi : Adalah ekspresi logika berupa pernyataan Boolean (true) dan (false) yang bertindak sebagai kriteria evaluasi sebelum iterasi dimulai.
3.	Post statement : Adalah pernyataan yang digunakan untuk memodifikasi atau memperbarui nilai variabel penghitung setelah setiap iterasi selesai.


### 2. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing.

For dengan tiga komponen umumnya digunakan ketika jumlah iterasi atau perulanagn sudah diketahui sejak awal.
#### Contoh kode
```go
[package main

import "fmt"

func main() {
    // Menggunakan 3 komponen (inisialisasi; kondisi; post statement)
    for i := 0; i < 5; i++ {
        fmt.Println("Iterasi ke:", i)
    }
}
]
```
Sedangkan For menggunakan kondisi umumnya digunakan Ketika jumlah iterasi atau perulangan tidak pasti sejak awal atau berubah-ubah tergantung kondisi eksternal yang berubah dalam blok kode.
#### Contoh kode
```go
[package main

import "fmt"

func main() {
    // Inisialisasi dilakukan di luar
    i := 0

    // Hanya menggunakan kondisi (berfungsi seperti while loop)
    for i < 5 {
        fmt.Println("Nilai i:", i)
        
        // Post statement / pembaruan nilai dilakukan di dalam blok
        i++
    }
}
]
```

### 3. [Ganjil.go]

```go
[package main

import "fmt"

func main () {
	var n int

	fmt.Scan(&n)

	jumlah := 0
	first := true

	for i := 1; i <= n; i += 2 {
		if !first {
			fmt.Print(" ")
		}
		fmt.Print(i)
		jumlah += i
		first = false
	
	}

	fmt.Println()
	fmt.Println(jumlah)

}
]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/ganjil/ganjil.png)


#### Deskripsi
[Proses dalam ganjil.go ini diawali dengan menerima masukan nilai n, serta menginisialisasi variabel jumlah untuk menampung total akumulasi dan variabel first sebagai penanda pemformatan spasi agar cetakan angka terlihat rapi. Selanjutnya, program menjalankan perulangan for yang dimulai dari i = 1 dengan penambahan i += 2 selama i <= n untuk memastikan hanya bilangan ganjil yang diproses. Pada setiap iterasi, angka ganjil dicetak ke layar disertai spasi pemisah antarangka dan langsung ditambah ke dalam variabel jumlah. Setelah perulangan selesai, program berpindah ke baris baru lalu menampilkan hasil akhir berupa total akumulasi penjumlahan tersebut.]


### 3. [Perkalian.go]
```go
[package main

import "fmt"

func main() {
	var a int //bilangan yang dikalikan
	var b int //batas perkalian

	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}
}
]
```

#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/perkalian/perkalian.png)

#### Deskripsi
[Proses dalam perkalian.go ini dimulai dengan menerima dua buah nilai masukan, yaitu variabel a (interger) sebagai angka yang akan dikalikan dan variabel b (interger) sebagai batas akhir perkalian. Selanjutnya, program menjalankan perulangan for klasik yang dimulai dari penghitung i = 1 dan akan terus berjalan selama i <= b, dengan penambahan nilai i++ pada setiap putarannya. Di dalam blok perulangan tersebut, program menggunakan perintah fmt.Printf untuk menampilkan format operasi perkalian secara lengkap—mulai dari nilai a, angka pengali i, hingga hasil akhir dari perkalian a * i yang otomatis tercetak baris demi baris hingga batas b terpenuhi.]


## Kesimpulan
[Kesimpulan dari seluruh rangkaian materi dan latihan di atas adalah bahwa struktur perulangan for memegang peran fundamental dan fleksibel dalam bahasa pemrograman Go, mengingat bahasa ini tidak menyediakan kata kunci pengulangan lain seperti while atau do-while. Melalui bentuk umumnya yang mencakup tiga komponen yaitu inisialisasi, kondisi, dan post statement—pemrogram dapat mengatur perulangan dengan jumlah iterasi yang terstruktur secara pasti, maupun menyederhanakannya hanya dengan menyertakan kondisi tertentu untuk mengakomodasi pola perulangan yang dinamis. Fleksibilitas ini terbukti efektif dalam berbagai implementasi program praktis, seperti menghasilkan deret bilangan ganjil secara terurut beserta akumulasi penjumlahannya, maupun mencetak tabel perkalian secara sistematis berdasarkan masukan dari pengguna.]