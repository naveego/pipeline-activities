package activities

import (
	"github.com/Sirupsen/logrus"
	"github.com/naveego/pipeline-api/types"
)

type PassthroughActivity struct {
}

func (p *PassthroughActivity) Execute(context types.ActivityContext, dataPoint types.DataPoint) error {
	logrus.Debugf("Passing through message")
	context.OutputCollector.Emit(dataPoint)
	return nil
}
