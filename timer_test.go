// Package belajar_golang_goroutines berisi implementasi testing untuk fungsi-fungsi timer di Go
package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestTimer menguji penggunaan dasar time.Timer
// Timer digunakan untuk menunda eksekusi kode selama durasi tertentu
func TestTimer(t *testing.T) {
	// Membuat timer baru dengan durasi 5 detik
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now()) // Mencetak waktu saat ini

	// Menunggu hingga timer selesai dan menerima waktu dari channel timer
	time := <-timer.C
	fmt.Println(time) // Mencetak waktu setelah timer selesai
}

// TestAfter menguji penggunaan time.After
// After adalah cara singkat untuk membuat timer sekali pakai
func TestAfter(t *testing.T) {
	// Membuat channel yang akan menerima waktu setelah 5 detik
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now()) // Mencetak waktu saat ini

	// Menunggu hingga 5 detik berlalu dan menerima waktu dari channel
	time := <-channel
	fmt.Println(time) // Mencetak waktu setelah delay
}

// TestAfterFunc menguji penggunaan time.AfterFunc
// AfterFunc mengeksekusi fungsi yang diberikan setelah durasi tertentu
func TestAfterFunc(t *testing.T) {
	// Inisialisasi WaitGroup untuk sinkronisasi
	group := sync.WaitGroup{}
	group.Add(1) // Menambah counter WaitGroup

	// Menjadwalkan fungsi untuk dijalankan setelah 5 detik
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now()) // Mencetak waktu saat fungsi dijalankan
		group.Done()            // Menandai bahwa goroutine telah selesai
	})
	fmt.Println(time.Now()) // Mencetak waktu saat ini

	// Menunggu hingga semua goroutine selesai
	group.Wait()
}


