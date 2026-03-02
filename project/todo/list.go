package todo

import "github.com/k0kubun/pp"

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) {
	l.tasks[task.Title] = task

}

func (l *List) Help() {

}

func (l *List) Del() {

}

func (l *List) ListTasks() map[string]Task {
	return l.tasks
}

func (l *List) Done(title string) {
	tasks, ok := l.tasks[title]
	if !ok {
		pp.Println("well, there's an error, no such task found")
	}
	tasks.DoneTask()
}

func (l *List) Events() {

}

func (l *List) Exit() {

}
