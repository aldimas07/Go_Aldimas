# Summary

## 1. Directory
### pwd: untuk menampilkan alamat direktori saat ini berada
### ls: menampilkan file dan folder apa saja dalam direktori saat ini user berada
### ls -l: menampilkan file dan folder apa saja dalam direktori saat ini dengan detail (permission, jumlah file, user, group, ukuran file, tanggal, nama file / direktori)
### mkdir: membuat direktori baru
### rmdir: menghapus sebuah direktori

## 2. Permission
### d|rwx|rwx|rwx
### d = merupakan tipe direktori, sedangkan tanda (-) merupakan file
### rwx = read write execute oleh tipe user owner
### rwx = read write execute oleh tipe user group dari owner file
### rwx = read write execute oleh semua tipe user
### rwx jika diterjemahkan ke biner menjadi 111 kalau diubah ke desimal menjadi 7
## 3. Program yang berfungsi menjembatani antara user dengan kernel dinamakan shell