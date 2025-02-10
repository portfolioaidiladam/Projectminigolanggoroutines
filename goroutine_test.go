// Package belajar_golang_goroutines berisi contoh implementasi goroutine di Go
package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// RunHelloWorld adalah fungsi sederhana yang mencetak "Hello World"
// Fungsi ini digunakan sebagai contoh dasar penggunaan goroutine
func RunHelloWorld() {
	fmt.Println("Hello World")
}

// TestCreateGoroutine menguji pembuatan goroutine sederhana
// Test ini mendemonstrasikan sifat asynchronous dari goroutine
// dengan menjalankan RunHelloWorld() secara concurrent
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // Menjalankan fungsi dalam goroutine terpisah
	fmt.Println("Ups")

	// Menunggu 1 detik untuk memastikan goroutine selesai dieksekusi
	// Catatan: Dalam praktik nyata, sebaiknya gunakan WaitGroup untuk sinkronisasi
	time.Sleep(1 * time.Second)
}

// DisplayNumber mencetak nomor yang diberikan dengan format "Display [nomor]"
// Parameter:
//   - number: nilai integer yang akan dicetak
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

// TestManyGoroutine menguji pembuatan banyak goroutine secara bersamaan
// Test ini mendemonstrasikan kemampuan Go untuk menangani ribuan goroutine
func TestManyGoroutine(t *testing.T) {
	// Membuat 100000 goroutine secara bersamaan
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	// Menunggu 5 detik untuk memastikan semua goroutine selesai
	// Catatan: Dalam praktik nyata, sebaiknya gunakan WaitGroup untuk sinkronisasi
	time.Sleep(5 * time.Second)
}
