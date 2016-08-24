package activities

import (
	"github.com/naveego/pipeline-activities/transforms"
	"github.com/naveego/pipeline-api/activity"
	"github.com/naveego/pipeline-api/types"
)

func init() {

	activity.RegisterActivityFactory("map", func() types.Activity { return &transforms.MapActivity{} })
	activity.RegisterActivityFactory("passthru", func() types.Activity { return &PassthroughActivity{} })

}
