package activities

import (
	"github.com/naveego/pipeline-activities/transforms"
	"github.com/naveego/pipeline-api/activity"
)

func init() {

	activity.RegisterActivityFactory("map", func() activity.Activity { return &transforms.MapActivity{} })
	activity.RegisterActivityFactory("passthrough", func() activity.Activity { return &PassthroughActivity{} })

}
