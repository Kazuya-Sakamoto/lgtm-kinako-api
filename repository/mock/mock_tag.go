package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

/*
	repository
*/
type MockTagRepository struct {
	mock.Mock
}

func (m *MockTagRepository) GetTags(tags *[]domain.Tag) error {
    args := m.Called(tags)
    if args.Get(0) != nil {
        *tags = args.Get(0).([]domain.Tag)
    }
    return args.Error(0)
}