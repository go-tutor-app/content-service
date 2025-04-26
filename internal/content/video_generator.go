package content

import (
	"context"
	"fmt"
)

type VideoLessonGenerator struct{}

func (g *VideoLessonGenerator) Generate(ctx context.Context, data string) (*Lesson, error) {
	return &Lesson{
		Type:    "video",
		Content: fmt.Sprintf("Video for topic %s", data),
	}, nil
}
