package brokers

import (
	"TasQ/app/tasks"
)

// broker is a queue for saving tasQ
type Broker interface {
	Acquire(string) *tasks.TasQ // acquire get a tasQ from broker
	Enqueue(*tasks.TasQ) string // enqueue put a tasQ to broker
	Update(*tasks.TasQ)         // update change tasQ status
	Cancel(*tasks.TasQ)         // cancel a tasQ
}
