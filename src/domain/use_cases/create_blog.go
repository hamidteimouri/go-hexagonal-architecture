package use_cases

import (
	"context"
	"github.com/hamidteimouri/go-hexagonal-architecture/domain/repositories"

	"github.com/hamidteimouri/go-hexagonal-architecture/domain/entities"
)

type CreateBlogInterface interface {
	Call(ctx context.Context, b entities.Blog) (*entities.Blog, error)
}

func NewBlog(blgRepo repositories.BlogRepository) CreateBlogInterface {
	return createBlog{
		blgRepo,
	}
}

type createBlog struct {
	repo repositories.BlogRepository
}

func (obj createBlog) Call(ctx context.Context, b entities.Blog) (*entities.Blog, error) {
	// todo : log and use error handling here
	blog, err := obj.repo.Create(ctx, b)
	if err != nil {
		return nil, err
	}
	return blog, nil
}
