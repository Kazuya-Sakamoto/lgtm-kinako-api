package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

type MockTagUsecase struct {
	mock.Mock
}

func (m *MockTagUsecase) GetTags() ([]domain.TagResponse, error) {
	args := m.Called()
	return args.Get(0).([]domain.TagResponse), args.Error(1)
}

func (m *MockTagUsecase) CreateTag(tag domain.Tag) (domain.TagResponse, error) {
	args := m.Called(tag)
	return args.Get(0).(domain.TagResponse), args.Error(1)
}

func (m *MockTagUsecase) DeleteTag(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}
