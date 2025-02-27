// This file contains the repository implementation layer.
package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type NewRepositoryOptions struct {
	DB *gorm.DB
}

func NewRepository(opts NewRepositoryOptions) *Repository {

	return &Repository{
		DB: opts.DB,
	}
}
