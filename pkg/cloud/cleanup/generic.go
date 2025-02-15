/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

package cleanup

import (
	"github.com/submariner-io/admiral/pkg/reporter"
	"github.com/submariner-io/cloud-prepare/pkg/api"
	"github.com/submariner-io/subctl/pkg/cloud/generic"
	"github.com/submariner-io/subctl/pkg/cluster"
)

func GenericCluster(clusterInfo *cluster.Info, status reporter.Interface) error {
	defer status.End()
	err := generic.RunOnCluster(clusterInfo, status,
		func(gwDeployer api.GatewayDeployer, status reporter.Interface) error {
			return gwDeployer.Cleanup(status) //nolint:wrapcheck // No need to wrap here
		})

	return status.Error(err, "Failed to cleanup generic K8s cluster")
}
