package main

import (
	_driverFactory "Go_Aldimas/22_docker/simpleapp/drivers"
	util "Go_Aldimas/22_docker/simpleapp/utils"

	// "encoding/json"
	"net/http"
	"testing"

	_userUsecase "Go_Aldimas/22_docker/simpleapp/businesses/users"
	_userController "Go_Aldimas/22_docker/simpleapp/controllers/users"
	"Go_Aldimas/22_docker/simpleapp/controllers/users/request"

	_dbDriver "Go_Aldimas/22_docker/simpleapp/drivers/mysql"
	"Go_Aldimas/22_docker/simpleapp/drivers/mysql/users"

	_middlewares "Go_Aldimas/22_docker/simpleapp/app/middlewares"
	_routes "Go_Aldimas/22_docker/simpleapp/app/routes"

	echo "github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func newApp() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT:      util.GetConfig("JWT_SECRET_KEY"),
		ExpireDuration: 1,
	}

	configLogger := _middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMiddleware:    configJWT.Init(),
		UserController:   *userCtrl,
	}

	routesInit.RouteRegister(e)

	return e
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		configDB := _dbDriver.ConfigDB{
			DB_USERNAME: util.GetConfig("DB_USERNAME"),
			DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
			DB_HOST:     util.GetConfig("DB_HOST"),
			DB_PORT:     util.GetConfig("DB_PORT"),
			DB_NAME:     util.GetConfig("DB_TEST_NAME"),
		}

		db := configDB.InitDB()

		_dbDriver.CleanSeeders(db)
	}
}

// func getJWTToken(t *testing.T) string {
// 	configDB := _dbDriver.ConfigDB{
// 		DB_USERNAME: util.GetConfig("DB_USERNAME"),
// 		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
// 		DB_HOST:     util.GetConfig("DB_HOST"),
// 		DB_PORT:     util.GetConfig("DB_PORT"),
// 		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
// 	}

// 	db := configDB.InitDB()

// 	user := _dbDriver.SeedUser(db)

// 	var userRequest *request.User = &request.User{
// 		Email:    user.Email,
// 		Password: user.Password,
// 	}

// 	var resp *http.Response = apitest.New().
// 		Handler(newApp()).
// 		Post("/v1/auth/login").
// 		JSON(userRequest).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End().Response

// 	var response map[string]string = map[string]string{}

// 	json.NewDecoder(resp.Body).Decode(&response)

// 	var token string = response["token"]

// 	var JWT_TOKEN = "Bearer " + token

// 	return JWT_TOKEN
// }

func getUser() users.User {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	db := configDB.InitDB()

	user := _dbDriver.SeedUser(db)

	return user
}

func TestRegister_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/v1/auth/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/v1/auth/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	user := getUser()

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Handler(newApp()).
		Post("/v1/auth/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/v1/auth/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		Handler(newApp()).
		Post("/v1/auth/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}