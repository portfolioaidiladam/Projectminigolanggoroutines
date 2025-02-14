// Package belajar_golang_goroutines berisi implementasi penggunaan goroutine dengan sync.Map
package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// AddToMap adalah fungsi yang digunakan untuk menambahkan data ke sync.Map secara concurrent
// Parameters:
//   - data: pointer ke sync.Map yang akan diisi
//   - value: nilai integer yang akan disimpan sebagai key dan value
//   - group: pointer ke WaitGroup untuk sinkronisasi goroutine
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	// Pastikan group.Done() dipanggil ketika fungsi selesai
	defer group.Done()

	// Tambahkan counter ke WaitGroup
	group.Add(1)
	// Simpan data ke sync.Map dengan key dan value yang sama
	data.Store(value, value)
}

// TestMap adalah fungsi test untuk mendemonstrasikan penggunaan sync.Map dalam concurrent programming
// Test ini menunjukkan bagaimana menyimpan 100 angka ke dalam sync.Map secara bersamaan
// menggunakan goroutine
func TestMap(t *testing.T) {
	// Inisialisasi sync.Map untuk menyimpan data secara thread-safe
	data := &sync.Map{}
	// Inisialisasi WaitGroup untuk menunggu semua goroutine selesai
	group := &sync.WaitGroup{}

	// Loop untuk membuat 100 goroutine
	for i := 0; i < 100; i++ {
		// Jalankan AddToMap sebagai goroutine
		go AddToMap(data, i, group)
	}

	// Tunggu sampai semua goroutine selesai
	group.Wait()

	// Tampilkan semua data yang tersimpan dalam sync.Map
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
