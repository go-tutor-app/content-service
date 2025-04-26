package content

import (
	"context"
	"promotion/pkg/databases"

	"gorm.io/gorm"
)

type Repo struct {
	db databases.MySQLDB
}

func newRepo(db databases.MySQLDB) *Repo {
	return &Repo{db}
}

func (r *Repo) dbWithContext(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}
