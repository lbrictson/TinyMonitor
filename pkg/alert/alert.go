package alert

type Alert interface {
	SendDown(monitorName string, message string) error
	SendUp(monitorName string, message string) error
}
