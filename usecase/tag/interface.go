package tag

import "lgtm-kinako-api/domain"

type ITagUsecase interface {
	GetTags() ([]domain.TagResponse, error)
	CreateTag(tag domain.Tag) (domain.TagResponse, error)
	DeleteTag(tagId uint) error
}
