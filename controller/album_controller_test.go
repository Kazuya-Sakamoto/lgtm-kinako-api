package controller

import (
	"encoding/json"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/usecase"
	"lgtm-kinako-api/usecase/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAlbumController_GetAllAlbums(t *testing.T) {
    expectedAlbums := []model.AlbumResponse{
        {
            ID:        1,
            Title:     "Album 1",
            Image:     "image1.jpg",
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
    }

    mockAlbumUsecase := new(mock.MockAlbumUsecase)
    mockAlbumUsecase.On("GetAllAlbums").Return(expectedAlbums, nil)
    controller := NewAlbumController(usecase.IAlbumUsecase(mockAlbumUsecase), nil)

    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/albums", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, controller.GetAllAlbums(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)

        var albums []model.AlbumResponse
        err := json.Unmarshal(rec.Body.Bytes(), &albums)
        assert.NoError(t, err)
        assert.Len(t, albums, len(expectedAlbums), "Unexpected number of albums")

		for i, expected := range expectedAlbums {
			assert.Equal(t, expected.ID, albums[i].ID, "Unexpected album ID at index", i)
			assert.Equal(t, expected.Title, albums[i].Title, "Unexpected album title at index", i)
			assert.Equal(t, expected.Image, albums[i].Image, "Unexpected album image at index", i)
		
			//* 時間に依存しないCreatedAtおよびUpdatedAtフィールドの比較
			assert.True(t, expected.CreatedAt.Equal(albums[i].CreatedAt), "Unexpected album CreatedAt at index", i)
			assert.True(t, expected.UpdatedAt.Equal(albums[i].UpdatedAt), "Unexpected album UpdatedAt at index", i)
		}
    }
}

func TestAlbumController_GetRandomAlbums(t *testing.T) {
    expectedAlbums := []model.AlbumResponse{
        {
            ID:        1,
            Title:     "Album 1",
            Image:     "image1.jpg",
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
        {
            ID:        2,
            Title:     "Album 2",
            Image:     "image2.jpg",
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
    }

    mockAlbumUsecase := new(mock.MockAlbumUsecase)
    mockAlbumUsecase.On("GetRandomAlbums").Return(expectedAlbums, nil)
    controller := NewAlbumController(usecase.IAlbumUsecase(mockAlbumUsecase), nil)

    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/random-albums", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, controller.GetRandomAlbums(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)

        var albums []model.AlbumResponse
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

