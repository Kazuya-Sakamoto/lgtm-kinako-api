package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type GetTagsUsecase struct {
	re repository.ITagRepository
}

func NewGetTagsUsecase(re repository.ITagRepository) *GetTagsUsecase {
	return &GetTagsUsecase{re}
}

func (u *GetTagsUsecase) GetTags() ([]domain.TagResponse, error) {
	tags := []domain.Tag{}
	if err := u.re.FindAll(&tags); err != nil {
		return nil, err
	}
	res := []domain.TagResponse{}
	for _, v := range tags {
		t := domain.TagResponse(v)
		res = append(res, t)
	}
	return res, nil
}
