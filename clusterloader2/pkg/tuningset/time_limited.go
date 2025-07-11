/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tuningset

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/perf-tests/clusterloader2/api"
)

type timeLimitedLoad struct {
	params *api.TimeLimitedLoad
}

func newTimeLimitedLoad(params *api.TimeLimitedLoad) TuningSet {
	return &timeLimitedLoad{
		params: params,
	}
}

func (t *timeLimitedLoad) Execute(actions []func()) {
	if len(actions) == 0 {
		return
	}
	sleepDuration := time.Duration(t.params.TimeLimit.ToTimeDuration().Nanoseconds() / int64(len(actions)))
	var wg wait.Group
	for i := range actions {
		wg.Start(actions[i])
		time.Sleep(sleepDuration)
	}
	wg.Wait()
}
