package repository

import "lgtm-kinako-api/domain"

type IUserRepository interface {
	domain.IUserRepository
}

type IAlbumRepository interface {
	domain.IAlbumRepository
}

type ITagRepository interface {
	domain.ITagRepository
}

type IAlbumTagRepository interface {
	domain.IAlbumTagRepository
}
