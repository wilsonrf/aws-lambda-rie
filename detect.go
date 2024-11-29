/*
 * Copyright 2012-2024 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
