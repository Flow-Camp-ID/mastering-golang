# FlowcampId LMS API

## Penulis

[Nama Anda]

## Deskripsi Singkat

Aplikasi LMS FlowcampId adalah API RESTful yang dirancang untuk mengelola data siswa, kelas, dan otentikasi admin.

## Implementasi Endpoint

### 1. Auth

* `POST api/admin/register` - Registrasi admin
* `POST api/admin/login` - Login admin

### 2. CRUD Siswa

* `POST api/student` - Membuat data siswa
* `PATCH api/student/:studentId` - Mengedit data siswa secara parsial berdasarkan `studentId`
* `DELETE api/student/:studentId` - Menghapus siswa berdasarkan `studentId`
* `GET api/student` - Mendapatkan daftar siswa

### 3. CRUD Kelas (Opsional)

* `POST api/class` - Membuat data kelas
* `GET api/class` - Mendapatkan daftar kelas
* `GET api/class/:classId` - Mendapatkan detail kelas berdasarkan `classId`
* `POST api/class/:classId/students` - Menambahkan siswa ke kelas berdasarkan `classId`
* `DELETE api/class/:classId/students` - Menghapus siswa dari kelas berdasarkan `classId`
* `PUT api/class/:classId` - Mengedit kelas secara keseluruhan berdasarkan `classId`
* `DELETE api/class/:classId` - Menghapus kelas berdasarkan `classId`

## Teknologi & Implementasi

* [Daftar library yang digunakan, contoh: Gin, GORM, dll.]
* [Jika menggunakan goroutine atau channel, sebutkan di sini.]
* [Arsitektur model yang digunakan, contoh: MVC, dll.]

## Perjalanan Pembuatan Tugas

* Tantangan: Implementasi middleware autentikasi menggunakan JWT.
* Solusi: Menggunakan library `github.com/golang-jwt/jwt/v4` untuk menghasilkan dan memvalidasi token.
* Pembelajaran: Menangani middleware secara efisien dapat meningkatkan keamanan API.

## Checklist PR

* [x] Kode sudah dites secara lokal
* [x] Tidak ada error atau bug yang menghambat fungsi utama
* [x] Sudah membuat README cara menjalankan RESTful API yang dibuat
* [x] PR ini sudah siap untuk direview

## Tanggung Jawab

Saya menyatakan bahwa:

* Kode dalam PR ini adalah hasil kerja saya sendiri tanpa menyalin dari sumber lain.
* Jika ada referensi yang digunakan, saya telah mencantumkannya dengan jelas.
* Saya menerima konsekuensi jika ada pelanggaran terhadap aturan ini.