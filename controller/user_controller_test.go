package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/mock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupUserControllerTest(mockUsecase *mock.MockUserUsecase, method, url string, body []byte) (*echo.Echo, *httptest.ResponseRecorder, echo.Context, IUserController) {
	controller := NewUserController(mockUsecase)
	e := echo.New()
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return e, rec, ctx, controller
}

func Test_UserController_CsrfToken(t *testing.T) {
	usecase := new(mock.MockUserUsecase)
	csrf_token := "expect_csrf_token"

	_, res, c, controller := setupUserControllerTest(usecase, http.MethodGet, "/csrf", nil)
	c.Set("csrf", csrf_token)

	if assert.NoError(t, controller.CsrfToken(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		var response map[string]string
		err := json.Unmarshal(res.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, csrf_token, response["csrf_token"])
	}
}

func Test_UserController_LogIn(t *testing.T) {
	usecase := new(mock.MockUserUsecase)
	token := "expected_token"
	user := domain.User{Email: "test@test.com", Password: "password"}
	usecase.On("Login", user).Return(token, nil)

	userJSON, _ := json.Marshal(user)
	_, res, c, controller := setupUserControllerTest(usecase, http.MethodPost, "/login", userJSON)

	if assert.NoError(t, controller.LogIn(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		response := strings.Trim(res.Body.String(), "\"\n")
		assert.Equal(t, token, response)
	}
}
