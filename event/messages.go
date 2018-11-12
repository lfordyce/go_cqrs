package event

import "time"

type Message interface {
	Key() string
}

type HeroCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

func (h *HeroCreatedMessage) Key() string {
	return "hero.created"
}
