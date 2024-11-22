package lambda

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
)

type Detect struct{}

const (
	PlanEntryAwsLambda             = "aws-lambda"
	PlanEntryCustomRuntimeEmulator = "aws-custom-runtime-emulator-extension"
)

func (d *Detect) Detect(context packit.DetectContext) (packit.DetectResult, error) {
	result := packit.DetectResult{
		Plan: packit.BuildPlan{
			Provides: []packit.BuildPlanProvision{
				{Name: PlanEntryAwsLambda},
			},
			Requires: []packit.BuildPlanRequirement{
				{Name: PlanEntryAwsLambda},
			},
		},
	}
	if env, ok := os.LookupEnv("BP_AWS_RIE"); ok {
		if env == "true" {
			result.Plan.Requires = append(result.Plan.Requires, packit.BuildPlanRequirement{Name: PlanEntryCustomRuntimeEmulator})
		}
	}

	return result, nil
}
