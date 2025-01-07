package services

type PushService struct {
	ProviderA string
	ProviderB string
}

func (ps *PushService) Send(target string, message string) error {
	return nil
}
