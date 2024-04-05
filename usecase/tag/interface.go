package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type ITagUsecase interface {
	GetTags() ([]domain.TagResponse, error)
	CreateTag(tag domain.Tag) (domain.TagResponse, error)
	DeleteTag(tagId uint) error
}

type TagUsecase struct {
	*GetTagsUsecase
	*CreateTagUsecase
	*DeleteTagUsecase
}

func NewTagUsecase(
	re repository.ITagRepository,
	ha handler.ITagHandler,
) *TagUsecase {
	return &TagUsecase{
		NewGetTagsUsecase(re),
		NewCreateTagUsecase(re, ha),
		NewDeleteTagUsecase(re),
	}
}
