package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Mutex untuk mengontrol akses ke resource yang dibagi
var locker = sync.Mutex{}

// Cond untuk mengimplementasikan mekanisme sinkronisasi menggunakan kondisi
var cond = sync.NewCond(&locker)

// WaitGroup untuk menunggu semua goroutine selesai
var group = sync.WaitGroup{}

// WaitCondition adalah fungsi yang akan dijalankan oleh goroutine
// Parameter value digunakan untuk mengidentifikasi goroutine
func WaitCondition(value int) {
	// Menandakan goroutine telah selesai saat fungsi berakhir
	defer group.Done()
	
	// Menambah counter WaitGroup
	group.Add(1)

	// Mengunci mutex sebelum mengakses conditional variable
	cond.L.Lock()
	
	// Menunggu sinyal dari goroutine lain
	cond.Wait()
	
	// Mencetak pesan setelah menerima sinyal
	fmt.Println("Done", value)
	
	// Membuka kunci mutex setelah selesai
	cond.L.Unlock()
}

// TestCond adalah fungsi test untuk mendemonstrasikan penggunaan sync.Cond
func TestCond(t *testing.T) {
	// Membuat 10 goroutine yang akan menunggu kondisi
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

//	Goroutine untuk mengirim sinyal satu per satu
	go func() {
		for i := 0; i < 10; i++ {
			// Menunggu 1 detik sebelum mengirim sinyal berikutnya
			time.Sleep(1 * time.Second)
			// Membangunkan satu goroutine yang menunggu
			cond.Signal()
		}
	}()

	//Alternatif menggunakan Broadcast (dikomentari)
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	// Membangunkan semua goroutine yang menunggu
	// 	cond.Broadcast()
	// }()

	// Menunggu semua goroutine selesai
	group.Wait()
}
