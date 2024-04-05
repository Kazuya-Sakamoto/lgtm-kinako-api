package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) SignUp(user domain.User) (domain.UserResponse, error) {
	args := m.Called(user)
	return args.Get(0).(domain.UserResponse), args.Error(1)
}

func (m *MockUserUsecase) Login(user domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
