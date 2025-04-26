package content

import (
	"context"
	"promotion/pkg/logger"
)

type Service struct {
	log           *logger.Logger
	repo          *Repo
	lessonManager *LessonManager
}

func NewService(log *logger.Logger, repo *Repo, generator LessonGeneratorStrategy) *Service {
	lessonManager := NewLessonManager(generator)

	return &Service{
		log:           log,
		repo:          repo,
		lessonManager: lessonManager,
	}
}

func (s *Service) GenerateLessonFromCode(ctx context.Context, code string) (*Lesson, error) {
	return s.lessonManager.GenerateLesson(ctx, code)
}
