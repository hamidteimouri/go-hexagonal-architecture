package postgres

import (
	"github.com/hamidteimouri/go-hexagonal-architecture/data/models"
	"gorm.io/gorm"
	"time"
)

type PgDataSource interface {
	CreateTable() *error
	InsertBlog(model *models.BlogModel) (int64, *error)
}

type pgDataSource struct {
	db *gorm.DB
}

func (p pgDataSource) CreateTable() *error {
	err := p.db.AutoMigrate(&PgBlogModel{})
	if err != nil {
		return &err
	}
	return nil
}

func (p pgDataSource) InsertBlog(b *models.BlogModel) (int64, *error) {

	b.ID = 0

	blog := new(PgBlogModel)

	blog.Title = b.Title
	blog.Description = b.Description
	blog.CreatedAt = b.CreatedAt
	blog.UpdatedAt = b.UpdatedAt

	p.db.Model().cre

}
