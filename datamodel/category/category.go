package category

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/common-library/pkg/nexus"
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var CategoryRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/category/{category.Category}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/categories",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:CategoryRestAPISpec
type Category struct {
	nexus.Node

	Desription string
}
