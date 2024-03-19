package tag

import (
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/stretchr/testify/require"
)

func setupDeleteTagUsecase(t *testing.T) (*mock.MockTagRepository, *mock.MockTagHandler, *DeleteTagUsecase, func()) {
	tr := new(mock.MockTagRepository)
	th := new(mock.MockTagHandler)
	usecase := NewDeleteTagUsecase(tr, th)

	return tr, th, usecase, func() {
		tr.AssertExpectations(t)
		th.AssertExpectations(t)
	}
}

func Test_TagUsecase_DeleteTag(t *testing.T) {
	t.Run("タグが正常に削除されること", func(t *testing.T) {
		tr, _, usecase, cleanup := setupDeleteTagUsecase(t)
		defer cleanup()

		tagId := uint(1)

		tr.On("DeleteByTagID", tagId).Return(nil)

		err := usecase.DeleteTag(tagId)

		require.NoError(t, err)
	})
}
