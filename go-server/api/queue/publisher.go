// 負責將任務推送到 Redis Queue，讓Python Worker取出並執行
package queue

import "context"

type Publisher struct {
	client Client
}

func NewPublisher(client Client) *Publisher {
	return &Publisher{
		client: client,
	}
}

func (p *Publisher) Publish(ctx context.Context, payload []byte) error {
	return p.client.Enqueue(ctx, payload)
}
