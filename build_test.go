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
	"bytes"
	"os"
	"testing"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/sclevine/spec"
	awslambdarie "github.com/wilsonrf/aws-lambda-rie-buildpack"

	. "github.com/onsi/gomega"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		layersDir  string
		workingDir string
		cnbDir     string
		buffer     *bytes.Buffer
		logger     scribe.Emitter
		build      packit.BuildFunc
	)

	it.Before(func() {
		var err error
		buffer = bytes.NewBuffer(nil)
		logger = scribe.NewEmitter(buffer)
		layersDir, err = os.MkdirTemp("", "layers")
		Expect(err).NotTo(HaveOccurred())

		cnbDir, err = os.MkdirTemp("", "cnb")
		Expect(err).NotTo(HaveOccurred())

		workingDir, err = os.MkdirTemp("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		build = awslambdarie.Build(logger)

	})

	it.After(func() {
		Expect(os.RemoveAll(layersDir)).To(Succeed())
		Expect(os.RemoveAll(cnbDir)).To(Succeed())
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	it("should add a runtime interface emulator", func() {
		br, err := build(packit.BuildContext{
			WorkingDir: workingDir,
			Layers:     packit.Layers{Path: layersDir},
			CNBPath:    cnbDir,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(br.Layers[0]).To(HaveField("Name", "emulator"))
		Expect(br.Layers[0]).To(HaveField("Launch", true))
		Expect(br.Layers[0]).To(HaveField("Build", true))
		Expect(br.Layers[0]).To(HaveField("Cache", true))
		Expect(br.Launch.DirectProcesses[0]).To(HaveField("Type", "aws-lambda-rie"))
		Expect(br.Launch.DirectProcesses[0]).To(HaveField("Command", []string{"/home/cnb/.aws-lambda-rie/aws-lambda-rie"}))
		Expect(br.Launch.DirectProcesses[0]).To(HaveField("Default", false))
	})
}
