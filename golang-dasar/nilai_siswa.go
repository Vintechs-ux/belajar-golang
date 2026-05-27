package main

import "fmt"

func hitungRataRata(nilai map[string]int) float64 {
	total := 0

	for _, v := range nilai {
		total += v
	}

	return float64(total) / float64(len(nilai))
}

func main() {
	fmt.Println("Nilai Siswa — Radit")

	Nama := map[string]int{
		"Matematika":       100,
		"Biologi":          99,
		"Bahasa Inggris":   98,
		"Kimia":            97,
		"Bahasa Indonesia": 90,
		"Informatika":      100,
		"Bahasa Jawa":      70,
	}

	// Tampilkan semua mata pelajaran + nilai:
	for mapel, nilai := range Nama {
		fmt.Printf("%s: %d\n", mapel, nilai)
	}

	// Hitung rata-rata:
	rata := hitungRataRata(Nama)
	fmt.Printf("\nRata-rata: %.2f\n", rata)
}
