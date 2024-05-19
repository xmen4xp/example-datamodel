package root

import (
	evaluation "example/evaluation"
	"example/tenant"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

// Datamodel graph root
type Root struct {
	nexus.SingletonNode

	Tenant tenant.Tenant `nexus:"children"`

	Evaluation evaluation.Evaluation `nexus:"child"`

	// Soft links specified with annotatation: `nexus:"link"`
}
