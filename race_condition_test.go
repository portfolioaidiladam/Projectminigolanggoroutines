// Package belajar_golang_goroutines berisi contoh implementasi goroutine di Go
package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// TestRaceCondition adalah fungsi test yang mendemonstrasikan masalah race condition
// dimana multiple goroutine mencoba mengakses dan memodifikasi variabel yang sama
// secara bersamaan, yang dapat menyebabkan hasil yang tidak terduga
func TestRaceCondition(t *testing.T) {
	// x adalah variabel counter yang akan diakses oleh multiple goroutine
	x := 0

	// Membuat 1000 goroutine yang masing-masing akan menambahkan nilai x
	for i := 1; i <= 1000; i++ {
		go func() {
			// Setiap goroutine akan menambahkan nilai x sebanyak 100 kali
			for j := 1; j <= 100; j++ {
				// PERINGATAN: Ini adalah operasi yang tidak aman (race condition)
				// karena multiple goroutine mencoba mengubah nilai x secara bersamaan
				x = x + 1
			}
		}()
	}

	// Menunggu 5 detik untuk memberi waktu goroutine menyelesaikan eksekusinya
	// CATATAN: Menggunakan time.Sleep bukanlah cara yang baik untuk sinkronisasi
	// Lebih baik menggunakan WaitGroup atau channel untuk koordinasi goroutine
	time.Sleep(5 * time.Second)
	
	// Mencetak nilai akhir dari x
	// CATATAN: Hasil mungkin tidak akan konsisten karena adanya race condition
	fmt.Println("Counter = ", x)
}
