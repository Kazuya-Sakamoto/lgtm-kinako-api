package tag

import (
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type DeleteTagUsecase struct {
	tr repository.ITagRepository
	th handler.ITagHandler
}

func NewDeleteTagUsecase(tr repository.ITagRepository, th handler.ITagHandler) *DeleteTagUsecase {
	return &DeleteTagUsecase{tr, th}
}

func (tu *DeleteTagUsecase) DeleteTag(tagId uint) error {
	if err := tu.tr.DeleteByTagID(tagId); err != nil {
		return err
	}
	return nil
}
