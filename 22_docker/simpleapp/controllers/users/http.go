package users

import (
  // "Go_Aldimas/21_clean_architecture/clean_architecture/app/middlewares"
  "Go_Aldimas/22_docker/simpleapp/businesses/users"
  "Go_Aldimas/22_docker/simpleapp/controllers"
  "Go_Aldimas/22_docker/simpleapp/controllers/users/request"
  "Go_Aldimas/22_docker/simpleapp/controllers/users/response"
  "net/http"
	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
  userUsecase users.Usecase
}

func NewUserController(userUC users.Usecase) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (ctrl *UserController) GetAll(c echo.Context) error {
	usersData := ctrl.userUsecase.GetAll()

	users := []response.User{}

	for _, user := range usersData {
    users = append(users, response.FromDomain(user))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all users", users)
}

func (ctrl *UserController) Register(c echo.Context) error {
	input := request.User{}
	
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	user := ctrl.userUsecase.Register(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "user registered", response.FromDomain(user))
}

func (ctrl *UserController) Login(c echo.Context) error {
	userInput := request.UserLogin{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token := ctrl.userUsecase.Login(userInput.ToDomain())
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

// func (ctrl *UserController) Logout(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)

// 	if isListed := middlewares.CheckToken(user.Raw); !isListed {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "invalid token",
// 		})
// 	}
// 	middlewares.Logout(user.Raw)
// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "logout succes",
// 	})
// }