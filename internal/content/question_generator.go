package content

import (
	"context"
	"fmt"
)

type QuestionLessonGenerator struct{}

func (g *QuestionLessonGenerator) Generate(ctx context.Context, data string) (*Lesson, error) {
	return &Lesson{
		Type:    "question",
		Content: fmt.Sprintf("Generated question from %s", data),
	}, nil
}
