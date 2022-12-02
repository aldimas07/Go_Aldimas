#Summary

### Arsitektur yg bisa terpisahkan secara layer dan membuat modular yg scalable
### di Clean Architecture, kita bisa merubah UI tanpa merubah business logicnya

### Benefitnya:
#### Standar Strukturnya sudah ada dan lebih cepat developmentnya
#### Unit testnya bisa berjalan dengan baik

### CA Layer:
#### Entity layer: berisi business object
#### Usecase - domain layer: berisi business logic
#### Controller - presentation layer: berisi framework. menyajikan data ke layar / user interaction
#### Driver - data layer: manajemen data, menerima data dari network / third party, retrieve data

### Approachnya
#### Approach 1: CA layers dalam 1 modul. app module ada presentation, domain, data
#### Approach 2: 1 CA layer berisi 1 modul
#### Approach 3: CA layers berada di dalam feature modul
#### Approach 4: CA layers berisi feature model

### DDD Domain Driven Design: teknik desain software