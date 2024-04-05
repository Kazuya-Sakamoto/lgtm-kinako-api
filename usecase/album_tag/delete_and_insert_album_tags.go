package album_tag

import (
	"lgtm-kinako-api/repository"
)

type DeleteAndInsertAlbumTagsUsecase struct {
	re repository.IAlbumTagRepository
}

func NewDeleteAndInsertAlbumTagsUsecase(re repository.IAlbumTagRepository) *DeleteAndInsertAlbumTagsUsecase {
	return &DeleteAndInsertAlbumTagsUsecase{re}
}

func (u *DeleteAndInsertAlbumTagsUsecase) DeleteAndInsertAlbumTags(albumId uint, tagIds []uint) error {
	if err := u.re.DeleteAndInsert(albumId, tagIds); err != nil {
		return err
	}
	return nil
}
