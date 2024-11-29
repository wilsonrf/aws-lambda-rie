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

package awslambdarie_test

import (
	"os"
	"testing"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/sclevine/spec"
	awslambdarie "github.com/wilsonrf/aws-lambda-rie-buildpack"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
		detect     packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = os.MkdirTemp("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		detect = awslambdarie.Detect()
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	it("should detect if BP_AWS_RIE is true", func() {
		Expect(os.Setenv("BP_AWS_RIE", "true")).To(Succeed())

		result, err := detect(packit.DetectContext{
			WorkingDir: workingDir,
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(result).To(Equal(packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: awslambdarie.PlanEntryAwsLambda},
				},
				Requires: []packit.BuildPlanRequirement{
					{Name: awslambdarie.PlanEntryAwsLambda},
					{Name: awslambdarie.PlanEntryCustomRuntimeEmulator},
				},
			},
		}))
	})

	it("should detect false if BP_AWS_RIE is false", func() {
		Expect(os.Setenv("BP_AWS_RIE", "false")).To(Succeed())

		result, err := detect(packit.DetectContext{
			WorkingDir: workingDir,
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(result).To(Equal(packit.DetectResult{}))

	})

	it("should detect false if BP_AWS_RIE is not present", func() {
		Expect(os.Unsetenv("BP_AWS_RIE")).To(Succeed())

		result, err := detect(packit.DetectContext{
			WorkingDir: workingDir,
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(result).To(Equal(packit.DetectResult{}))
	})
}
