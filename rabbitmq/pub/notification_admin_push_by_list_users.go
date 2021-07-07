package rabbitmqpub

import (
	"github.com/Otobook-vn/modules/queue"
	queuemodel "github.com/Otobook-vn/modules/queue/model"
	"github.com/Otobook-vn/modules/rabbitmq"
)

// NotificationAdminPushByListUsers ...
func NotificationAdminPushByListUsers(data queuemodel.NotificationAdminPushByListUsers) {
	queueName := queue.TaskNotificationAdminPushByListUsers

	// Publish message
	if err := rabbitmq.Publish(queueName, data); err != nil {
		rabbitmq.PublishErrorLog(queueName, err)
	}
}
