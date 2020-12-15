package repositories

import (
	"context"

	"github.com/hamidteimouri/go-hexagonal-architecture/src/domain/entities"
)


type BlogRepository interface {
	Create(ctx context.Context, blog entities.Blog) (*entities.Blog, error)


	//Delete(ctx context.Context, b entities.Blog) (*entities.Blog, error)
	//Update(ctx context.Context, b entities.Blog) (*entities.Blog, error)
	//ToggleDisplay(ctx context.Context, b entities.Blog) (*entities.Blog, error)
}
