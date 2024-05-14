package event

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var EventRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/${tenant.Tenant}/event/{event.Event}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/${tenant.Tenant}/events",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:EventRestAPISpec
type Event struct {
	nexus.Node

	Description string
	MeetingLink string
	DateTime    string
	Public      bool
}
