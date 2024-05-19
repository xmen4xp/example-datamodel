package evaluation

import (
	"example/evaluation/quiz"

	"github.com/vmware-tanzu/graph-framework-for-microservices/common-library/pkg/nexus"
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type Evaluation struct {
	nexus.SingletonNode

	Quiz quiz.Quiz `nexus:"children"`
}
