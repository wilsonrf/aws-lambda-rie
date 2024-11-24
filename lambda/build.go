package lambda

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
)

type Build struct{}

func (b *Build) Build(context packit.BuildContext) (packit.BuildResult, error) {

	result := packit.BuildResult{}

	if env, ok := os.LookupEnv("BP_AWS_RIE"); ok {
		if env == "true" {
			layer, err := context.Layers.Get("emulator")
			if err != nil {
				return result, err
			}
			layer.Reset()
			layer.Launch = true
			layer.Cache = true
			layer.Build = true

			p := packit.DirectProcess{
				Type:    "aws-lambda-rie",
				Command: []string{"/home/cnb/.aws-lambda-rie/aws-lambda-rie"},
				Default: false,
			}
			result.Layers = append(result.Layers, layer)
			result.Launch.DirectProcesses = append(result.Launch.DirectProcesses, p)
		}
	}

	return result, nil
}
