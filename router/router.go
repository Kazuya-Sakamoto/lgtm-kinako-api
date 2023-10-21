package router

import (
	"lgtm-kinako-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ac controller.IAlbumController, uc controller.IUserController) *echo.Echo {
	logConfig := middleware.LoggerConfig{
		Format: "${host}${uri} METHOD:${method} STATUS:${status} TIME:${time_rfc3339} LATENCY:${latency_human} ERROR:${error}\n",
	}

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(logConfig))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieSecure: true,
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))
	// * User
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.GET("/csrf", uc.CsrfToken)
	// * Album
	a := e.Group("/album")
	a.GET("/random", ac.GetRandomAlbums)
	a.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	a.GET("", ac.GetAllAlbums)
	a.POST("", ac.CreateAlbum)
	a.DELETE("/:albumId", ac.DeleteAlbum)

	return e
}