package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/album_tag"
	"lgtm-kinako-api/usecase/mock"
	"lgtm-kinako-api/usecase/tag"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupTagControllerTest(mockTagUsecase *mock.MockTagUsecase, mockAlbumTagUsecase *mock.MockAlbumTagUsecase, url string) (*echo.Echo, *httptest.ResponseRecorder, echo.Context, ITagController) {
	controller := NewTagController(tag.ITagUsecase(mockTagUsecase), album_tag.IAlbumTagUsecase(mockAlbumTagUsecase))
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return e, rec, c, controller
}

func Test_TagController_GetTags(t *testing.T) {
	e := []domain.TagResponse{
		{
			ID:   1,
			Name: "tag1",
		},
		{
			ID:   2,
			Name: "tag2",
		},
	}
	mt := new(mock.MockTagUsecase)
	mat := new(mock.MockAlbumTagUsecase)
	mt.On("GetTags").Return(e, nil)
	_, rec, c, controller := setupTagControllerTest(mt, mat, "/tags")

	if assert.NoError(t, controller.GetTags(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var tags []domain.TagResponse
		err := json.Unmarshal(rec.Body.Bytes(), &tags)
		assert.NoError(t, err)

		assert.Len(t, tags, len(e), "Unexpected number of tags")
		for i, expected := range e {
			assert.Equal(t, expected.ID, tags[i].ID, "Unexpected tag ID at index", i)
		}
	}
}
