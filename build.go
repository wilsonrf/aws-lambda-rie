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
