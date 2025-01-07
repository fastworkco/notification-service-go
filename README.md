# Software Engineer Challenge @ Fastwork

You are going to implement a notification service that will send notifications to users through different channels.

For example, on an e-commerce platform, when a buyer purchases a product, the notification service should send both email and push notifications to the seller.


![Untitled-2024-06-19-1904 (2)](https://gist.github.com/assets/9070844/8c9c0304-0443-4a9b-b159-634cc3c0d5c3)


## Target events
| Event                                 | Noti channels     | Noti target   |
| --------                              | --------          | --------      |
| Buyer sends chat message to seller    | Email, Push       | Seller        |
| Buyer purchases product               | Email, Push       | Seller        |
| Remind to pay for pending order       | Push              | Buyer         |
| Purchased items are shipped           | Push              | Buyer         |


## Requirements
- The service should be able to handle huge traffic spikes, e.g., sale events on double day 11.11.
- Notification channels should be configurable, e.g., the system should be able to send notifications via email, push, or both.
- Notification channel providers should be configurable and able to automatically switch over when one is down.
    - For example, if we have providers A and B for email, the system should automatically switch to B when A is down.
- Feel free to add any additional features you think might be necessary.
- The team should be able to monitor and verify if notifications are sent successfully.
- The final outcome, if not runnable, should at least include unit tests that cover most scenarios.


## Technologies Stack
**Programming Language**
- Golang

**Framework**
- Your choice

**Database**
- PostgreSQL, MySQL

## Example Project structure
```
notification-service-go/
├── cmd/
│   └── main.go
├── config/
│   └── config.yaml
├── internal/
│   ├── handlers/
│   │   └── notification.go
│   ├── models/
│   │   └── notification.go
│   ├── services/
│   │   ├── email_service.go
│   │   ├── push_service.go
│   │   └── notification_service.go
├── tests/
│   └── notification_test.go
├── config.yaml
├── docker-compose.yaml
├── go.mod
└── go.sum
```
