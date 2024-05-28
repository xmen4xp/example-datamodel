package runtimeuser

import (
	"example/tenant/config/user"
	"example/tenant/runtime/runtimeuser/runtimeevaluation"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var RuntimeUserRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/{tenant.Tenant}/runtime/user/${runtimeuser.RuntimeUser}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/{tenant.Tenant}/runtime/users",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:RuntimeUserRestAPISpec
type RuntimeUser struct {
	nexus.Node

	Evaluation runtimeevaluation.RuntimeEvaluation `nexus:"child"`
	User       user.User                           `nexus:"link"`
}
