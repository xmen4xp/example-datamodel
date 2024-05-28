package runtimeevaluation

import (
	"example/tenant/runtime/runtimeuser/runtimeevaluation/runtimequiz"

	"github.com/vmware-tanzu/graph-framework-for-microservices/common-library/pkg/nexus"
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type RuntimeEvaluation struct {
	nexus.SingletonNode

	Quiz runtimequiz.RuntimeQuiz `nexus:"children"`
}
