# FlowCamp Tugas API Ajitirto

Proyek ini adalah RESTful API yang dibangun menggunakan Go (Golang) dan framework Gin. API ini menyediakan endpoint GET, POST, PATCH, DELETE

## Struktur Direktori

```bash

╰─❯ tree
.
├── docker-compose.yml
├── Dockerfile
├── FlowCamp-Tugas-Api-Ajitirto.postman_collection.json
├── gin.log
├── go.mod
├── go.sum
├── issue.md
├── main.go
├── Makefile
├── README.md
├── script.sh
└── src
    ├── config
    │   ├── database.go
    │   └── jwt.go
    ├── controllers
    │   ├── auth_controller.go
    │   └── student_controller.go
    ├── middleware
    │   └── auth_middleware.go
    ├── models
    │   ├── admin.go
    │   ├── class.go
    │   ├── student-class.go
    │   └── student.go
    └── utils
        ├── create-table.go
        ├── drop-table.go
        └── jwt.go

6 directories, 23 files
```


* `docker-compose.yml`: Konfigurasi Docker Compose untuk menjalankan aplikasi dan database MySQL.
* `Dockerfile`: Konfigurasi Docker untuk membangun image aplikasi Go.
* `FlowCamp-Tugas-Api-Ajitirto.postman_collection.json`: Koleksi Postman untuk menguji API.
* `gin.log`: Log dari aplikasi Gin.
* `go.mod` dan `go.sum`: File dependensi Go Modules.
* `issue.md`: Berkas untuk mencatat masalah.
* `main.go`: File entry point aplikasi.
* `Makefile`: Skrip build dan run aplikasi.
* `README.md`: Dokumen ini.
* `script.sh`: Skrip shell untuk beberapa keperluan.
* `src/`: Direktori kode sumber aplikasi.
    * `config/`: Konfigurasi database dan JWT.
    * `controllers/`: Handler untuk endpoint API.
    * `middleware/`: Middleware untuk otentikasi.
    * `models/`: Model data.
    * `utils/`: Fungsi utilitas.

## Prasyarat

* Go (Golang) versi 1.17 atau lebih tinggi.
* Docker dan Docker Compose (jika ingin menjalankan aplikasi menggunakan Docker).
* MySQL (jika ingin menjalankan aplikasi secara lokal tanpa Docker).
* Postman atau alat sejenis untuk menguji API.

## Instalasi dan Penggunaan

### Menjalankan dengan Docker Compose

1.  Pastikan Docker dan Docker Compose sudah terinstal.
2.  Buka terminal di direktori proyek `ajitirto`.
3.  Jalankan `docker compose up -d`.
4.  Jalankan `go run main.go` di root folder
5.  Aplikasi akan berjalan di `http://localhost:8080`.



### Menggunakan Makefile

Makefile menyediakan perintah yang mudah untuk menjalankan aplikasi:

* `make run`: Menjalankan aplikasi secara lokal.
* `make deps`: Mendownload depedensi.
* `make help`: Melihat bantuan perintah 

### Menguji API dengan Postman

1.  Impor koleksi `FlowCamp-Tugas-Api-Ajitirto.postman_collection.json` ke Postman.
2.  Gunakan koleksi tersebut untuk menguji endpoint API.

### Konfigurasi Variabel Lingkungan

Variabel lingkungan berikut digunakan untuk konfigurasi:

* `DB_HOST`: Host database MySQL.
* `DB_USER`: Username database MySQL.
* `DB_PASSWORD`: Password database MySQL.
* `DB_NAME`: Nama database MySQL.

* `JWT_SECRET_KEY`: Secret key untuk JWT.
* `JWT_EXPIRES_IN`: Expired untuk JWt.

### Endpoint API

* `POST api/admin/register` - Registrasi admin
* `POST api/admin/login` - Login admin

* `POST api/student` - Membuat data siswa
* `PATCH api/student/:studentId` - Mengedit data siswa secara parsial berdasarkan `studentId`
* `DELETE api/student/:studentId` - Menghapus siswa berdasarkan `studentId`
* `GET api/student` - Mendapatkan daftar siswa