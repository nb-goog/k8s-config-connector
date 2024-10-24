// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kms

import (
	//refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KeyHandle_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KeyHandle {
	if in == nil {
		return nil
	}
	out := &krm.KeyHandle{}
	out.Name = direct.LazyPtr(in.GetName())
	out.KmsKey = direct.LazyPtr(in.GetKmsKey())
	out.ResourceTypeSelector = direct.LazyPtr(in.GetResourceTypeSelector())
	return out
}
func KeyHandle_ToProto(mapCtx *direct.MapContext, in *krm.KeyHandle) *pb.KeyHandle {
	if in == nil {
		return nil
	}
	out := &pb.KeyHandle{}
	out.Name = direct.ValueOf(in.Name)
	out.KmsKey = direct.ValueOf(in.KmsKey)
	out.ResourceTypeSelector = direct.ValueOf(in.ResourceTypeSelector)
	return out
}
