package tenant

import (
	"example/tenant/config"
	"example/tenant/interest"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type Tenant struct {
	nexus.Node

	Config   config.Config     `nexus:"child"`
	Interest interest.Interest `nexus:"children"`
}
