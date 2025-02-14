// Package belajar_golang_goroutines berisi contoh penggunaan goroutine di Go
package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// TestAtomic menguji penggunaan atomic operation untuk menghindari race condition
func TestAtomic(t *testing.T) {
	// Inisialisasi variabel counter atomic
	var x int64 = 0
	// Membuat WaitGroup untuk sinkronisasi goroutine
	group := sync.WaitGroup{}

	// Membuat 1000 goroutine yang akan menambah nilai counter
	for i := 1; i <= 1000; i++ {
		// Menambah counter WaitGroup sebelum membuat goroutine baru
		group.Add(1)
		// Membuat goroutine menggunakan anonymous function
		go func() {
			// Menandai bahwa goroutine telah selesai saat fungsi berakhir
			defer group.Done()
			
			// Setiap goroutine akan menambah counter sebanyak 100 kali
			for j := 1; j <= 100; j++ {
				// Menggunakan atomic.AddInt64 untuk menghindari race condition
				atomic.AddInt64(&x, 1)
			}
		}()
	}
	
	// Menunggu semua goroutine selesai
	group.Wait()
	// Menampilkan hasil akhir counter
	fmt.Println("Counter = ", x)
}
