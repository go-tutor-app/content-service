package content

type Lesson struct {
	Code    string
	Type    string
	Content string
}

func (u *Lesson) TableName() string {
	return "lesson"
}
