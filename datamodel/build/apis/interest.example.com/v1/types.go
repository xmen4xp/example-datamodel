// Code generated by nexus. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"example/build/common"
)

// +k8s:openapi-gen=true
type Child struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

// +k8s:openapi-gen=true
type Link struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

// +k8s:openapi-gen=true
type SyncerStatus struct {
	EtcdVersion    int64 `json:"etcdVersion, omitempty" yaml:"etcdVersion, omitempty"`
	CRGenerationId int64 `json:"cRGenerationId, omitempty" yaml:"cRGenerationId, omitempty"`
}

// +k8s:openapi-gen=true
type NexusStatus struct {
	SourceGeneration int64        `json:"sourceGeneration, omitempty" yaml:"sourceGeneration, omitempty"`
	RemoteGeneration int64        `json:"remoteGeneration, omitempty" yaml:"remoteGeneration, omitempty"`
	SyncerStatus     SyncerStatus `json:"syncerStatus, omitempty" yaml:"syncerStatus, omitempty"`
}

/* ------------------- CRDs definitions ------------------- */

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Interest struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              InterestSpec        `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status            InterestNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type InterestNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *Interest) CRDName() string {
	return "interests.interest.example.com"
}

func (c *Interest) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type InterestSpec struct {
	Name string `json:"name" yaml:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type InterestList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Interest `json:"items" yaml:"items"`
}
