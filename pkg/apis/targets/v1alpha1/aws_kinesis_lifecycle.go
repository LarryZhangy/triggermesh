/*
Copyright 2022 TriggerMesh Inc.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/triggermesh/triggermesh/pkg/apis/common/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/reconciler/resource"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (*AWSKinesisTarget) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AWSKinesisTarget")
}

// GetConditionSet implements duckv1.KRShaped.
func (*AWSKinesisTarget) GetConditionSet() apis.ConditionSet {
	return v1alpha1.DefaultConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (t *AWSKinesisTarget) GetStatus() *duckv1.Status {
	return &t.Status.Status
}

// GetStatusManager implements Reconcilable.
func (t *AWSKinesisTarget) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: t.GetConditionSet(),
		Status:       &t.Status,
	}
}

// GetAdapterOverrides implements AdapterConfigurable.
func (t *AWSKinesisTarget) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return t.Spec.AdapterOverrides
}

// WantsOwnServiceAccount implements ServiceAccountProvider.
func (t *AWSKinesisTarget) WantsOwnServiceAccount() bool {
	return t.Spec.Auth.EksIAMRole != nil
}

// ServiceAccountOptions implements ServiceAccountProvider.
func (t *AWSKinesisTarget) ServiceAccountOptions() []resource.ServiceAccountOption {
	var saOpts []resource.ServiceAccountOption

	if iamRole := t.Spec.Auth.EksIAMRole; iamRole != nil {
		saOpts = append(saOpts, v1alpha1.AwsIamRoleAnnotation(*iamRole))
	}

	return saOpts
}
