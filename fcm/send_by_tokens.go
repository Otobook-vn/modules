package fcm

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/messaging"
	"github.com/thoas/go-funk"
)

const maxTokensPerSend = 1000

// SendByTokens ...
func SendByTokens(tokens []string, batchID string, payload messaging.Message) (result Result, err error) {
	ctx := context.Background()
	result.BatchID = batchID

	// Send
	for {
		listTokens, restTokens := separateTokens(tokens)

		// Return if list tokens length < 0
		if len(listTokens) <= 0 {
			break
		}

		// Prepare message
		message := &messaging.MulticastMessage{
			Tokens:       listTokens,
			Data:         payload.Data,
			Notification: payload.Notification,
			Android:      payload.Android,
		}

		// Send
		resp, err := client.SendMulticast(ctx, message)
		if err != nil {
			fmt.Println(fmt.Printf("Error when push notification with sendID %s, error: %s", batchID, err.Error()))
			return
		}

		result.Success += resp.SuccessCount
		result.Failure += resp.FailureCount

		// Assign token with rest tokens
		tokens = restTokens
	}

	return Result{}, nil
}

// separate tokens for multiple times send, due to FCM limited
func separateTokens(tokens []string) (list, rest []string) {
	list = tokens
	if len(tokens) > maxTokensPerSend {
		list = removeEmptyToken(tokens[:maxTokensPerSend])
		rest = tokens[maxTokensPerSend:]
	}
	return
}

// remove empty token
func removeEmptyToken(tokens []string) []string {
	result := funk.FilterString(tokens, func(t string) bool {
		return t != ""
	})
	return result
}
