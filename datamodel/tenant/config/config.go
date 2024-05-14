package config

import (
	"example/tenant/config/event"
	"example/tenant/config/user"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type Config struct {
	nexus.SingletonNode

	User  user.User   `nexus:"children"`
	Event event.Event `nexus:"children"`
}
