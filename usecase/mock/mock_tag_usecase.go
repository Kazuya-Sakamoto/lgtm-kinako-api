package mock

import (
	"lgtm-kinako-api/domain"

	"github.com/stretchr/testify/mock"
)

type MockTagUsecase struct {
    mock.Mock
}

func (mu *MockTagUsecase) GetTags() ([]domain.TagResponse, error) {
    args := mu.Called()
    return args.Get(0).([]domain.TagResponse), args.Error(1)
}