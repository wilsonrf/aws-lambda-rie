package lambda

import "github.com/paketo-buildpacks/packit/v2"

type Detect struct{}

const (
	PlanEntryAwsLambda             = "aws-lambda"
	PlanEntryCustomRuntimeEmulator = "aws-custom-runtime-emulator-extension"
)

func (d *Detect) Detect(context packit.DetectContext) (packit.DetectResult, error) {
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
}
