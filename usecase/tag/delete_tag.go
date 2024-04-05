package tag

import (
	"lgtm-kinako-api/repository"
)

type DeleteTagUsecase struct {
	re repository.ITagRepository
}

func NewDeleteTagUsecase(re repository.ITagRepository) *DeleteTagUsecase {
	return &DeleteTagUsecase{re}
}

func (u *DeleteTagUsecase) DeleteTag(tagId uint) error {
	if err := u.re.DeleteByTagID(tagId); err != nil {
		return err
	}
	return nil
}
