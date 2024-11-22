package lambda

import "github.com/paketo-buildpacks/packit/v2"

type Detect struct{}

func (d *Detect) Detect(context packit.DetectContext) (packit.DetectResult, error) {
	return packit.DetectResult{}, nil
}
