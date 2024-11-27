package awslambdarie

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func Build(logger scribe.Emitter) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {

		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		result := packit.BuildResult{}

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

		return result, nil
	}
}
