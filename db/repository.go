package db

import (
	"context"
	"github.com/lfordyce/hero_cqrs/schema"
)

type Repository interface {
	Close()
	InsertHero(ctx context.Context, hero schema.Hero) error
	ListHero(ctx context.Context, skip uint64, take uint64) ([]schema.Hero, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertHero(ctx context.Context, hero schema.Hero) error {
	return impl.InsertHero(ctx, hero)
}

func ListHero(ctx context.Context, skip uint64, take uint64) ([]schema.Hero, error) {
	return impl.ListHero(ctx, skip, take)
}
