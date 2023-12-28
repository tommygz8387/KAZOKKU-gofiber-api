# Aplikasi Go Fiber

Aplikasi ini adalah contoh sederhana dari aplikasi web menggunakan framework Go Fiber untuk membuat REST API. Aplikasi ini memiliki fitur registrasi pengguna dan beberapa endpoint untuk mengelola data pengguna.

## Fitur
- Registrasi pengguna
- Pengambilan daftar pengguna
- Pengambilan pengguna berdasarkan ID

## Instalasi

### Persyaratan
- Go (minimal versi 1.21)
- MySQL (atau database lain yang didukung oleh GORM)

### Langkah-langkah instalasi

1. **Clone Repositori**
    ```bash
    git clone https://github.com/tommygz8387/KAZOKKU-gofiber-api.git
    cd KAZOKKU-gofiber-api
    ```

2. **Install Dependensi**
    ```bash
    go mod tidy
    ```

3. **Konfigurasi Database**
    - Buat database di MySQL (atau database lainnya)
    - Salin file `.env.example` menjadi `.env` dan sesuaikan konfigurasi database di dalamnya.

4. **Jalankan Aplikasi**
    ```bash
    go run main.go
    ```

5. **Akses Aplikasi**
    Aplikasi akan berjalan di `http://localhost:3000`. Buka browser dan akses URL tersebut.

## Penggunaan

1. **Registrasi Pengguna**
    - Buat pengguna baru dengan mengirimkan POST request ke `http://localhost:3000/user/register` dengan body JSON yang berisi informasi pengguna.

2. **Ambil Daftar Pengguna**
    - Ambil daftar pengguna dengan mengirimkan GET request ke `http://localhost:3000/user/list`.

3. **Ambil Pengguna Berdasarkan ID**
    - Ambil pengguna berdasarkan ID dengan mengirimkan GET request ke `http://localhost:3000/user/:id`, mengganti `:id` dengan ID pengguna yang diinginkan.

## NB

Dokumentasi Hitpoin API bisa dilihat di file api.rest dan bisa dijalankan dengan menggunakan ekstensi rest client.