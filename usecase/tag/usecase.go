package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/repository"
)

type TagUsecase struct {
	GetTagsUsecase 	*GetTagsUsecase
}

func NewTagUsecase(tr repository.ITagRepository) *TagUsecase {
	return &TagUsecase{
		GetTagsUsecase: 	NewGetTagsUsecase(tr),
	}
}

func (tu *TagUsecase) GetTags() ([]domain.TagResponse, error) {
	return tu.GetTagsUsecase.GetTags()
}
