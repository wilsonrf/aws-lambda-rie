package main

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/wilsonrf/aws-lambda-rie-buildpack/lambda"
)

func main() {
	detect := lambda.Detect{}
	build := lambda.Build{}
	packit.Run(detect.Detect, build.Build)
}
