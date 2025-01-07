package services

import (
	"notification-service-go/internal/models"
)

type NotificationService struct {
	EmailService EmailService
	PushService  PushService
}

func (ns *NotificationService) Notify(event models.Event) error {
	return nil
}
