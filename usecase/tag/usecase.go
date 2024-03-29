package tag

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
)

type TagUsecase struct {
	GetTagsUsecase   *GetTagsUsecase
	CreateTagUsecase *CreateTagUsecase
	DeleteTagUsecase *DeleteTagUsecase
}

func NewTagUsecase(tr repository.ITagRepository, th handler.ITagHandler) *TagUsecase {
	return &TagUsecase{
		GetTagsUsecase:   NewGetTagsUsecase(tr, th),
		CreateTagUsecase: NewCreateTagUsecase(tr, th),
		DeleteTagUsecase: NewDeleteTagUsecase(tr, th),
	}
}

func (tu *TagUsecase) GetTags() ([]domain.TagResponse, error) {
	return tu.GetTagsUsecase.GetTags()
}

func (tu *TagUsecase) CreateTag(tag domain.Tag) (domain.TagResponse, error) {
	return tu.CreateTagUsecase.CreateTag(tag)
}

func (tu *TagUsecase) DeleteTag(tagId uint) error {
	return tu.DeleteTagUsecase.DeleteTag(tagId)
}
