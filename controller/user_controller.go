package controller

import (
	"net/http"
	"os"
	"time"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/user"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(ctx echo.Context) error
	LogIn(ctx echo.Context) error
	CsrfToken(ctx echo.Context) error
}

type userController struct {
	uc user.IUserUsecase
}

func NewUserController(uc user.IUserUsecase) IUserController {
	return &userController{uc}
}

func (c *userController) SignUp(ctx echo.Context) error {
	user := domain.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := c.uc.SignUp(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, userRes)
}

func (c *userController) LogIn(ctx echo.Context) error {
	user := domain.User{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := c.uc.Login(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	setTokenCookie(ctx, tokenString)

	return ctx.JSON(http.StatusOK, tokenString)
}

func (c *userController) CsrfToken(ctx echo.Context) error {
	token := ctx.Get("csrf").(string)
	return ctx.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func setTokenCookie(ctx echo.Context, tokenString string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
}
