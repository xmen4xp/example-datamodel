package wanna

import (
	"example/tenant/interest"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var WannaRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/${tenant.Tenant}/user/${user.User}/wanna/{wanna.Wanna}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/${tenant.Tenant}/user/${user.User}/wannas",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:WannaRestAPISpec
type Wanna struct {
	nexus.Node

	Name string

	Interest interest.Interest `nexus:"link"`
}
