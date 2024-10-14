package main

import "fmt"

func main() {

	hargaJual := 150000.0
	hargaBeli := 100000.0
	biayaOperasional := 1000.0
	diskon := 15.0
	jumlahTerjual := 100

	hargaJualSetelahDiskon := hargaJual * (1 - (diskon / 100))
	totalPendapatan := hargaJualSetelahDiskon * float64(jumlahTerjual)
	totalBiaya := (hargaBeli + biayaOperasional) * float64(jumlahTerjual)
	totalKeuntungan := totalPendapatan - totalBiaya

	fmt.Printf("Harga Jual Setelah Diskon: Rp%.2f\n", hargaJualSetelahDiskon)
	fmt.Printf("Total Pendapatan: Rp%.2f\n", totalPendapatan)
	fmt.Printf("Total Biaya: Rp%.2f\n", totalBiaya)
	fmt.Printf("Total Keuntungan: Rp%.2f\n", totalKeuntungan)
}
