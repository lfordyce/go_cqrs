package search

import (
	"context"
	"encoding/json"
	"github.com/lfordyce/hero_cqrs/schema"
	"github.com/olivere/elastic"
	"log"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) InsertHero(ctx context.Context, hero schema.Hero) error {
	_, err := r.client.Index().
		Index("heros").
		Type("hero").
		Id(hero.ID).
		BodyJson(hero).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *ElasticRepository) SearchHeros(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Hero, error) {
	result, err := r.client.Search().
		Index("heros").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	heros := []schema.Hero{}
	for _, hit := range result.Hits.Hits {
		var hero schema.Hero
		if err = json.Unmarshal(*hit.Source, &hero); err != nil {
			log.Println(err)
		}
		heros = append(heros, hero)
	}
	return heros, nil
}