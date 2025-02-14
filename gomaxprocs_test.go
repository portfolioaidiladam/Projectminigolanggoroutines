// Package belajar_golang_goroutines berisi implementasi testing untuk GOMAXPROCS
package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// TestGetGomaxprocs adalah fungsi test untuk mendemonstrasikan penggunaan GOMAXPROCS
// dan informasi terkait goroutine dan CPU
func TestGetGomaxprocs(t *testing.T) {
	// Inisialisasi WaitGroup untuk sinkronisasi goroutine
	group := sync.WaitGroup{}

	// Membuat 100 goroutine yang berjalan secara concurrent
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			// Simulasi pekerjaan dengan sleep selama 3 detik
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	// Mendapatkan jumlah CPU yang tersedia di sistem
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// Mendapatkan jumlah thread yang digunakan Go runtime
	// Parameter -1 berarti hanya mengambil nilai tanpa mengubah
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// Mendapatkan jumlah goroutine yang sedang berjalan
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	// Menunggu semua goroutine selesai
	group.Wait()
}

// TestChangeThreadNumber adalah fungsi test untuk mendemonstrasikan
// cara mengubah jumlah thread yang digunakan oleh Go runtime
func TestChangeThreadNumber(t *testing.T) {
	// Inisialisasi WaitGroup untuk sinkronisasi goroutine
	group := sync.WaitGroup{}

	// Membuat 100 goroutine yang berjalan secara concurrent
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			// Simulasi pekerjaan dengan sleep selama 3 detik
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	// Mendapatkan jumlah CPU yang tersedia di sistem
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// Mengubah jumlah maksimum thread menjadi 20
	runtime.GOMAXPROCS(20)
	// Mengambil nilai GOMAXPROCS yang telah diubah
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// Mendapatkan jumlah goroutine yang sedang berjalan
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	// Menunggu semua goroutine selesai
	group.Wait()
}

