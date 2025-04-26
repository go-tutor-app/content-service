// generator_strategy.go
package content

import (
	"context"
)

// LessonGeneratorStrategy defines the interface that all generators must implement
type LessonGeneratorStrategy interface {
	Generate(ctx context.Context, data string) (*Lesson, error)
}
