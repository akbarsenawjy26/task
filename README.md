# Task Management API

API ini bertujuan untuk mengelola task seperti menambahkan, mengupdate, menghapus, dan mengambil task berdasarkan status.

## Fitur

- **Create Task**: Menambahkan task baru.
- **Update Task**: Mengupdate task berdasarkan ID.
- **Delete Task**: Menghapus task berdasarkan ID.
- **Get All Tasks**: Mengambil semua task.
- **Get Task by Status**: Mengambil task berdasarkan status.

## Arsitektur

Proyek ini menggunakan pendekatan Service-Oriented Architecture (SOA) dengan struktur folder sebagai berikut:

```bash
├── cmd/
│   └── api
│       └── routes/
├── common/
│   ├── helper/
├── src/
   ├── repository/
   └── task/
       ├── application/
       ├── controller/
       └── service/
```
## Prasyarat
- Go: v1.20 atau lebih baru
- PostgreSQL
- SQLC
- Framework : Echo

### Cek Spesifikasi API untuk detail
#### @Akbarsenawjy-24