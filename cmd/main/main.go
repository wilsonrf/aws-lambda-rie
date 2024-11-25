package main

import (
	"github.com/paketo-buildpacks/packit/v2"
	awslambdarie "github.com/wilsonrf/aws-lambda-rie-buildpack/aws-lambda-rie"
)

func main() {
	detect := awslambdarie.Detect{}
	build := awslambdarie.Build{}
	packit.Run(detect.Detect, build.Build)
}
