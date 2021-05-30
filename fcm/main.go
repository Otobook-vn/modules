package fcm

import (
	"context"
	b64 "encoding/base64"
	"errors"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

// Config ...
type Config struct {
	// Project id
	ProjectID string

	// Original is JSON format, but encoded with base64, need to call base64 decode to get byte data
	Credential string
}

// Result of each send
type Result struct {
	BatchID     string
	Success     int
	Failure     int
	ErrorTokens []string
}

var client *messaging.Client

// Init ...
func Init(cfg Config) error {
	ctx := context.Background()

	// Setup
	credential := base64Decode(cfg.Credential)
	opts := option.WithCredentialsJSON(credential)
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: cfg.ProjectID,
	}, opts)
	if err != nil {
		return errors.New("error when init Firebase app")
	}

	// Init messaging client
	client, err = app.Messaging(ctx)
	if err != nil {
		return errors.New("error when init Firebase messaging client")
	}

	return nil
}

// base64Decode ...
func base64Decode(text string) []byte {
	sDec, _ := b64.StdEncoding.DecodeString(text)
	return sDec
}
