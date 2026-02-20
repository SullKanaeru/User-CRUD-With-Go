# Akses link ini untuk testing api yang sudah dipublic: https://drive.google.com/file/d/1sQmG_FObWsVq7AFPbCVWWmKRwPovnwrh/view?usp=sharing

# CRUD User API - Role Management System

RESTful API berkinerja tinggi yang dibangun menggunakan **Go (Golang)** dan **Fiber Framework**. API ini dirancang khusus untuk mengelola data pengguna dengan sistem berbasis peran (*Role-Based Access Control*) yang mencakup `admin`, `owner`, dan `customer`. Sangat cocok digunakan sebagai *backend* untuk aplikasi direktori UMKM.

## 🚀 Fitur Utama

* **RESTful CRUD Operations:** Endpoint lengkap untuk Create, Read, Update, dan Delete pengguna.
* **Role-Based Filtering:** Pencarian pengguna secara spesifik berdasarkan *role* menggunakan *Query Parameters*.
* **Enum & Default Values:** Validasi ketat pada level aplikasi dan database. Jika peran tidak dikirimkan, sistem otomatis mengatur pengguna sebagai `customer`.
* **Auto-Migration:** Sinkronisasi struktur tabel otomatis ke database menggunakan GORM.
* **Cloud-Ready:** Terintegrasi mulus dengan PostgreSQL (Neon.tech / Supabase).

## 🛠️ Tech Stack

* **Bahasa Pemrograman:** [Go (Golang)](https://golang.org/)
* **Web Framework:** [Fiber v2](https://gofiber.io/)
* **ORM:** [GORM](https://gorm.io/)
* **Database:** PostgreSQL (via [Neon.tech](https://neon.tech/))
* **Environment Manager:** [godotenv](https://github.com/joho/godotenv)

---

## ⚙️ Persiapan & Instalasi

Pastikan kamu sudah menginstal **Go** di komputermu.

1. **Clone repositori ini atau masuk ke direktori proyek:**
   ```bash
   cd crud-user-api

```

2. **Unduh semua *dependencies* yang dibutuhkan:**
```bash
go mod tidy

```


3. **Atur Environment Variables:**
Buat sebuah file bernama `.env` di direktori utama (sejajar dengan `main.go` atau `go.mod`), lalu tambahkan konfigurasi berikut:
```env
# Ganti dengan URL dari database PostgreSQL (Neon/Supabase) milikmu
DATABASE_URL=postgresql://user:password@host/dbname?sslmode=require&channel_binding=require
PORT=3000

```


4. **Jalankan Aplikasi:**
```bash
go run cmd/api/main.go

```


*Catatan: Saat pertama kali dijalankan, GORM akan otomatis membuat tabel `users` di dalam database jika belum ada.*

---

## 📡 Dokumentasi API (Endpoints)

Base URL lokal: `http://localhost:3000`

| Method | Endpoint | Deskripsi |
| --- | --- | --- |
| **POST** | `/api/users` | Menambahkan pengguna baru. |
| **GET** | `/api/users` | Mengambil daftar semua pengguna. |
| **GET** | `/api/users?role={role}` | Mengambil pengguna berdasarkan peran (`admin`, `owner`, `customer`). |
| **PUT** | `/api/users/:id` | Memperbarui data pengguna berdasarkan ID. |
| **DELETE** | `/api/users/:id` | Menghapus data pengguna berdasarkan ID. |

---

## 💡 Contoh Penggunaan (Payload JSON)

### 1. Create User (POST `/api/users`)

Secara *default*, jika kunci `role` tidak diikutsertakan, sistem akan mendaftarkan pengguna sebagai `customer`.

**Request Body (Admin/Owner):**

```json
{
    "name": "Nama Pengguna",
    "email": "email@example.com",
    "role": "owner" 
}

```

### 2. Update User (PUT `/api/users/:id`)

**Request Body:**

```json
{
    "name": "Nama Baru",
    "email": "emailbaru@example.com",
    "role": "admin"
}

```

---

## 🛡️ Aturan Peran (Roles)

Sistem ini menggunakan validasi *Enum* yang ketat. Nilai `role` yang diizinkan hanya:

* `admin`
* `owner`
* `customer` (Nilai *Default*)

Mengirimkan *role* di luar ketiga nilai tersebut akan menghasilkan status `400 Bad Request`.
