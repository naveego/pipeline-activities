package activities

import "github.com/naveego/pipeline-api/types"

type PassthroughActivity struct {
}

func (p *PassthroughActivity) Execute(context types.ActivityContext, dataPoint types.DataPoint) error {
	context.OutputCollector.Emit(dataPoint)
	return nil
}
