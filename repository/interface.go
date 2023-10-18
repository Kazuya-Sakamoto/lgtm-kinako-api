package repository

import "lgtm-kinako-api/domain"

type AllRepository interface {
	IUserRepository
	IAlbumRepository
}

type IUserRepository interface {
	domain.IUserRepository
}

type IAlbumRepository interface {
	domain.IAlbumRepository
}