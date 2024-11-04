# Contoh-CRUD-Library

Dengan Bahasa Golang

1. package dan import
  package main: Menandakan ini adalah package utama yang bisa dieksekusi:
  Import packages yang dibutuhkan:
    1. bufio: Untuk membaca input dari user.
    2. fmt Untuk input/output formatting.
    3. os: Untuk akses sistem operasi
    4. strconv: Untuk konversi string ke number.

2. Variable Global:
   Membuat slice string kosong untuk menyimpan daftar buku.
3. Fungsi Main:
   1. Membuat scanner untuk membaca input.
   2. Loop infinite untuk menampilkan menu terus menerus.
   3. Switch case untuk menjalankan fungsi sesuai pilihan user.
4. Fungsi AddBook:
   4. Menerima input judul buku dari user.
   5. Menambahkan buku ke slice books
   6. Menampilkan konfirmasi dan list buku terbaru.
5. Fungsi viewBooks:
   1. Mengecek apakah ada buku dalam list.
   2. Menampilkan semua buku dengan nomor urut.
6. Fungsi deleteBook:
   1. Menampilkan list buku terlebih dahulu.
   2. Meminta input nomor buku yang akan dihapus.
   3. Validasi input user.
   4. Menghapus buku dari slice menggunakan slice operations,
   5. Menampilkan konfirmasi dan list buku terbaru.
   
## RINGKASAN DARI PROGRAM INI

1. Menambah Judul buku baru.
2. Melihat daftar buku.
3. Menghapus buku yang ada.
4. Exit atau keluar dari program.

### Keterangan:

Semua operasi disimpan dalam sebuah memori (slice) dan menghilang jikalau program ditutup.

## TERIMAKASIH SEKIAN, JANGAN LUPA SEMANGAT NGODINGNYA 🔥