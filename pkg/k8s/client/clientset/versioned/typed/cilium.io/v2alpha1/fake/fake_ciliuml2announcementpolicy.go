// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	ciliumiov2alpha1 "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/typed/cilium.io/v2alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCiliumL2AnnouncementPolicies implements CiliumL2AnnouncementPolicyInterface
type fakeCiliumL2AnnouncementPolicies struct {
	*gentype.FakeClientWithList[*v2alpha1.CiliumL2AnnouncementPolicy, *v2alpha1.CiliumL2AnnouncementPolicyList]
	Fake *FakeCiliumV2alpha1
}

func newFakeCiliumL2AnnouncementPolicies(fake *FakeCiliumV2alpha1) ciliumiov2alpha1.CiliumL2AnnouncementPolicyInterface {
	return &fakeCiliumL2AnnouncementPolicies{
		gentype.NewFakeClientWithList[*v2alpha1.CiliumL2AnnouncementPolicy, *v2alpha1.CiliumL2AnnouncementPolicyList](
			fake.Fake,
			"",
			v2alpha1.SchemeGroupVersion.WithResource("ciliuml2announcementpolicies"),
			v2alpha1.SchemeGroupVersion.WithKind("CiliumL2AnnouncementPolicy"),
			func() *v2alpha1.CiliumL2AnnouncementPolicy { return &v2alpha1.CiliumL2AnnouncementPolicy{} },
			func() *v2alpha1.CiliumL2AnnouncementPolicyList { return &v2alpha1.CiliumL2AnnouncementPolicyList{} },
			func(dst, src *v2alpha1.CiliumL2AnnouncementPolicyList) { dst.ListMeta = src.ListMeta },
			func(list *v2alpha1.CiliumL2AnnouncementPolicyList) []*v2alpha1.CiliumL2AnnouncementPolicy {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v2alpha1.CiliumL2AnnouncementPolicyList, items []*v2alpha1.CiliumL2AnnouncementPolicy) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
