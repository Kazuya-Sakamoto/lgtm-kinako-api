package tag

import (
	"testing"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupCreateTagUsecase(t *testing.T) (*mock.MockTagRepository, *mock.MockTagHandler, *CreateTagUsecase, func()) {
	re := new(mock.MockTagRepository)
	ha := new(mock.MockTagHandler)
	usecase := NewCreateTagUsecase(re, ha)

	return re, ha, usecase, func() {
		re.AssertExpectations(t)
		ha.AssertExpectations(t)
	}
}

func Test_TagUsecase_CreateTag(t *testing.T) {
	t.Run("タグが正常に作成されること", func(t *testing.T) {
		re, ha, usecase, cleanup := setupCreateTagUsecase(t)
		defer cleanup()

		tag := domain.Tag{
			ID:   1,
			Name: "TestTag",
		}

		ha.On("TagHandler", tag).Return(nil)
		re.On("Create", &tag).Return(nil)

		res, err := usecase.CreateTag(tag)

		require.NoError(t, err)
		require.Equal(t, tag.ID, res.ID)
		require.Equal(t, tag.Name, res.Name)
	})
}
