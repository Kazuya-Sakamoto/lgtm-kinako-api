// mocks.go
package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

/*
	repository
*/
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(user *domain.User, email string) error {
	args := m.Called(user, email)
	if args.Get(0) != nil {
		*user = args.Get(0).(domain.User)
	}
	return args.Error(1)
}
/*
	handler
*/

type MockUserHandler struct {
	mock.Mock
}

func (m *MockUserHandler) UserHandler(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}
