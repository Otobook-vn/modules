package audit

import (
	"context"

	"github.com/kr/pretty"
)

// CreatePayload ...
type CreatePayload struct {
	Target   string
	TargetID string
	Action   string
	Data     string
	Author   CreatePayloadAuthor
}

// CreatePayloadAuthor ...
type CreatePayloadAuthor struct {
	ID   string
	Name string
}

// Create ...
func (s Service) Create(payload CreatePayload) {
	ctx := context.Background()

	// Get document
	doc := Audit{
		ID:       newMongoID(),
		Source:   s.Source,
		Target:   payload.Target,
		TargetID: payload.TargetID,
		Action:   payload.Action,
		Data:     payload.Data,
		Author: Author{
			ID:   payload.Author.ID,
			Name: payload.Author.Name,
		},
		CreatedAt: timeNow(),
	}

	// Insert to db
	colName := s.Source + "_" + payload.Target
	if _, err := s.MongoDB.Collection(colName).InsertOne(ctx, doc); err != nil {
		pretty.Println("*** Audit create log error", err.Error())
		pretty.Println("*** Payload", payload)
		pretty.Println("*****************")
	}
}
