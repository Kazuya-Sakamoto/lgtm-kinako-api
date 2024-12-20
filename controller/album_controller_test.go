package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/album"
	"lgtm-kinako-api/usecase/mock"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupAlbumControllerTest(mockUsecase *mock.MockAlbumUsecase, url string) (*echo.Echo, *httptest.ResponseRecorder, echo.Context, IAlbumController) {
	controller := NewAlbumController(album.IAlbumUsecase(mockUsecase), nil)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return e, rec, ctx, controller
}

func setJWTToken(ctx echo.Context, userID uint) {
	ctx.Set("user", &jwt.Token{
		Claims: jwt.MapClaims{
			"user_id": float64(userID),
		},
	})
}

func Test_AlbumController_GetAllAlbums(t *testing.T) {
	expectedAlbums := []domain.AlbumResponse{
		{
			ID:        1,
			Title:     "title2",
			Image:     "image1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "title2",
			Image:     "image2.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	usecase := new(mock.MockAlbumUsecase)
	usecase.On("GetAllAlbums").Return(expectedAlbums, nil)
	_, rec, ctx, controller := setupAlbumControllerTest(usecase, "/album")
	setJWTToken(ctx, 1)

	if assert.NoError(t, controller.GetAllAlbums(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var albums []domain.AlbumResponse
		err := json.Unmarshal(rec.Body.Bytes(), &albums)
		assert.NoError(t, err)

		assert.Len(t, albums, len(expectedAlbums), "Unexpected number of albums")
		for i, expected := range expectedAlbums {
			assert.Equal(t, expected.ID, albums[i].ID, "Unexpected album ID at index", i)
			assert.Equal(t, expected.Title, albums[i].Title, "Unexpected album title at index", i)
			assert.Equal(t, expected.Image, albums[i].Image, "Unexpected album image at index", i)

			assert.True(t, expected.CreatedAt.Equal(albums[i].CreatedAt), "Unexpected album CreatedAt at index", i)
			assert.True(t, expected.UpdatedAt.Equal(albums[i].UpdatedAt), "Unexpected album UpdatedAt at index", i)
		}
	}
}

func Test_AlbumController_GetRandomAlbums(t *testing.T) {
	expectedAlbums := []domain.AlbumResponse{
		{
			ID:        1,
			Title:     "title2",
			Image:     "image1.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "title2",
			Image:     "image2.jpg",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	usecase := new(mock.MockAlbumUsecase)
	usecase.On("GetRandomAlbums").Return(expectedAlbums, nil)
	_, rec, ctx, controller := setupAlbumControllerTest(usecase, "/album/random")

	if assert.NoError(t, controller.GetRandomAlbums(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var albums []domain.AlbumResponse
		err := json.Unmarshal(rec.Body.Bytes(), &albums)
		assert.NoError(t, err)

		assert.Len(t, albums, len(expectedAlbums), "Unexpected number of albums")
		for i, expected := range expectedAlbums {
			assert.Equal(t, expected.ID, albums[i].ID, "Unexpected album ID at index", i)
			assert.Equal(t, expected.Title, albums[i].Title, "Unexpected album title at index", i)
			assert.Equal(t, expected.Image, albums[i].Image, "Unexpected album image at index", i)

			assert.True(t, expected.CreatedAt.Equal(albums[i].CreatedAt), "Unexpected album CreatedAt at index", i)
			assert.True(t, expected.UpdatedAt.Equal(albums[i].UpdatedAt), "Unexpected album UpdatedAt at index", i)
		}
	}
}
