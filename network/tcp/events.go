package tcp

import "github.com/finderAUT/hive.go/v3/events"

type tcpServerEvents struct {
	Start    *events.Event
	Shutdown *events.Event
	Connect  *events.Event
	Error    *events.Event
}
