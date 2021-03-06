package fcm

import (
	"context"
	"errors"
	"fmt"

	"github.com/Otobook-vn/modules/utils"
)

// UnsubscribeTokensFromTopic ...
func (s Service) UnsubscribeTokensFromTopic(batchID, topic string, tokens []string) error {
	if topic == "" || len(tokens) == 0 {
		return errors.New("invalid data")
	}

	if !isTopicAllowed(topic) {
		return errors.New("invalid topic")
	}

	ctx := context.Background()
	resp, err := s.Client.UnsubscribeFromTopic(ctx, tokens, topic)
	if err != nil {
		fmt.Printf("*** Unsubscribe tokens from topic %s error: %s \n", topic, err.Error())
		return err
	}

	fmt.Printf("Unsubscribe tokens from topic %s: success %d, failed %d \n", topic, resp.SuccessCount, resp.FailureCount)
	if len(resp.Errors) > 0 {
		return errors.New(resp.Errors[0].Reason)
	}

	// Save log
	go func() {
		if batchID != "" {
			log := Log{
				ID:           utils.GenerateUUID(),
				Action:       LogActionUnsubscribeTokens,
				BatchID:      batchID,
				Topics:       utils.ConvertToDataTypesJSON([]string{topic}),
				TokenCount:   len(tokens),
				SuccessCount: 0,
				FailureCount: 0,
				CreatedAt:    utils.TimeNow(),
			}
			s.saveLog(log)
		}
	}()

	return nil
}
