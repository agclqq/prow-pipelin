package register

import (
	"github.com/agclqq/prow-framework/event"

	"github.com/agclqq/prow-pipeline/app/events"
)

type Demo struct {
}

func Register() {
	event.Register(&events.Demo{})
}
