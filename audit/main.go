package audit

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// Config ...
type Config struct {
	// Source of server, e.g: otobook
	Source string

	// Target of log
	Targets []string

	// MongoDB instance, for save documents
	MongoDB *mongo.Database
}

// Service ...
type Service struct {
	Config
}

// NewInstance ...
func NewInstance(config Config) (*Service, error) {
	if config.Source == "" || config.MongoDB == nil {
		return nil, errors.New("please provide all information that needed: source, mongodb")
	}

	s := Service{
		Config: config,
	}

	if config.MongoDB != nil {
		// index
		s.indexDB()
	}

	return &s, nil
}
