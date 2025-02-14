// Package belajar_golang_goroutines berisi contoh implementasi goroutine di Go
package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// RunAsynchronous adalah fungsi yang dijalankan secara asynchronous
// menggunakan WaitGroup untuk sinkronisasi
// Parameter:
// - group: pointer ke sync.WaitGroup untuk mengelola goroutine
func RunAsynchronous(group *sync.WaitGroup) {
	// Pastikan group.Done() dipanggil ketika fungsi selesai
	defer group.Done()

	// Tambahkan counter goroutine ke WaitGroup
	group.Add(1)

	// Cetak pesan dan tunggu 1 detik untuk simulasi proses
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

// TestWaitGroup adalah fungsi test untuk mendemonstrasikan penggunaan WaitGroup
// dalam mengelola multiple goroutines
func TestWaitGroup(t *testing.T) {
	// Inisialisasi WaitGroup
	group := &sync.WaitGroup{}

	// Jalankan 100 goroutine secara bersamaan
	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	// Tunggu semua goroutine selesai
	group.Wait()
	fmt.Println("Selesai")
}

