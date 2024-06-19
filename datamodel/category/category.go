package category

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/common-library/pkg/nexus"
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type Category struct {
	nexus.Node

	Desription string
}
