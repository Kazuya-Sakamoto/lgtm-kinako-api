package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	testify_mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupGetTagsUsecase(t *testing.T) (*mock.MockTagRepository, *GetTagsUsecase, func()) {
	tr := new(mock.MockTagRepository)
	usecase := NewGetTagsUsecase(tr)

	return tr, usecase, func() {
		tr.AssertExpectations(t)
	}
}

func Test_TagUsecase_GetTags(t *testing.T) {
	t.Run("正常にタグが取得できること", func(t *testing.T) {
		mr, usecase, cleanup := setupGetTagsUsecase(t)
		defer cleanup()

		ea := []domain.Tag{
			{
				ID:        1,
				Name:      "Name1",
			},
			{
				ID:        2,
				Name:      "Name2",
			},
		}

		mr.On("GetTags", testify_mock.AnythingOfType("*[]domain.Tag")).Return(nil).Run(func(args testify_mock.Arguments) {
			arg := args.Get(0).(*[]domain.Tag)
			*arg = ea
		})

		res, err := usecase.GetTags()

		require.NoError(t, err)
		assert.Equal(t, len(ea), len(res))
		for i, r := range res {
			assert.Equal(t, ea[i].ID, r.ID)
			assert.Equal(t, ea[i].Name, r.Name)
		}
	})
}
