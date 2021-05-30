package fcm

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/messaging"
)

// SubscribeTokensToTopic ...
func SubscribeTokensToTopic(topic string, tokens []string) (result Result, err error) {
	if topic == "" || len(tokens) == 0 {
		return
	}

	if !isTopicAllowed(topic) {
		return
	}

	ctx := context.Background()
	for {
		listTokens, restTokens := separateTokens(tokens)

		// Return if list tokens length < 0
		if len(listTokens) <= 0 {
			break
		}

		resp, err := client.SubscribeToTopic(ctx, tokens, topic)
		if err != nil {
			fmt.Printf("*** Subscribe tokens to topic %s error: %s \n", topic, err.Error())
		}

		// Assign result
		result.Success += resp.SuccessCount
		result.Failure += resp.FailureCount

		// Get list error tokens
		if len(resp.Errors) > 0 {
			result.ErrorTokens = append(result.ErrorTokens, getFailureTokensFromSubscribe(resp.Errors, listTokens)...)
		}

		// Assign token with rest tokens
		tokens = restTokens
	}

	return
}

func getFailureTokensFromSubscribe(r []*messaging.ErrorInfo, inputTokens []string) (tokens []string) {
	for _, info := range r {
		tokens = append(tokens, inputTokens[info.Index])
	}
	return
}
