package tasks

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/x022513319/asynchronous-processing-practice/api/task"
)

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.ID = uuid.NewString()
	t.Status = task.StatusPending
	t.CreatedAt = time.Now().UTC()

	/*
		API 層 → 傳結構化的 Task
		Queue / Redis 層 → 負責序列化成 []byte
	*/
	if err := h.publisher.Publish(r.Context(), t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("task enqueued: " + t.ID)
}
