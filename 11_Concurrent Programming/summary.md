Summary Concurrent Programming

Sequential Program: menyelesaikan satu task dahulu setelah itu lanjut ke task berikutnya
Parallel Program: dapat menyelesaikan beberapa task dalam waktu yang bersamaan tergantung pada CPUnya
Concurrent Program: menjalankan task secara bersamaan tidak membutuhkan multiple core CPU. Cara kerjanya yaitu membagi-bagi beberapa task tersebut


Goroutines & feature
- func main adalah concurrent utama
- Channel digunakan untuk komunikasi antar beberapa task untuk mengatur task yang akan dieksekusi
- Select: mengontrol jalan jalannya task didalam Channel

goroutines: sebuah function yang dijalankan secara concurrent