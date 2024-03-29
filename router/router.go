package router

import (
	"net/http"
	"os"

	"lgtm-kinako-api/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	ac controller.IAlbumController,
	uc controller.IUserController,
	tc controller.ITagController,
	atu controller.IAlbumTagController,
) *echo.Echo {
	logConfig := middleware.LoggerConfig{
		Format: "${host}${uri} METHOD:${method} STATUS:${status} TIME:${time_rfc3339} LATENCY:${latency_human} ERROR:${error}\n",
	}

	e := echo.New()
	//* Logger
	e.Use(middleware.LoggerWithConfig(logConfig))
	e.Use(middleware.Recover())
	//* HSTS
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
			return next(c)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieSecure:   true,
		// CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge:   60,
	}))

	//* API Version 1 Group
	v1 := e.Group("/api/v1")

	// * User Routes
	v1.POST("/signup", uc.SignUp)
	v1.POST("/login", uc.LogIn)
	v1.GET("/csrf", uc.CsrfToken)

	// * Album Routes
	albums := v1.Group("/albums")
	albums.GET("", ac.GetAlbums)
	albums.GET("/random", ac.GetRandomAlbums)
	albums.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	albums.GET("/all", ac.GetAllAlbums)
	albums.POST("", ac.CreateAlbum)
	albums.DELETE("/:albumId", ac.DeleteAlbum)

	// * AlbumTag Routes
	albums.POST("/tags/upsert", atu.ResetAndSetAlbumTags)

	// * Tag Routes
	tags := v1.Group("/tags")
	tags.GET("", tc.GetTags)
	tags.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	tags.POST("", tc.CreateTag)
	tags.DELETE("/:tagId", tc.DeleteTag)

	return e
}
