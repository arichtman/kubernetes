/*
Copyright 2017 The Kubernetes Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackPatch represents a desire to update Kubernetes resources with data from the world
type BackPatch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackPatchSpec   `json:"spec"`
	Status BackPatchStatus `json:"status"`
}

// Contains desired backpatch and parameters to achieve it
type BackPatchSpec struct {
	// +kubebuilder:validation:MinLength=1
	ConfigMapName string `json:"configMapName"`
}

// Contains information on patching status
// +kubebuilder:subresource:status
type BackPatchStatus struct {
	LastPatchTimestamp string `json:"lastPatchTimestamp"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type BackPatchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []BackPatch `json:"items"`
}
