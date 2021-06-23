package rabbitmqpub

import (
	"github.com/Otobook-vn/modules/queue"
	queuemodel "github.com/Otobook-vn/modules/queue/model"
	"github.com/Otobook-vn/modules/rabbitmq"
)

// CreateAndPushSingleNotification ...
func CreateAndPushSingleNotification(data queuemodel.NotificationCreateAndPushSingle) {
	queueName := queue.TaskNotificationCreateAndPushSingle

	// Publish message
	if err := rabbitmq.Publish(queueName, data); err != nil {
		rabbitmq.PublishErrorLog(queueName, err)
	}
}
