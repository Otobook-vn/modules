package fcm

import (
	"context"
	"errors"
	"fmt"
)

// UnsubscribeTokensFromTopic ...
func UnsubscribeTokensFromTopic(topic string, tokens []string) error {
	if topic == "" || len(tokens) == 0 {
		return errors.New("invalid data")
	}

	if !isTopicAllowed(topic) {
		return errors.New("invalid topic")
	}

	ctx := context.Background()
	resp, err := client.UnsubscribeFromTopic(ctx, tokens, topic)
	if err != nil {
		fmt.Printf("*** Unsubscribe tokens from topic %s error: %s \n", topic, err.Error())
		return err
	}

	fmt.Printf("Unsubscribe tokens from topic %s: success %d, failed %d \n", topic, resp.SuccessCount, resp.FailureCount)
	if len(resp.Errors) > 0 {
		return errors.New(resp.Errors[0].Reason)
	}
	return nil
}
