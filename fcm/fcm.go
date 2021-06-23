package fcm

import (
	"context"
	b64 "encoding/base64"
	"errors"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/thoas/go-funk"
	"google.golang.org/api/option"
	"gorm.io/gorm"

	"github.com/Otobook-vn/modules/postgresql"
)

// PGConfig ...
type PGConfig struct {
	Host, User, Password, DBName, Port string
}

// Config ...
type Config struct {
	// Project id
	ProjectID string
	// Original is JSON format, but encoded with base64, need to call base64 decode to get byte data
	Credential string
	// Database config
	PostgreSQL PGConfig
}

// Result of each send
type Result struct {
	BatchID     string
	Success     int
	Failure     int
	ErrorTokens []string
}

// List topics
const (
	TopicAll     = "all"
	TopicIOS     = "iOS"
	TopicAndroid = "android"
)

var allowedTopics = []string{TopicAll, TopicIOS, TopicAndroid}

// Service ...
type Service struct {
	Config
	Client *messaging.Client
	DB     *gorm.DB
}

// NewInstance for push notification
func NewInstance(config Config) (*Service, error) {
	if config.ProjectID == "" || config.Credential == "" || config.PostgreSQL.Host == "" {
		return nil, errors.New("please provide all information that needed: projectId, credential, postgresql")
	}

	ctx := context.Background()

	// Connect db first
	err := postgresql.Connect(
		config.PostgreSQL.Host,
		config.PostgreSQL.User,
		config.PostgreSQL.Password,
		config.PostgreSQL.DBName,
		config.PostgreSQL.Port,
	)
	if err != nil {
		fmt.Println("Cannot init module FCM", err)
		return nil, err
	}

	// Setup
	credential := base64Decode(config.Credential)
	opts := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: config.ProjectID,
	}, opts)
	if err != nil {
		return nil, errors.New("error when init Firebase app")
	}

	// Init messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, errors.New("error when init Firebase messaging client")
	}

	s := Service{
		Config: config,
		Client: client,
		DB:     postgresql.GetInstance(ctx),
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
