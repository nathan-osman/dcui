package server

import (
	"sync"
)

const (
	alertInfo   = "info"
	alertDanger = "danger"
)

// alert represents an individual alert message.
type alert struct {
	Type string
	Body string
}

var (
	alertMutex sync.Mutex
	alerts     []*alert
)

// addAlert adds a new alert.
func addAlert(type_, body string) {
	alertMutex.Lock()
	defer alertMutex.Unlock()
	alerts = append(alerts, &alert{
		Type: type_,
		Body: body,
	})
}

// getAlerts retrieves all current alerts.
func getAlerts() []*alert {
	alertMutex.Lock()
	defer alertMutex.Unlock()
	ret := alerts
	alerts = []*alert{}
	return ret
}
