package transforms

import (
	"github.com/Sirupsen/logrus"
	"github.com/naveego/pipeline-api/types"
)

type MapActivity struct {
}

func (m *MapActivity) Execute(context types.ActivityContext, dataPoint types.DataPoint) error {
	logrus.Info("Hello from map")
	return nil
}
