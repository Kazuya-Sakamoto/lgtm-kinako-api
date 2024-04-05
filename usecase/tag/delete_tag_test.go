package tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupDeleteTagUsecase(t *testing.T) (*mock.MockTagRepository, *DeleteTagUsecase, func()) {
	re := new(mock.MockTagRepository)
	usecase := NewDeleteTagUsecase(re)

	return re, usecase, func() {
		re.AssertExpectations(t)
	}
}

func Test_TagUsecase_DeleteTag(t *testing.T) {
	t.Run("タグが正常に削除されること", func(t *testing.T) {
		re, usecase, cleanup := setupDeleteTagUsecase(t)
		defer cleanup()

		tagId := uint(1)

		re.On("DeleteByTagID", tagId).Return(nil)

		err := usecase.DeleteTag(tagId)

		require.NoError(t, err)
	})
}
