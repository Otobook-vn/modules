package rabbitmqpub

import (
	"github.com/Otobook-vn/modules/queue"
	queuemodel "github.com/Otobook-vn/modules/queue/model"
	"github.com/Otobook-vn/modules/rabbitmq"
)

// NotificationAdminPushAll ...
func NotificationAdminPushAll(data queuemodel.NotificationAdminPushAll) {
	queueName := queue.TaskNotificationAdminPushAll

	// Publish message
	if err := rabbitmq.Publish(queueName, data); err != nil {
		rabbitmq.PublishErrorLog(queueName, err)
	}
}
