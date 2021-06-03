// Package classification Tracker.
//
// Terms Of Service:
//
//     Schemes: http
//     Version: 1.0.0
//     Contact: Pavel Aborilov<pavel@aborilov.ru>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"go.mongodb.org/mongo-driver/mongo"

	sessionAPI "github.com/aborilov/tech-challenge-time/backend/v1/api/session"
	"github.com/aborilov/tech-challenge-time/backend/v1/pkg/session"
	sessionRepo "github.com/aborilov/tech-challenge-time/backend/v1/pkg/session/repo/mongo"
	"github.com/gorilla/mux"
)

// Build service and API endpoints
func Build(router *mux.Router, mongoClient *mongo.Client) {
	repo := sessionRepo.NewRepository(mongoClient)
	sessionSvc := session.NewService(repo)
	router = router.PathPrefix("/session").Subrouter()
	sessionAPI.NewHTTPController(sessionSvc, router)
}
