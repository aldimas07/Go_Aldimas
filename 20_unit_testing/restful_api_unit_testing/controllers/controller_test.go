package controllers

import (
	database "Go_Aldimas/20_unit_testing/restful_api_unit_testing/databases"
	"Go_Aldimas/20_unit_testing/restful_api_unit_testing/models"
	"bytes"
	"fmt"
	"strconv"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {

	if Addition(2, 4) != 6 {
		t.Error("2+4 seharusnya 6")
	}

	if Addition(-2, 4) != 2 {
		t.Error("-2+4 seharusnya 2")
	}
}

func InitEchoTestAPI() *echo.Echo {
	database.InitTestDB()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := models.User{
		Name:     "Alta",
		Email:    "alta@gmail.com",
		Password: "123",
	}

	var err error
	if err = database.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func InsertDataUserForGetBooks() error {
	book := models.Book{
		Name:        "bookname",
		Genre:       "bookgenre",
		Description: "bookdesc",
	}

	var err error
	if err = database.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func TestCreateUserController(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}}

	e := InitEchoTestAPI()
	jsonBody := []byte(`{"name":"tess","email:"asdad@gmail.com",password:"12132"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/users", bodyReader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{
		{
			name:                   "get user normal",
			path:                   "/users",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEchoTestAPI()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}

	}
}

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		unexpectedStatus       int
		expectedBodyStartsWith string
	}{
		{
			name:                   "get user normal",
			path:                   "/books",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEchoTestAPI()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetBookByIDController(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/books",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}}

	e := InitEchoTestAPI()
	book := database.SeedBook()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetUserByIDController(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/users",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"message\":",
	}}

	e := InitEchoTestAPI()
	book := database.SeedUser()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestLoginUser_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/login",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"token\":",
	},
	}

	e := InitEchoTestAPI()

	user := database.SeedUser()
	fmt.Println(user.Email, user.Password)

	jsonBody := []byte(`{"email":"` + user.Email + `","password":"` + user.Password + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/login", bodyReader)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}

	}

}

func TestLoginUser_Failed(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "failed",
		path:                   "/login",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"token\":",
	},
	}

	e := InitEchoTestAPI()
	jsonBody := []byte(`{"email":"wrong@mail.com","password":"wrong"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/login", bodyReader)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		err := Login(c)

		if assert.Error(t, err) {
			expected := echo.NewHTTPError(http.StatusUnauthorized, "login failed")
			assert.Equal(t, expected, err)
		}

	}

}
