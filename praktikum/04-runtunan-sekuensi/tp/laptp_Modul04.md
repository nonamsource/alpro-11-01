# <h1 align="center">Tugas Pendahuluan Modul [04] - [Bahasa Pemrograman GO/Runtunan atau Sekuensi]</h1>
<p align="center">[Nadhifa Khumaira Zalfa] - [109092630004]</p>

### 1. Evaluasi EKspresi Kontrol dalam GO

```go
[package main

import "fmt"

func main () {
	intNum := 5 
	intOther := 10 
	var sngNum float64 = -3.0
	fmt.Println(intOther + 2 * intNum != 30 || !(sngNum > 0))
	fmt.Println(4 / 2 == intOther / intNum )
	fmt.Println(intOther + 2 * intNum != 30 || !(sngNum > 0))
	fmt.Println(intNum > 0 || (sngNum <= 0 && intOther == 13))
}]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/ekspresi/ekspresi.png)


#### Deskripsi
[Pada bagian ini, praktikum difokuskan untuk memahami cara kerja operator aritmatika, perbandingan, dan logika dalam bahasa pemrograman Go. Di sini saya mencoba mengevaluasi beberapa ekspresi boolean yang cukup kompleks dengan menggabungkan variabel bertipe integer (intNum, intOther) dan float (sngNum). Melalui latihan ini, saya jadi lebih paham bagaimana Go memprioritaskan urutan eksekusi antar operator, mulai dari operasi aritmatika dasar, perbandingan relasional, hingga operator logika seperti && (AND), || (OR), dan ! (NOT) sehingga menghasilkan nilai akhir boolean yang sesuai dengan logika program.]

### 2. [Tracing]

```go
[package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z 
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}
	fmt.Println("Nilai akhir result:", result)
}
]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/tracing/tracing.png)


#### Deskripsi
[Latihan ini berfokus pada proses tracing atau penelusuran alur eksekusi logika percabangan bertingkat (nested if-else) menggunakan beberapa variabel integer (x, y, z) serta sebuah variabel penampung result. Selama proses pengerjaan, saya melacak langkah demi langkah bagaimana nilai result berubah di setiap blok kondisi, mulai dari penambahan dasar, operasi aritmatika lanjutan, hingga percabangan alternatif dengan operator penugasan gabungan. Kegiatan ini sangat membantu saya dalam melatih ketelitian membaca alur program yang bercabang sekaligus memahami bagaimana suatu kondisi menentukan jalur eksekusi kode di Go.]

### 3. [Jumlah hari]
```go
[package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun)
	fmt.Scan(&bulan)

	switch bulan {
		case "Januari", "Maret", "Mei", "Juli", "Agustus", "Oktober", "Desember":
			fmt.Println("Jumlah hari:", 31)
		case "April", "Juni", "September", "November":
			fmt.Println("Jumlah hari:", 30)
		case "Februari":
			if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
				fmt.Println("Jumlah hari:", 29)
			} else {
				fmt.Println("Jumlah hari:", 28)
			}
		default:
			fmt.Println("Bulan tidak valid")
	}
}
]
```

#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/jumlahhari/jumlahhari%2004.png)

#### Deskripsi
[Pada soal ketiga ini, praktikum membahas implementasi struktur kontrol percabangan switch-case untuk menentukan jumlah hari dalam suatu bulan berdasarkan tahun yang diinputkan. Program dibuat agar mampu mengelompokkan bulan-bulan yang memiliki 31 hari, 30 hari, serta menangani kasus khusus untuk bulan Februari yang jumlah harinya dipengaruhi oleh kondisi tahun kabisat (kelipatan 4, dengan aturan pengecualian pada tahun abad). Hasil dari praktikum ini menunjukkan bahwa penggunaan switch-case yang dikombinasikan dengan kondisi if di dalamnya membuat kode program menjadi jauh lebih rapi, terstruktur, dan mudah dibaca ketimbang menggunakan if-else yang terlalu panjang.]

### 4. Switch
```go
[package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Bukan 1, 2, atau 3")
	}
}
]
```

#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/switch/switch.png)

#### Deskripsi
[Bagian ini mempelajari penggunaan dasar dari struktur kontrol switch-case sederhana untuk mencetak nama bilangan berdasarkan input angka integer dari pengguna. Program membaca nilai variabel x lalu mencocokkannya dengan beberapa pilihan case yang tersedia (1, 2, atau 3), serta menyediakan blok default untuk mengantisipasi jika nilai input tidak sesuai dengan pilihan yang ada. Dari latihan ini, saya dapat memahami sintaks dasar percabangan banyak arah di Go yang membuat proses seleksi kondisi menjadi lebih bersih dan efisien dibandingkan percabangan if-else bertingkat.]


## Kesimpulan
[Secara keseluruhan, praktikum kali ini memberikan pemahaman yang mendalam mengenai berbagai macam struktur kontrol percabangan dan evaluasi ekspresi dalam bahasa pemrograman Go. Dari rangkaian latihan yang telah dikerjakan, mulai dari evaluasi ekspresi logika, penelusuran alur program (tracing), hingga penerapan percabangan menggunakan if-else dan switch-case saya belajar bagaimana menyusun logika pemrograman yang sistematis untuk menyelesaikan berbagai kasus komputasi yang berbeda. Pemahaman ini sangat penting sebagai fondasi dasar dalam membangun alur program yang logis, efisien, dan bebas dari kesalahan logika ke depannya.]