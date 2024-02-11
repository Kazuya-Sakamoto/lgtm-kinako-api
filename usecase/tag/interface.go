package tag

import "lgtm-kinako-api/domain"

type ITagUsecase interface {
	GetTags() ([]domain.TagResponse, error)
}
