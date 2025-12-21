package tasks

import "github.com/x022513319/asynchronous-processing-practice/api/queue"

type TaskHandler struct {
	publisher *queue.Publisher
}

func NewTaskHandler(publisher *queue.Publisher) *TaskHandler {
	return &TaskHandler{
		publisher: publisher,
	}
}
