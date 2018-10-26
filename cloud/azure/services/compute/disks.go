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
package compute

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
)

func (s *Service) DeleteManagedDisk(resourceGroup string, name string) (compute.DisksDeleteFuture, error) {
	return s.DisksClient.Delete(s.ctx, resourceGroup, name)
}

func (s *Service) WaitForDisksDeleteFuture(future compute.DisksDeleteFuture) error {
	return future.Future.WaitForCompletionRef(s.ctx, s.DisksClient.Client)
}
