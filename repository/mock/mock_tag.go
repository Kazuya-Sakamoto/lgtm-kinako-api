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

func (m *MockTagRepository) FindAll(tags *[]domain.Tag) error {
	args := m.Called(tags)
	if args.Get(0) != nil {
		*tags = args.Get(0).([]domain.Tag)
	}
	return args.Error(0)
}

func (m *MockTagRepository) Create(tag *domain.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}

func (m *MockTagRepository) DeleteByTagID(tagId uint) error {
	args := m.Called(tagId)
	return args.Error(0)
}

/*
handler
*/
type MockTagHandler struct {
	mock.Mock
}

func (m *MockTagHandler) TagHandler(tag domain.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}
