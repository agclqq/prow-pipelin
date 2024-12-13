package register

import (
	"github.com/agclqq/prowjob"

	"github.com/agclqq/prow-pipeline/app/console/command"
)

func Register(eng *prowjob.CommandEngine) {
	eng.Add(&command.Demo{})
}
