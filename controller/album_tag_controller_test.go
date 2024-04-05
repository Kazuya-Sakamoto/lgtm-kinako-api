package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/mock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupAlbumTagControllerTest(mockUsecase *mock.MockAlbumTagUsecase) (*echo.Echo, *httptest.ResponseRecorder, echo.Context, IAlbumTagController) {
	controller := NewAlbumTagController(mockUsecase)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/album-tags/counts", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return e, rec, ctx, controller
}

func Test_AlbumTagController_GetAlbumCountsByTag(t *testing.T) {
	expectedCounts := []domain.TagCount{
		{TagID: 1, Count: 5},
		{TagID: 2, Count: 3},
		{TagID: 3, Count: 4},
		{TagID: 4, Count: 0},
	}
	usecase := new(mock.MockAlbumTagUsecase)
	usecase.On("GetAlbumCountsByTag").Return(expectedCounts, nil)

	_, rec, c, controller := setupAlbumTagControllerTest(usecase)

	if assert.NoError(t, controller.GetAlbumCountsByTag(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var counts []domain.TagCount
		err := json.Unmarshal(rec.Body.Bytes(), &counts)
		assert.NoError(t, err)

		assert.Equal(t, expectedCounts, counts, "Returned counts should match the expected counts")
	}

	usecase.AssertExpectations(t)
}
