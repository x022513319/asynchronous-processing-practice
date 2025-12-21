// Task schema
// v1先不分 entity 與 DTO
package task

import "encoding/json"

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

type Task struct {
	ID      string          `json:"id"`
	Type    string          `json:"type"`
	Status  Status          `json:"status"`
	Payload json.RawMessage `json:"payload"`
}
