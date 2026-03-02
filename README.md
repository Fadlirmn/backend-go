🚀 Belajar Go API - Level 2 (PostgreSQL Integration)
Project ini adalah tahap pembelajaran Backend Go tingkat menengah. Fokus utamanya adalah berpindah dari penyimpanan data in-memory (slice) ke database relasional PostgreSQL menggunakan arsitektur yang bersih (Clean Architecture).
🛠️ Tech Stack
Language: Go (Golang)
Database: PostgreSQL
Driver: pgx/v5 (Standard Library Compatible)
OS Environment: Arch Linux
🏗️ Struktur Proyek
Proyek ini menggunakan pemisahan layer agar kode mudah dimaintain:

backend-api-belajar/
├── config/ # Koneksi database PostgreSQL
├── handler/ # HTTP Request & Response (Controller)
├── model/ # Struct data (User)
├── repository/ # Query langsung ke database (SQL)
├── service/ # Logika bisnis
└── main.go # Entry point & Dependency Injection
