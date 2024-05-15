package event

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var EventRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/{tenant.Tenant}/event/{event.Event}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/{tenant.Tenant}/events",
			Methods: nexus.HTTPListResponse,
		},
	},
}

type Time struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Min    int
	Second int
	Zone   string
}

// nexus-rest-api-gen:EventRestAPISpec
type Event struct {
	nexus.Node

	Description string
	MeetingLink string
	Time        Time
	Public      bool

	Status Status `nexus:"status"`
}

type Status struct {
	Time                string
	AssignedMeetingLink string
}
