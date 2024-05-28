package runtime

import (
	"example/tenant/runtime/runtimeuser"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type Runtime struct {
	nexus.SingletonNode

	User runtimeuser.RuntimeUser `nexus:"children"`
}
