//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MetaOperatorSetSpec defines the desired state of MetaOperatorSet
// +k8s:openapi-gen=true
type MetaOperatorSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// +listType=set
	Services []SetService `json:"services"`
}

// SetService is the service configuration for a common service
type SetService struct {
	Name    string `json:"name"`
	Channel string `json:"channel,omitempty"`
	State   string `json:"state"`
}

// ConditionType is the condition of a service
type ConditionType string

// ClusterPhase is the phase of the installation
type ClusterPhase string

// Constants are used for state
const (
	ConditionInstall ConditionType = "Install"
	ConditionUpdate  ConditionType = "Update"
	ConditionDelete  ConditionType = "Delete"

	ClusterPhaseNone     ClusterPhase = ""
	ClusterPhaseCreating ClusterPhase = "Creating"
	ClusterPhaseUpdating ClusterPhase = "Updating"
	ClusterPhaseDeleting ClusterPhase = "Deleting"
	ClusterPhaseRunning  ClusterPhase = "Running"
	ClusterPhaseFailed   ClusterPhase = "Failed"
)

// Condition defines the current state of operator deploy
type Condition struct {
	Name           string        `json:"name,omitempty"`
	Type           ConditionType `json:"type,omitempty"`
	Status         string        `json:"status,omitempty"`
	LastUpdateTime string        `json:"lastUpdateTime,omitempty"`
	Reason         string        `json:"reason,omitempty"`
	Message        string        `json:"message,omitempty"`
}

// MetaOperatorSetStatus defines the observed state of MetaOperatorSet
// +k8s:openapi-gen=true
type MetaOperatorSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// Conditions represents the current state of the Set Service
	// +optional
	// +listType=set
	Conditions []Condition `json:"conditions,omitempty"`
	// Members represnets the current operators of the set
	// +optional
	Members MembersStatus `json:"members,omitempty"`
	// Phase is the cluster running phase
	Phase ClusterPhase `json:"phase"`
}

// MembersStatus show if the Operator is ready
type MembersStatus struct {
	// Ready are the operator members that are ready
	// The member names are the same as the subscription
	Ready []string `json:"ready,omitempty"`
	// Unready is that operator members are not ready
	Unready []string `json:"unready,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MetaOperatorSet is the Schema for the metaoperatorsets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=metaoperatorsets,shortName=moset,scope=Namespaced
type MetaOperatorSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetaOperatorSetSpec   `json:"spec,omitempty"`
	Status MetaOperatorSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MetaOperatorSetList contains a list of MetaOperatorSet
type MetaOperatorSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetaOperatorSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetaOperatorSet{}, &MetaOperatorSetList{})
}
