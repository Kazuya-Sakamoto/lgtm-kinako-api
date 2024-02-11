package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type GetTagsUsecase struct {
	tr repository.ITagRepository
}

func NewGetTagsUsecase(tr repository.ITagRepository) *GetTagsUsecase {
	return &GetTagsUsecase{tr}
}

func (tu *GetTagsUsecase) GetTags() ([]domain.TagResponse, error) {
	tags := []domain.Tag{}
	if err := tu.tr.GetTags(&tags); err != nil {
		return nil, err
	}
	res := []domain.TagResponse{}
	for _, v := range tags {
		t := domain.TagResponse(v)
		res = append(res, t)
	}
	return res, nil
}