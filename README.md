# cake-store

#DB Migration
pastikan bahwa anda telah menginstal golang-migrate
jika belum anda bisa melihat dokumentasinya di [sini](https://github.com/golang-migrate/migrate)

migration database dengan command line:
```bash
migrate -path db/migration -database "mysql://root:@tcp(localhost:3306)/cake" up 1
```