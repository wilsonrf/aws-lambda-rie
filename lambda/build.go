package lambda

import "github.com/paketo-buildpacks/packit/v2"

type Build struct{}

func (b *Build) Build(context packit.BuildContext) (packit.BuildResult, error) {
	return packit.BuildResult{}, nil
}
