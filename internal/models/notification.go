package models

type Event struct {
	EventType            string
	NotificationChannels []string
	TargetUser           string
	Message              string
}

const (
	EventChatMessage  = "chat_message"
	EventPurchase     = "purchase"
	EventPendingOrder = "pending_order"
	EventItemShipped  = "item_shipped"
)
