package fcm

import (
	"context"
	"fmt"
	"strings"

	"firebase.google.com/go/v4/messaging"
	"github.com/thoas/go-funk"
)

// List allowed topics
const (
	AllowedTopicAll     = "all"
	AllowedTopicIOS     = "iOS"
	AllowedTopicAndroid = "android"
)

var allowedTopics = []string{AllowedTopicAll, AllowedTopicIOS, AllowedTopicAndroid}

// SendByTopics ...
func SendByTopics(topics []string, batchID string, payload messaging.Message) {
	ctx := context.Background()

	// Return if have no topics
	if len(topics) == 0 {
		fmt.Println(fmt.Sprintf("*** Empty topics array with batch id %s", batchID))
		return
	}

	// Get topic condition
	payload.Condition = getTopicCondition(topics)

	// Return if there is no condition
	if payload.Condition == "" {
		fmt.Println(fmt.Sprintf("*** No valid topics array with batch id %s: %v", batchID, topics))
		return
	}

	_, err := client.Send(ctx, &payload)
	if err != nil {
		fmt.Println(fmt.Sprintf("*** Send topic error with batch id %s: %s", batchID, err.Error()))
		fmt.Println("*** Topics:", topics)
	}
}

// getTopicCondition ...
func getTopicCondition(topics []string) string {
	var topicCond []string

	for _, topic := range topics {
		if !isTopicAllowed(topic) {
			continue
		}
		cond := fmt.Sprintf("'%s' in topics", topic)
		topicCond = append(topicCond, cond)
	}

	if len(topicCond) == 0 {
		return ""
	}
	return strings.Join(topicCond, " && ")
}

// isTopicAllowed ...
func isTopicAllowed(topic string) bool {
	return funk.ContainsString(allowedTopics, topic)
}
