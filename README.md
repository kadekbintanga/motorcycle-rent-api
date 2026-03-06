# 🚀 Motorcycle Rent API Backend

Proyek ini dibangun menggunakan bahasa pemrograman **Go (Golang)** dengan framework **Gin** struktur folder modular untuk memastikan kode mudah dipelihara, diuji, dan dikembangkan.

---

## 📂 Struktur Proyek

Berikut adalah penjelasan mengenai struktur folder yang digunakan dalam sistem ini:

| Folder | Tanggung Jawab |
| :--- | :--- |
| **`appconfig`** | Inisialisasi fungsionalitas utama (koneksi database, konfigurasi logger, dll). |
| **`constant`** | Menyimpan variabel statis/tetap yang digunakan di seluruh sistem (e.g., Error Messages, Status Codes). |
| **`global`** | Menampung variabel konfigurasi dari environment (`.env`) yang diakses secara global. |
| **`handler`** | Layer Controller yang menangani HTTP request dan response (Input/Output). |
| **`helper`** | Kumpulan fungsi utilitas mandiri (e.g., hashing password, formatting date, string manipulation). |
| **`middleware`** | Fungsi pengaman (JWT), inisialisasi logging per-request, dan generator Request ID. |
| **`model`** | Definisi struktur tabel database (GORM structs) yang merepresentasikan entitas data. |
| **`repository`** | Layer akses data yang berinteraksi langsung dengan database (Query SQL/GORM). |
| **`resource`** | Menampung struct khusus untuk Request (Input binding) dan Formatting Response. |
| **`router`** | Definisi jalur (endpoint) API dan pengelompokan rute (Route Grouping). |
| **`seeder`** | Fungsi untuk keperluan migrasi data awal atau pengisian data dummy ke database. |
| **`service`** | Pusat logika bisnis (Business Logic) yang menghubungkan Handler dengan Repository. |
| **`main.go`** | Titik masuk utama (Entry Point) aplikasi. |

---

## 🛠️ Persiapan & Instalasi

### 1. Prasyarat
* **Go** (versi 1.25++)
* **Framework** (Gin)
* **Database Engine** (PostgreSQL)
* **Git**

### 2. Cara Menjalankan Project
1. Clone repositori ini:
   ```bash
   git clone https://github.com/kadekbintanga/motorcycle-rent-api
   ```
2. Masuk ke direktori:
    ```
    cd motorcycle-rent-api
    ```
3. Instal dependensi:
    ```
    go mod tidy
    ```
4. Salin file environment:
    ```
    cp .env.example .env
    ```
5. Setup config env, seperti credential database dan lain - lainnya jika perlu diubah
6. Jalankan aplikasi:
    ```
    go run main.go
    ```
7. API sudah dapat di akses melalui postman atau tools lainya

### This Project created by : Kadek Bintang Anjasmara