package content

import (
	"context"
	"database/sql"
	"fmt"
)

type VideoLessonGenerator struct{}

func (g *VideoLessonGenerator) Generate(ctx context.Context, data string) (*Lesson, error) {
	var content string
	query := `SELECT content FROM lessons WHERE topic = $1 AND type = 'video' LIMIT 1`

	err := g.repo.db.QueryRowContext(ctx, query, data).Scan(&content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no lesson found for topic: %s", data)
		}
		return nil, err
	}

	return &Lesson{Content: content}, nil
}
