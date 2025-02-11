package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// TestCreateChannel mendemonstrasikan penggunaan dasar channel
// Channel digunakan untuk komunikasi dan sinkronisasi antar goroutine
func TestCreateChannel(t *testing.T) {
	// Membuat channel string baru yang tidak di-buffer
	channel := make(chan string)
	// Memastikan channel ditutup setelah fungsi selesai
	defer close(channel)

	// Membuat goroutine anonymous untuk mengirim data
	go func() {
		time.Sleep(2 * time.Second)  // Simulasi proses yang memakan waktu
		channel <- "Aidil Adam Baik Hati"  // Mengirim data ke channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	// Menerima data dari channel (operasi blocking)
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)  // Memberikan waktu untuk goroutine selesai
}

// GiveMeResponse adalah fungsi helper yang mengirim data ke channel setelah delay
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Aidil Adam Baik Hati"
}

// TestChannelAsParameter mendemonstrasikan cara menggunakan channel sebagai parameter fungsi.
// Fungsi ini menunjukkan pola umum dalam komunikasi antar goroutine menggunakan channel.
func TestChannelAsParameter(t *testing.T) {
	// Membuat channel string baru yang tidak di-buffer
	// Channel ini akan digunakan untuk komunikasi antar goroutine
	channel := make(chan string)

	// Memastikan channel ditutup setelah fungsi selesai
	// untuk mencegah memory leak dan deadlock
	defer close(channel)

	// Menjalankan fungsi GiveMeResponse dalam goroutine terpisah
	// dan mengirimkan channel sebagai parameter
	go GiveMeResponse(channel)

	// Menerima data dari channel (operasi blocking)
	// Program akan menunggu sampai ada data yang dikirim
	data := <-channel

	// Menampilkan data yang diterima dari channel
	fmt.Println(data)

	// Memberikan waktu untuk memastikan goroutine selesai dieksekusi
	// Catatan: Dalam praktik nyata, lebih baik menggunakan WaitGroup
	// daripada time.Sleep untuk sinkronisasi
	time.Sleep(5 * time.Second)
}

// OnlyIn mendemonstrasikan channel yang hanya bisa menerima data (write-only)
// chan<- menandakan channel hanya bisa digunakan untuk mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Aidil Adam Baik Hati"
}

// OnlyOut mendemonstrasikan channel yang hanya bisa membaca data (read-only)
// <-chan menandakan channel hanya bisa digunakan untuk menerima data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// TestInOutChannel mendemonstrasikan penggunaan channel satu arah (unidirectional channel)
// untuk memastikan keamanan tipe dan mencegah penggunaan channel yang tidak diinginkan
func TestInOutChannel(t *testing.T) {
	// Membuat channel string tanpa buffer
	channel := make(chan string)
	// Menutup channel setelah fungsi selesai untuk mencegah memory leak
	defer close(channel)

	// Menjalankan fungsi OnlyIn sebagai goroutine terpisah
	// OnlyIn hanya dapat menulis ke channel (chan<-)
	go OnlyIn(channel)
	// Menjalankan fungsi OnlyOut sebagai goroutine terpisah
	// OnlyOut hanya dapat membaca dari channel (<-chan)
	go OnlyOut(channel)

	// Memberikan waktu untuk goroutine menyelesaikan eksekusinya
	// Note: Dalam produksi, lebih baik menggunakan sync.WaitGroup
	time.Sleep(5 * time.Second)
}

// TestBufferedChannel mendemonstrasikan penggunaan buffered channel
// yang memungkinkan pengiriman data asynchronous sampai batas buffer tertentu
func TestBufferedChannel(t *testing.T) {
	// Membuat buffered channel dengan kapasitas 3
	// Channel dapat menampung 3 pesan sebelum blocking
	channel := make(chan string, 3)
	// Menutup channel setelah fungsi selesai
	defer close(channel)

	// Goroutine pertama untuk mengirim data
	go func() {
		// Mengirim 3 string ke channel
		// Operasi ini tidak akan blocking karena masih dalam kapasitas buffer
		channel <- "Aidil"
		channel <- "Adam"
		channel <- "Baik"
	}()

	// Goroutine kedua untuk membaca data
	go func() {
		// Membaca dan mencetak setiap nilai dari channel
		// Operasi ini akan mengosongkan buffer satu per satu
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	// Memberikan waktu untuk goroutine menyelesaikan eksekusinya
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// TestRangeChannel mendemonstrasikan cara aman untuk membaca
// semua data dari channel menggunakan range loop
func TestRangeChannel(t *testing.T) {
	// Membuat unbuffered channel
	channel := make(chan string)

	// Goroutine untuk mengirim data
	go func() {
		// Mengirim 10 pesan ke channel
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		// Penting: Menutup channel setelah selesai mengirim
		// untuk memberitahu range loop bahwa tidak ada data lagi
		close(channel)
	}()

	// Range loop akan terus berjalan sampai channel ditutup
	// Ini adalah cara yang aman untuk membaca semua data dari channel
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// TestSelectChannel menguji penggunaan select untuk menangani multiple channel secara bersamaan.
// Select memungkinkan kita untuk menunggu dan menerima data dari beberapa channel sekaligus.
func TestSelectChannel(t *testing.T) {
	// Membuat dua channel string tanpa buffer
	channel1 := make(chan string)
	channel2 := make(chan string)
	
	// Memastikan kedua channel ditutup setelah fungsi selesai
	// untuk mencegah memory leak
	defer close(channel1)
	defer close(channel2)

	// Menjalankan dua goroutine yang akan mengirim data ke masing-masing channel
	// setelah delay 2 detik
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// Counter untuk melacak jumlah data yang telah diterima
	counter := 0
	
	// Loop tak terbatas yang akan berhenti setelah menerima data dari kedua channel
	for {
		// Block select akan menunggu sampai salah satu case siap dieksekusi
		select {
		case data := <-channel1:
			// Case ini dieksekusi ketika data tersedia di channel1
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			// Case ini dieksekusi ketika data tersedia di channel2
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		// Keluar dari loop setelah menerima data dari kedua channel
		if counter == 2 {
			break
		}
	}
}

// TestDefaultSelectChannel menguji penggunaan select dengan case default.
// Fungsi ini menunjukkan bagaimana menangani situasi ketika semua channel blocking.
func TestDefaultSelectChannel(t *testing.T) {
	// Membuat dua channel string tanpa buffer
	channel1 := make(chan string)
	channel2 := make(chan string)
	
	// Memastikan kedua channel ditutup setelah fungsi selesai
	defer close(channel1)
	defer close(channel2)

	// Menjalankan dua goroutine yang akan mengirim data ke masing-masing channel
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// Counter untuk melacak jumlah data yang telah diterima
	counter := 0
	
	// Loop tak terbatas yang akan berhenti setelah menerima data dari kedua channel
	for {
		// Block select dengan default case
		select {
		case data := <-channel1:
			// Case ini dieksekusi ketika data tersedia di channel1
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			// Case ini dieksekusi ketika data tersedia di channel2
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			// Case ini dieksekusi ketika kedua channel masih blocking (belum ada data)
			// Mencegah CPU bound karena tight loop
			fmt.Println("Menunggu Data")
		}

		// Keluar dari loop setelah menerima data dari kedua channel
		if counter == 2 {
			break
		}
	}
}
