package awslambdarie

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

const (
	PlanEntryAwsLambda             = "aws-lambda"
	PlanEntryCustomRuntimeEmulator = "aws-custom-runtime-emulator-extension"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		logger := scribe.NewLogger(os.Stdout)
		if env, ok := os.LookupEnv("BP_AWS_RIE"); ok {
			logger.Detail("BP_AWS_RIE: %v", env)
			if env == "true" {
				logger.Detail("BP_AWS_RIE is true")
				return packit.DetectResult{
					Plan: packit.BuildPlan{
						Provides: []packit.BuildPlanProvision{
							{Name: PlanEntryAwsLambda},
						},
						Requires: []packit.BuildPlanRequirement{
							{Name: PlanEntryAwsLambda},
							{Name: PlanEntryCustomRuntimeEmulator},
						},
					},
				}, nil
			} else {
				return packit.DetectResult{}, nil
			}
		}

		return packit.DetectResult{}, nil
	}
}
