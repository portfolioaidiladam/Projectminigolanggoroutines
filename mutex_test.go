// Package belajar_golang_goroutines berisi implementasi dan pengujian untuk konsep mutex dan goroutine
package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestMutex mendemonstrasikan penggunaan dasar mutex untuk mengamankan akses concurrent ke variable
func TestMutex(t *testing.T) {
	x := 0                  // Variabel yang akan diakses secara concurrent
	var mutex sync.Mutex   // Mutex untuk mengamankan akses ke variabel x

	// Membuat 1000 goroutine yang masing-masing akan menambah nilai x sebanyak 100 kali
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()   // Mengunci akses ke critical section
				x = x + 1      // Critical section
				mutex.Unlock() // Membuka kunci setelah selesai
			}
		}()
	}

	time.Sleep(5 * time.Second) // Menunggu semua goroutine selesai
	fmt.Println("Counter = ", x)
}

// BankAccount merepresentasikan rekening bank dengan RWMutex untuk mengamankan akses concurrent
type BankAccount struct {
	RWMutex sync.RWMutex // RWMutex untuk membedakan operasi read dan write
	Balance int          // Saldo rekening
}

// AddBalance menambahkan sejumlah amount ke saldo rekening dengan menggunakan write lock
func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()   // Mengunci untuk operasi write
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock() // Membuka kunci setelah selesai
}

// GetBalance mengambil nilai saldo rekening dengan menggunakan read lock
func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()  // Mengunci untuk operasi read
	balance := account.Balance
	account.RWMutex.RUnlock() // Membuka kunci read
	return balance
}

// TestRWMutex menguji penggunaan RWMutex dalam operasi concurrent read/write pada rekening bank
// Test ini mendemonstrasikan bagaimana multiple goroutine dapat mengakses dan memodifikasi saldo
// secara aman menggunakan RWMutex
func TestRWMutex(t *testing.T) {
	// Inisialisasi rekening bank baru dengan saldo awal 0
	account := BankAccount{}

	// Membuat 100 goroutine yang akan melakukan operasi secara concurrent
	// Setiap goroutine akan menambah saldo dan membaca saldo sebanyak 100 kali
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)        // Operasi write: menambah saldo
				fmt.Println(account.GetBalance()) // Operasi read: membaca saldo
			}
		}()
	}

	// Memberikan waktu agar semua goroutine selesai dieksekusi
	time.Sleep(5 * time.Second)
	// Menampilkan saldo akhir setelah semua operasi selesai
	fmt.Println("Total Balance", account.GetBalance())
}

// UserBalance merepresentasikan entitas pengguna dengan saldo
// Struct ini menggunakan embedded mutex untuk thread-safety
type UserBalance struct {
	sync.Mutex           // Mutex yang di-embed untuk sinkronisasi akses ke data
	Name    string      // Nama pemilik rekening
	Balance int         // Jumlah saldo yang dimiliki
}

// Lock mengimplementasikan method Lock dari interface sync.Locker
// Method ini harus dipanggil sebelum mengakses atau memodifikasi data UserBalance
func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

// Unlock mengimplementasikan method Unlock dari interface sync.Locker
// Method ini harus dipanggil setelah selesai mengakses atau memodifikasi data UserBalance
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

// Change memodifikasi saldo pengguna
// Method ini mengasumsikan pemanggil sudah melakukan Lock() sebelumnya
// Parameter amount: jumlah perubahan saldo (positif untuk penambahan, negatif untuk pengurangan)
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

// Transfer melakukan pemindahan dana antar rekening
// PERINGATAN: Implementasi ini rentan terhadap deadlock karena penguncian tidak berurutan
// Parameters:
// - user1: pengirim dana
// - user2: penerima dana
// - amount: jumlah yang ditransfer
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	// Mengunci akses ke rekening pengirim
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount) // Mengurangi saldo pengirim

	time.Sleep(1 * time.Second) // Simulasi proses yang memakan waktu

	// Mengunci akses ke rekening penerima
	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount) // Menambah saldo penerima

	time.Sleep(1 * time.Second)

	// Membuka kunci kedua rekening
	user1.Unlock()
	user2.Unlock()
}

// TestDeadlock mendemonstrasikan potensi deadlock dalam transfer concurrent
// Test ini menunjukkan bagaimana dua transfer yang berjalan bersamaan dapat
// menyebabkan deadlock karena pola penguncian yang tidak konsisten
func TestDeadlock(t *testing.T) {
	// Inisialisasi dua rekening dengan saldo awal
	user1 := UserBalance{
		Name:    "Aidil",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}

	// Menjalankan dua transfer secara concurrent dengan arah berlawanan
	// Transfer pertama: Eko -> Budi (100000)
	// Transfer kedua: Budi -> Eko (200000)
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	// Menunggu proses transfer selesai atau terjadi deadlock
	time.Sleep(10 * time.Second)

	// Menampilkan saldo akhir kedua rekening
	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}
