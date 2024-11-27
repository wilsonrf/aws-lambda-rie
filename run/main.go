package main

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	awslambdarie "github.com/wilsonrf/aws-lambda-rie-buildpack"
)

func main() {
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	packit.Run(awslambdarie.Detect(), awslambdarie.Build(logger))
}
