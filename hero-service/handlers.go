package main

import (
	"github.com/lfordyce/hero_cqrs/db"
	"github.com/lfordyce/hero_cqrs/event"
	"github.com/lfordyce/hero_cqrs/schema"
	"github.com/lfordyce/hero_cqrs/util"
	"github.com/segmentio/ksuid"
	"html/template"
	"log"
	"net/http"
	"time"
)

func createHeroHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}

	ctx := r.Context()

	// Read parameters
	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	// Create hero
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create hero")
		return
	}
	hero := schema.Hero{
		ID:        id.String(),
		Body:      body,
		CreatedAt: createdAt,
	}
	if err := db.InsertHero(ctx, hero); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create hero")
		return
	}

	// Publish event
	if err := event.PublishHeroCreated(hero); err != nil {
		log.Println(err)
	}

	// Return new hero
	util.ResponseOk(w, response{ID: hero.ID})
}