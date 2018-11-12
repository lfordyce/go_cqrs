package event

import "github.com/lfordyce/hero_cqrs/schema"

type EventStore interface {
	Close()
	PublishHeroCreated(h schema.Hero) error
	SubscribeHeroCreated() (<- chan HeroCreatedMessage, error)
	OnHeroCreated(f func(HeroCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishHeroCreated(h schema.Hero) error{
	return impl.PublishHeroCreated(h)
}

func SubscribeHeroCreated() (<- chan HeroCreatedMessage, error) {
	return impl.SubscribeHeroCreated()
}

func OnHeroCreated(f func(HeroCreatedMessage)) error {
	return impl.OnHeroCreated(f)
}