// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudevent

import (
	"yunion.io/x/onecloud/pkg/mcclient"
	modules "yunion.io/x/onecloud/pkg/mcclient/modules/cloudevent"
	"yunion.io/x/onecloud/pkg/mcclient/options"
)

func init() {

	type CloudeventListOptions struct {
		options.BaseListOptions

		Since string `help:"since time"`
		Until string `hlep:"until time"`
	}
	R(&CloudeventListOptions{}, "cloud-event-list", "List cloud events", func(s *mcclient.ClientSession, opts *CloudeventListOptions) error {
		params, err := options.ListStructToParams(opts)
		if err != nil {
			return err
		}
		result, err := modules.Cloudevents.List(s, params)
		if err != nil {
			return err
		}
		printList(result, modules.Cloudevents.GetColumns(s))
		return nil
	})

	type CloudeventLogsPurgeOptions struct {
	}
	R(&CloudeventLogsPurgeOptions{}, "cloud-event-purge", "Purge obsolete cloud event logs", func(s *mcclient.ClientSession, opts *CloudeventLogsPurgeOptions) error {
		_, err := modules.Cloudevents.PerformClassAction(s, "purge-splitable", nil)
		if err != nil {
			return err
		}
		_, err = modules.CloudeventLogs.PerformClassAction(s, "purge-splitable", nil)
		if err != nil {
			return err
		}
		return nil
	})
}
