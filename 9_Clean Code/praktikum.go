package main

type user struct {
	id       int //penamaan id ini perlu diperjelas lagi sehingga dapat diketahui id ini milik siapa
	username int 
	password int //penamaan password perlu diperjelas lagi sehingga dapat diketahui variabel kepemilikan password ini
}

type userservice struct {
	t []user /* Penamaan t ini kurang jelas apa artinya sehingga fungsinya lebih susah dipahami */
}

func (u  /* Penamaan u ini kurang jelas sehingga fungsinya lebih susah dipahami */ userservice) getallusers() []user {
	return u.t //u.t ini kurang jelas, sehingga harus membaca ulang kode sebelumnya untuk memahami apa u.t ini
} 
/* Ada 7 kekurangan */
func (u /* Penamaan u ini kurang jelas sehingga fungsinya lebih susah dipahami */ userservice) getuserbyid(id int) user {
	for _, r /* Penamaan r ini kurang jelas sehingga fungsinya lebih susah dipahami */ := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}

//===============================================================================

type structmobil struct {
	totalroda int
	kecepatanperjam int
}

func (mobil structmobil) berjalan() {
	mobil.tambahkecepatan(10)
}

func (mobil structmobil) tambahkecepatan(kecepatanbaru int) {
	mobil.kecepatanperjam = mobil.kecepatanperjam + kecepatanbaru
}

func main() {
	mobilcepat := structmobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := structmobil{}
	mobillamban.berjalan()
}
