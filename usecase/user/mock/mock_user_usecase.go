package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (mu *MockUserUsecase) SignUp(user domain.User) (domain.UserResponse, error) {
	args := mu.Called(user)
	return args.Get(0).(domain.UserResponse), args.Error(1)
}

func (mu *MockUserUsecase) Login(user domain.User) (string, error) {
	args := mu.Called(user)
	return args.String(0), args.Error(1)
}

