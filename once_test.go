// Package belajar_golang_goroutines berisi implementasi testing untuk sync.Once
package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// counter adalah variabel global yang akan diincrement
var counter = 0

// OnlyOnce adalah fungsi yang akan dieksekusi hanya sekali menggunakan sync.Once
func OnlyOnce() {
	counter++
}

// TestOnce adalah fungsi testing untuk memastikan bahwa sync.Once berfungsi dengan benar
// dengan menjalankan fungsi OnlyOnce dalam multiple goroutine
func TestOnce(t *testing.T) {
	// Inisialisasi sync.Once untuk memastikan fungsi hanya dijalankan sekali
	once := sync.Once{}
	// Inisialisasi WaitGroup untuk menunggu semua goroutine selesai
	group := sync.WaitGroup{}

	// Loop untuk membuat 100 goroutine
	for i := 0; i < 100; i++ {
		go func() {
			// Menambah counter WaitGroup sebelum eksekusi
			group.Add(1)
			// Menggunakan sync.Once untuk memastikan OnlyOnce hanya dijalankan sekali
			once.Do(OnlyOnce)
			// Menandakan goroutine telah selesai
			group.Done()
		}()
	}

	// Menunggu semua goroutine selesai
	group.Wait()
	// Mencetak nilai akhir counter
	fmt.Println("Counter", counter)
}

