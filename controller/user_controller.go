// lgtm-kinako-api/controller/user_controller.go
package controller

import (
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/usecase/user"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
    SignUp(c echo.Context) error
    LogIn(c echo.Context) error
    CsrfToken(c echo.Context) error
}

type userController struct {
    uu user.IUserUsecase
}

func NewUserController(uu user.IUserUsecase) IUserController {
    return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
    user := model.User{}
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    userRes, err := uc.uu.SignUp(user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
    user := model.User{}
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    tokenString, err := uc.uu.Login(user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    setTokenCookie(c, tokenString)

    return c.JSON(http.StatusOK, tokenString)
}

func (uc *userController) CsrfToken(c echo.Context) error {
    token := c.Get("csrf").(string)
    return c.JSON(http.StatusOK, echo.Map{
        "csrf_token": token,
    })
}

func setTokenCookie(c echo.Context, tokenString string) {
    cookie := new(http.Cookie)
    cookie.Name = "token"
    cookie.Value = tokenString
    cookie.Expires = time.Now().Add(24 * time.Hour)
    cookie.Path = "/"
    cookie.Domain = os.Getenv("API_DOMAIN")
    cookie.Secure = true
    cookie.HttpOnly = true
    cookie.SameSite = http.SameSiteNoneMode
    c.SetCookie(cookie)
}