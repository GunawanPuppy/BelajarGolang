# GoRoutine
- lightweight thread yang di manage oleh GO runtime
- memiliki sifat asynchronous jadi tidak saling menunggu dengan Go routine lain
- pembuatan goroutine baru diawali dgn sintaks go diikuti nama method
- definisi : unit eksekusi ringan yg dijalankan bersamaan (asynchronous process) dalam satu program

# Chanel
- penghubung antara go routine satu dengan lainnya. 
- Mekanisme serah terima 

## Deklarasi Channel
- utk buat channel baru contoh: angka := make(chan int).
- Cara mengirim a <- 10
- Cara menerima angka := < - a
## Menutup Channel
- Pengirim punya kemampuan menutup channel utk memberitahu penerima sudah tidak ada data lagi
- Penerima dpt memberikan tambahan kondisi apakah channel uda ditutup atau belum