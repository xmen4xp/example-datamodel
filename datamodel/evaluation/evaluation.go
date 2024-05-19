package evaluation

import (
	"example/evaluation/quiz"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type Evaluation struct {
	nexus.Node

	Quiz quiz.Quiz `nexus:"children"`
}
