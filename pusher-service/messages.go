package main

import "time"

const (
	KindHeroCreated = iota + 1
)

type HeroCreatedMessage struct {
	Kind uint32 `json:"kind"`
	ID string `json:"id"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func newHeroCreatedMessage(id string, body string, createdAt time.Time) *HeroCreatedMessage {
	return &HeroCreatedMessage{
		Kind: KindHeroCreated,
		ID:id,
		Body:body,
		CreatedAt:createdAt,
	}
}