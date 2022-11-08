package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

type UserEdit struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var userId string = c.Param("id")
	id, _ := strconv.Atoi(userId)

	var foundUser User

	for _, user := range users {
		if user.Id == id {
			foundUser = user
		}

	}

	return c.JSON(http.StatusOK, map[string]any{
		"messages": "Success get user",
		"user":     foundUser,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)

	var newUser []User /*= []User{} */ //slice kosong

	for _, valuser := range users { // valuser merepresentasikan isi dari slice users
		if valuser.Id != id { // filtering id yang bukan targetnya
			newUser = append(newUser, valuser) // jika data tidak sama dengan target id dimasukkan ke slice user yang kosong tadi
		}
	}

	// assign semua data user yang bukan target ke slice users
	users = newUser

	// for valuser

	return c.JSON(http.StatusOK, map[string]any{
		"message": "User successfully deleted",
		"user ":   newUser,
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here

	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)
	userEdit := UserEdit{}
	//agar input req bisa di baca perlu di binding
	c.Bind(&userEdit) // parsing di req body ke updatedUser

	updatedUser := User{}

	for index, valuser := range users {
		if valuser.Id == id { //jika ada user yang sesuai dengan req parameter
			// updatedUser = append(updatedUser, valuser)
			valuser.Email = userEdit.Email
			valuser.Name = userEdit.Name
			valuser.Password = userEdit.Password
			users[index] = valuser
			updatedUser = users[index]
		}
	}
	return c.JSON(http.StatusOK, map[string]any{
		"messages": "User updated",
		"user":     updatedUser,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user) //request parsing

	if len(users) == 0 { //penambahan data dengan storage lokal
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user) //menambah item ke dalam slice, users ditambahkan ke slice user. yg ditambahkan yaitu user
	return c.JSON(http.StatusOK, map[string]any{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)

	e.POST("/users", CreateUserController)
	e.GET("/users/:id",GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id",DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
