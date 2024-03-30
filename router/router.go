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

	// * Albums Routes
	albums := v1.Group("/albums")
	albums.GET("", ac.GetAlbums)
	albums.GET("/random", ac.GetRandomAlbums)
	albums.GET("/tags/count", atu.GetAlbumCountsByTag)
	albumsSecured := albums.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	albumsSecured.GET("/all", ac.GetAllAlbums)
	albumsSecured.POST("", ac.CreateAlbum)
	albumsSecured.DELETE("/:albumId", ac.DeleteAlbum)
	albumsSecured.POST("/tags/update", atu.DeleteAndInsertAlbumTags)

	// * Tag Routes
	tags := v1.Group("/tags")
	tags.GET("", tc.GetTags)
	tagsSecured := tags.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	tagsSecured.POST("", tc.CreateTag)
	tagsSecured.DELETE("/:tagId", tc.DeleteTag)

	return e
}
