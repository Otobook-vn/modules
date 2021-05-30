package fcm

import (
	"context"
	b64 "encoding/base64"
	"errors"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/thoas/go-funk"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

// Config ...
type Config struct {
	// Project id
	ProjectID string
	// Original is JSON format, but encoded with base64, need to call base64 decode to get byte data
	Credential string
	// Postgresql instance, for saving logs
	PostgreSQL *gorm.DB
	// Log table name
	LogTableName string
}

// Result of each send
type Result struct {
	BatchID     string
	Success     int
	Failure     int
	ErrorTokens []string
}

// List allowed topics
const (
	AllowedTopicAll     = "all"
	AllowedTopicIOS     = "iOS"
	AllowedTopicAndroid = "android"
)

// Service ...
type Service struct {
	Config
	Client *messaging.Client
}

// NewInstance for push notification
func NewInstance(cfg Config) (*Service, error) {
	ctx := context.Background()

	// Setup
	credential := base64Decode(cfg.Credential)
	opts := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: cfg.ProjectID,
	}, opts)
	if err != nil {
		return nil, errors.New("error when init Firebase app")
	}

	// Init db
	if cfg.PostgreSQL != nil {
		if err = cfg.PostgreSQL.AutoMigrate(
			&Log{tableName: cfg.LogTableName},
		); err != nil {
			return nil, errors.New("error when create new fcm log table for logging")
		}

		// TODO: add index for db (field action, created_at)
		//
	}

	// Init messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, errors.New("error when init Firebase messaging client")
	}

	s := Service{
		Config: cfg,
		Client: client,
	}
	return &s, nil
}

// base64Decode ...
func base64Decode(text string) []byte {
	sDec, _ := b64.StdEncoding.DecodeString(text)
	return sDec
}

// isTopicAllowed ...
func isTopicAllowed(topic string) bool {
	return funk.ContainsString(allowedTopics, topic)
}
