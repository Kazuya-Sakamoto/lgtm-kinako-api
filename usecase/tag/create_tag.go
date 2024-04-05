package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type CreateTagUsecase struct {
	re repository.ITagRepository
	ha handler.ITagHandler
}

func NewCreateTagUsecase(re repository.ITagRepository, ha handler.ITagHandler) *CreateTagUsecase {
	return &CreateTagUsecase{re, ha}
}

func (u *CreateTagUsecase) CreateTag(tag domain.Tag) (domain.TagResponse, error) {
	if err := u.ha.TagHandler(tag); err != nil {
		return domain.TagResponse{}, err
	}
	if err := u.re.Create(&tag); err != nil {
		return domain.TagResponse{}, err
	}
	res := domain.TagResponse(tag)

	return res, nil
}
