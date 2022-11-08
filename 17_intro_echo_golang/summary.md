# Summary

### Di echo konsepnya adalah: Application - Route - Controller - Service - Repository - Database

### Contoh: GET /notes

### Application akan mengarahkan req dari user (Get ke /notes), lalu route akan memetakan (GET /notes) ini akan ditangani oleh controller apa, lalu di controller yang menangani akan di eksekusi. Controller akan menggunakan komponen Service untuk berinteraksi dengan repository. Lalu repository yang akan melakukan business logic / query ke Database. Lalu dari Database akan dikembalikan secara descending dalam bentuk http response

### AIR ini merupakan sebuah tools untuk mengaktifkan hot reload di app yg dikembangkan dengan golang, semisal ada perubahan di source code waktu menjalankan aplikasi, maka app akan otomatis merestart
