package activities

import (
	"github.com/naveego/pipeline-activities/transforms"
	"github.com/naveego/pipeline-api"
	"github.com/naveego/pipeline-api/types"
)

func init() {

	pipeline.RegisterActivityFactory("map", func() types.Activity { return &transforms.MapActivity{} })

}
