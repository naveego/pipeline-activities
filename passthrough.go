package activities

import (
	"github.com/Sirupsen/logrus"
	"github.com/naveego/naveego-api/types/pipeline"
	"github.com/naveego/pipeline-api/activity"
)

type PassthroughActivity struct {
}

func (p *PassthroughActivity) Execute(context activity.Context, dataPoint pipeline.DataPoint) error {
	logrus.Debugf("Passing through message")
	context.OutputCollector.Emit(dataPoint)
	return nil
}
