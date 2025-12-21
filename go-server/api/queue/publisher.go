// 負責將任務推送到 Redis Queue，讓Python Worker取出並執行
package queue

import (
	"context"
	"encoding/json"

	"github.com/x022513319/asynchronous-processing-practice/api/task"
)

type Publisher struct {
	client Client
}

func NewPublisher(client Client) *Publisher {
	return &Publisher{
		client: client,
	}
}

func (p *Publisher) Publish(ctx context.Context, t task.Task) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return p.client.Enqueue(ctx, data)
}
