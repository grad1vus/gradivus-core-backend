# Gradivus Core Backend (Identity Service)

Service fondasi ekosistem Gradivus: auth, user, dan organization/workspace.
Semua modul lain (Board, HR, Plan, dll) bergantung ke service ini buat login & validasi identitas — jadi user cukup **login sekali** buat akses semua modul (SSO).

## Tanggung Jawab
- Register & login (`/api/auth/register`, `/api/auth/login`)
- Terbitkan JWT token
- Validasi token buat modul lain (`/api/auth/me`)
- Kelola data user & organization/workspace

## Cara Jalanin

```bash
cp .env.example .env
go mod tidy
go run cmd/api/main.go
```

Default jalan di port `8081` (beda port dari modul lain, misal Board di `8080`).

Apply migration:
```bash
psql "$DATABASE_URL" -f migrations/0001_init.sql
```

## Cara Modul Lain Pakai Service Ini

Modul lain (misal `gradivus-board-backend`) tinggal:
1. Terima `Authorization: Bearer <token>` dari frontend
2. Panggil `GET {CORE_URL}/api/auth/me` dengan token itu buat verifikasi & ambil `user_id`, `organization_id`, `role`
3. Kalau valid, lanjut proses request; kalau tidak, return 401

## Status
Skeleton awal — handler `Register`, `Login`, query database masih `// TODO`, belum ada rate limiting untuk login.
