package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// TestTicker menguji fungsionalitas time.Ticker untuk mengeksekusi kode secara periodik
// dengan interval waktu 1 detik dan berhenti setelah 5 detik.
func TestTicker(t *testing.T) {
	// Membuat ticker baru yang akan mengirim sinyal setiap 1 detik
	ticker := time.NewTicker(1 * time.Second)

	// Goroutine untuk menghentikan ticker setelah 5 detik
	go func() {
		// Menunggu selama 5 detik
		time.Sleep(5 * time.Second)
		// Menghentikan ticker untuk mencegah memory leak
		ticker.Stop()
	}()

	// Loop untuk menerima nilai waktu dari channel ticker
	// dan mencetak setiap timestamp yang diterima
	for time := range ticker.C {
		fmt.Println(time)
	}
}

// TestTick menguji fungsi time.Tick yang merupakan versi sederhana dari time.Ticker
// CATATAN: time.Tick tidak memiliki mekanisme untuk dihentikan, sehingga lebih cocok
// untuk program yang berjalan terus-menerus
func TestTick(t *testing.T) {
	// Membuat channel yang akan mengirim sinyal setiap 1 detik
	channel := time.Tick(1 * time.Second)

	// Loop untuk menerima nilai waktu dari channel
	// dan mencetak setiap timestamp yang diterima
	for time := range channel {
		fmt.Println(time)
	}
}
