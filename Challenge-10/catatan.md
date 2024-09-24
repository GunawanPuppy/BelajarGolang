# Defer
- defer function dieksekusi terakhir
- function yg dijalankan setelah eksekusi sebuah function
- defer ini seperti finally pada try-catch-finally

# Panic
- function utk menghentikan program
- dipanggil ketika terjadi error saat program berjalan
- saat panic function dipanggil, program akan terhenti
- defer function tetap tereksekusi
- panic ini seperti throw new Error {name : "asd"}

# Recover 
- function utk menangkap data panic
- panic terhenti , program ttp berjalan

## Panic-Recover 
- mirip seperti throw dan try-catch di javascript
- ketika throw new Error {name: "asd"}, di golang kita bisa buat dalam panic("asd") message:= recover ()  

## Double Defer 
- Saat ada Panic , semua kode selanjutnya dihentikan, jadi pastikan defer yang gak ada panic dibiarin jalan duluan dan juga pemanggilan function lainnya

## package Flag
- Pada program ini package flag mengambil nilai panjang dan lebar dari input command line 
- Contoh Command : go run tugas10.go -panjang=10 -lebar=5