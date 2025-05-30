/*
This file is part of the KubeVirt project

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Copyright The KubeVirt Authors.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1alpha1 "kubevirt.io/api/migrations/v1alpha1"
	scheme "kubevirt.io/client-go/kubevirt/scheme"
)

// MigrationPoliciesGetter has a method to return a MigrationPolicyInterface.
// A group's client should implement this interface.
type MigrationPoliciesGetter interface {
	MigrationPolicies() MigrationPolicyInterface
}

// MigrationPolicyInterface has methods to work with MigrationPolicy resources.
type MigrationPolicyInterface interface {
	Create(ctx context.Context, migrationPolicy *v1alpha1.MigrationPolicy, opts v1.CreateOptions) (*v1alpha1.MigrationPolicy, error)
	Update(ctx context.Context, migrationPolicy *v1alpha1.MigrationPolicy, opts v1.UpdateOptions) (*v1alpha1.MigrationPolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, migrationPolicy *v1alpha1.MigrationPolicy, opts v1.UpdateOptions) (*v1alpha1.MigrationPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MigrationPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MigrationPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MigrationPolicy, err error)
	MigrationPolicyExpansion
}

// migrationPolicies implements MigrationPolicyInterface
type migrationPolicies struct {
	*gentype.ClientWithList[*v1alpha1.MigrationPolicy, *v1alpha1.MigrationPolicyList]
}

// newMigrationPolicies returns a MigrationPolicies
func newMigrationPolicies(c *MigrationsV1alpha1Client) *migrationPolicies {
	return &migrationPolicies{
		gentype.NewClientWithList[*v1alpha1.MigrationPolicy, *v1alpha1.MigrationPolicyList](
			"migrationpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.MigrationPolicy { return &v1alpha1.MigrationPolicy{} },
			func() *v1alpha1.MigrationPolicyList { return &v1alpha1.MigrationPolicyList{} }),
	}
}
