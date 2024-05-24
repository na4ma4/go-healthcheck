package healthcheck

type Report interface {
	GetStatus() string
}
