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

package shell

import (
	"yunion.io/x/pkg/util/shellutils"

	"yunion.io/x/cloudmux/pkg/multicloud/google"
)

func init() {
	type RegionListOptions struct {
	}
	shellutils.R(&RegionListOptions{}, "region-list", "List regions", func(cli *google.SRegion, args *RegionListOptions) error {
		regions := cli.GetClient().GetRegions()
		printList(regions, 0, 0, 0, nil)
		return nil
	})

	type ResourceShowOptions struct {
		ID string
	}

	shellutils.R(&ResourceShowOptions{}, "resource-show", "show resource", func(cli *google.SRegion, args *ResourceShowOptions) error {
		ret := struct{}{}
		return cli.GetBySelfId(args.ID, &ret)
	})

	shellutils.R(&ResourceShowOptions{}, "resource-post", "post resource", func(cli *google.SRegion, args *ResourceShowOptions) error {
		ret := struct{}{}
		return cli.PostBySelfId(args.ID, &ret)
	})

}
