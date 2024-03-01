package user

import (
	"example/tenant/config/user/wanna"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var UserRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/${tenant.Tenant}/user/${user.User}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/${tenant.Tenant}/users",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:UserRestAPISpec
type User struct {
	nexus.Node

	Username  string `json:"username" yaml:"username"`
	Mail      string `json:"email,omitempty" yaml:"mail,omitempty"`
	FirstName string `json:"firstName,omitempty" yaml:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" yaml:"lastName,omitempty"`
	Password  string `json:"password" yaml:"password"`
	Realm     string `json:"realm,omitempty" yaml:"realm,omitempty"`

	Wanna wanna.Wanna `nexus:"children"`
}
