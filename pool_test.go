// Package belajar_golang_goroutines berisi implementasi dan pengujian goroutines
package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestPool menguji implementasi sync.Pool untuk penggunaan resource pooling
// sync.Pool berguna untuk menyimpan dan menggunakan kembali objek temporary
func TestPool(t *testing.T) {
	// Inisialisasi sync.Pool dengan fungsi New yang akan dipanggil
	// ketika pool kosong dan membutuhkan objek baru
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	// Memasukkan beberapa string ke dalam pool
	pool.Put("Aidil")        // Menambahkan "Aidil" ke pool
	pool.Put("Adam")  // Menambahkan "Adam" ke pool
	pool.Put("Baik Hati")   // Menambahkan "Baik Hati" ke pool

	// Membuat 10 goroutine yang akan mengakses pool secara concurrent
	for i := 0; i < 10; i++ {
		go func() {
			// Mengambil data dari pool
			data := pool.Get()
			// Menampilkan data yang diambil
			fmt.Println(data)
			// Simulasi pemrosesan dengan delay 1 detik
			time.Sleep(1 * time.Second)
			// Mengembalikan data ke pool untuk digunakan kembali
			pool.Put(data)
		}()
	}

	// Menunggu semua goroutine selesai (11 detik untuk memastikan semua selesai)
	time.Sleep(11 * time.Second)
	// Menampilkan pesan setelah semua goroutine selesai
	fmt.Println("Selesai")
}
