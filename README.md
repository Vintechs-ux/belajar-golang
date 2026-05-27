# belajar-golang

Personal learning repository dokumentasi perjalanan gue belajar Go dari nol setelah migrasi dari Node.js. Ini bukan tutorial, bukan course notes. Ini lebih ke catatan teknikal yang gue tulis buat diri sendiri, tapi cukup structured buat siapapun yang mau baca.

Fokusnya disini adalah pemahaman fundamental kenapa sesuatu bekerja seperti itu.

---

## Materi yang Sudah Dicakup

### Go Fundamentals
Dasar-dasar bahasa tipe data, control flow, function, pointer, struct, interface, goroutine, channel. Fokus di memahami type system dan memory model Go yang berbeda dari JavaScript.

### Standard Library
Eksplorasi paket-paket bawaan Go tanpa dependency eksternal:
- `errors` — error handling idiomatis, wrapping, sentinel errors
- `reflect` — reflection system dan kapan (tidak) menggunakannya
- `regexp` — regex engine dan kompilasi pattern
- `encoding` — JSON, base64, dan format lainnya
- `bufio` / `bytes` — buffer I/O dan manajemen memori
- `os` / `io` — file manipulation dan permission system

### Go Modules
Arsitektur dependency management Go — `go.mod`, `go.sum`, MVS (Minimum Version Selection), local module linking via `replace` directive, dan kenapa Go memilih decentralized model dibanding centralized registry.

### Unit Testing
Testing idiomatis di Go menggunakan package `testing` bawaan  table-driven tests, subtests, benchmark, dan test coverage.
Memahami Test Driven Programming dalam pengembangan projek untuk mempermudah scaling  

---

## Upcoming

- Concurrency patterns (goroutine, channel, sync primitives)
- HTTP server dengan `net/http`
- Database interaction
- CLI tooling
- dan lainnya

---

## Catatan

Gue migrasi dari Node.js, jadi beberapa catatan di sini mungkin ada perbandingan atau konteks dari ekosistem JavaScript. Bukan berarti salah satu lebih baik cuma biar lebih mudah nangkep perbedaan desain filosofinya.

Environment: Arch Linux / CachyOS · Neovim · Terminal-only workflow.
