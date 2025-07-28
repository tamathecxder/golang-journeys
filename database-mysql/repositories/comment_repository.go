package repositories

import (
	"context"
	"golang-journeys/database-mysql/entities"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entities.Comment) (entities.Comment, error)
	FindById(ctx context.Context, id int64) (entities.Comment, error)
	FindAll(ctx context.Context) ([]entities.Comment, error)
}
