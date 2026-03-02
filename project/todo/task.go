package todo

import "time"

type Task struct {
	Title       string
	Text        string
	CreatedAt   time.Time
	IsDone      bool
	CompletedAt *time.Time
}

func NewTask(title string, text string) Task {
	return Task{
		Title:       title,
		Text:        text,
		CreatedAt:   time.Now(),
		IsDone:      false,
		CompletedAt: nil,
	}
}

func (t *Task) DoneTask(title string) {
	doneTime := time.Now()
	t.IsDone = true
	t.CompletedAt = &doneTime
}
