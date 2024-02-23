package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type CreateTagUsecase struct {
	tr repository.ITagRepository
	th handler.ITagHandler
}

func NewCreateTagUsecase(tr repository.ITagRepository, th handler.ITagHandler) *CreateTagUsecase {
	return &CreateTagUsecase{tr, th}
}

func (tu *CreateTagUsecase) CreateTag(tag domain.Tag) (domain.TagResponse, error) {
	if err := tu.th.TagHandler(tag); err != nil {
		return domain.TagResponse{}, err
	}
	if err := tu.tr.CreateTag(&tag); err != nil {
		return domain.TagResponse{}, err
	}
	res := domain.TagResponse(tag)

	return res, nil
}
