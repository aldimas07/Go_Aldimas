package users

import(
  "Go_Aldimas/21_clean_architecture/clean_architecture/businesses/users"
	"golang.org/x/crypto/bcrypt"
  "gorm.io/gorm"
  "fmt"
)

type userRepository struct {
  conn *gorm.DB
}


func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) GetAll() []users.Domain {
  var rec []User

  ur.conn.Find(&rec)
  userDomain := []users.Domain{}

  for _, user := range rec {
    userDomain = append(userDomain, user.ToDomain())
  }
  return userDomain
}

func (ur *userRepository) Register(userDomain *users.Domain) users.Domain {
  password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
  rec := FromDomain(userDomain)
  rec.Password = string(password)
  result := ur.conn.Create(&rec)

  result.Last(&rec)
  return rec.ToDomain()
}

func (ur *userRepository) GetByEmail(userDomain *users.Domain) users.Domain {
  var user User
  ur.conn.First(&user, "email=?", userDomain.Email)

  if user.ID == 0 {
    fmt.Println("not found")
    return users.Domain{}
  }
  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password)); err != nil{
    fmt.Println("wrong password")
    return users.Domain{}
  }
  return user.ToDomain()
}