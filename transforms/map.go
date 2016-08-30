package transforms

import (
	"github.com/Sirupsen/logrus"
	"github.com/naveego/naveego-api/types/pipeline"
	"github.com/naveego/pipeline-api/activity"
)

type MapActivity struct {
}

func (m *MapActivity) Execute(context activity.Context, dataPoint pipeline.DataPoint) error {
	logrus.Info("Hello from map")
	return nil
}
