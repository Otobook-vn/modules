package audit

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Otobook-vn/modules/mongodb"
)

// MongoDBConfig ...
type MongoDBConfig struct {
	Host, User, Password, DBName, Mechanism, Source string
}

// Config ...
type Config struct {
	// Source of server, e.g: otobook
	Source string
	// MongoDB config, for save documents
	MongoDB MongoDBConfig
}

// Service ...
type Service struct {
	Config
	DB *mongo.Database
}

var s *Service

// NewInstance ...
func NewInstance(config Config) error {
	if config.Source == "" || config.MongoDB.Host == "" {
		return errors.New("please provide all information that needed: source, mongodb")
	}

	// Connect MongoDB
	err := mongodb.Connect(
		config.MongoDB.Host,
		config.MongoDB.User,
		config.MongoDB.Password,
		config.MongoDB.DBName,
		config.MongoDB.Mechanism,
		config.MongoDB.Source,
	)
	if err != nil {
		fmt.Println("Cannot init module AUDIT", err)
		return err
	}

	s = &Service{
		Config: config,
		DB:     mongodb.GetInstance(),
	}

	// index mongo
	s.indexDB()

	return nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}

// getColName ...
func getColName(source, target string) string {
	return source + "-" + target
}
