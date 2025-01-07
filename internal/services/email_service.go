package services

type EmailService struct {
	ProviderA string
	ProviderB string
}

func (es *EmailService) Send(to string, message string) error {
	return nil
}
