package repositories

import "github.com/hamidteimouri/go-hexagonal-architecture/src/data/datasource/db"

type blogRepository struct {
	dataSource db.DatabaseInterface
}

func (b *blogRepository) Save() {
	_ = b.dataSource.Insert()
}

func (b *blogRepository) Find() {
	b.dataSource.Find()
}
