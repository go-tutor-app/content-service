package content

import "context"

type LessonManager struct {
	generator LessonGeneratorStrategy
}

func NewLessonManager(generator LessonGeneratorStrategy) *LessonManager {
	return &LessonManager{generator: generator}
}

func (m *LessonManager) GenerateLesson(ctx context.Context, input string) (*Lesson, error) {
	return m.generator.Generate(ctx, input)
}
