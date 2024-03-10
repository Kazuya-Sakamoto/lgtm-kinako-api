package tag

import (
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupCreateTagUsecase(t *testing.T) (*mock.MockTagRepository, *mock.MockTagHandler, *CreateTagUsecase, func()) {
	tr := new(mock.MockTagRepository)
	th := new(mock.MockTagHandler)
	usecase := NewCreateTagUsecase(tr, th)

	return tr, th, usecase, func() {
		tr.AssertExpectations(t)
		th.AssertExpectations(t)
	}
}

func Test_TagUsecase_CreateTag(t *testing.T) {
	t.Run("タグが正常に作成されること", func(t *testing.T) {
		tr, th, usecase, cleanup := setupCreateTagUsecase(t)
		defer cleanup()

		tag := domain.Tag{
			ID:   1,
			Name: "TestTag",
		}

		th.On("TagHandler", tag).Return(nil)
		tr.On("CreateTag", &tag).Return(nil)

		res, err := usecase.CreateTag(tag)

		require.NoError(t, err)
		require.Equal(t, tag.ID, res.ID)
		require.Equal(t, tag.Name, res.Name)
	})
}
