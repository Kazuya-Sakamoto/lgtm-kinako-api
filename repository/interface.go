package repository

import "lgtm-kinako-api/domain"

type AllRepository interface {
	IUserRepository
	IAlbumRepository
	ITagRepository
}

type IUserRepository interface {
	domain.IUserRepository
}

type IAlbumRepository interface {
	domain.IAlbumRepository
}

type ITagRepository interface {
	domain.ITagRepository
}