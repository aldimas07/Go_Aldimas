API merupakan sekumpulan prosedur dan function yang dapat mengakses fitur dari sebuah sistem operasi, aplikasi atau layanan lainnya
Cara kerja APi yaitu dari client membuat request ke server, lalu server merespon sesuai permintaan client
Salah satu kegunaan restful API yaitu mengintegrasikan antara app berbasis frontend dan backend
REST merupakan singkatan dari Representational State Transfer: Sekumpulan aturan untuk membangun sistem terdistribusi
REST menggunakan http protocol untuk berkomunikasi dengan resource
Terdapat http request method: yaitu GET untuk mengambil data, POST untuk insert data, PUT untuk edit data, dan DELETE untuk menghapus data
Untuk request dan response formatnya ada JSON, XML, dan SOAP

JSON (Javascript Object Notation) merupakan format data untuk penukaran dan penyimpanan data, mirip struktur data MAP ada key dan value

HTTP response code untuk identifikasi request yang kita lakukan berhasil atau tidaknya. Jenisnya:
200 = OK - berhasil
201 = Created - berhasil namun ada proses PUT / POST data
400 = Bad Request - terdapat proses gagal pada POST / PUT, biasanya pada validasi
404 = Not Found - Resource yang kita request itu tidak ada
401 = Unauthorized - Resource yang kita akses memerlukan autentikasi terlebih dahulu
405 = Method Not Allowed - Di respond code kita tidak menggunakan method yang tidak sesuai dengan resource
500 = Internal Server Error - Terdapat error pada resource pada sisi server

Untuk pendeklarasiannya diharuskan menggunakan kata benda dan jamak karena naming method get post sudah mewakili dan menjelaskan kegunaannya
Untuk membuat desain endpoint yang baik yaitu dengan urutan:
/users : list user
/users/123 : user secara spesifik
/users/123/orders : mendapatkan data order dari user tersebut
/users/123/orders/0001 : mendapatkan detail ordernya dari user tersebut

OpenAPI: merupakan API yang dapat digunakan secara bebas dan publik untuk belajar dan eksplor
package net/http ini merupakan package yang bisa membuat server API
package encoding/json ini berguna untuk decode JSON ke struktur data yang ada di golang