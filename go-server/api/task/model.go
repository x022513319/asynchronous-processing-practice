// Task schema
// v1先不分 entity 與 DTO
package task

import (
	"encoding/json"
	"time"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

// 傳輸用的結構 (越輕量越好，節省 Redis 記憶體)
type TaskMessage struct {
	ID      string          `json:"id"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// 資料庫用的結構 (包含生命週期管理)(Domain Model)
type Task struct {
	TaskMessage
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
