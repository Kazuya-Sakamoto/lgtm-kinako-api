package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type GetTagsUsecase struct {
	tr repository.ITagRepository
	th handler.ITagHandler
}

func NewGetTagsUsecase(tr repository.ITagRepository, th handler.ITagHandler) *GetTagsUsecase {
	return &GetTagsUsecase{tr, th}
}

func (tu *GetTagsUsecase) GetTags() ([]domain.TagResponse, error) {
	tags := []domain.Tag{}
	if err := tu.tr.FindAll(&tags); err != nil {
		return nil, err
	}
	res := []domain.TagResponse{}
	for _, v := range tags {
		t := domain.TagResponse(v)
		res = append(res, t)
	}
	return res, nil
}
