package interest

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var InterestRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/{tenant.Tenant}/interest/{interest.Interest}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/{tenant.Tenant}/interests",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:InterestRestAPISpec
type Interest struct {
	nexus.Node

	Name string
}
