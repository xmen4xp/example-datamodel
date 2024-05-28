package runtimequiz

import (
	"example/evaluation/quiz"
	"example/tenant/runtime/runtimeuser/runtimeevaluation/runtimequiz/runtimeanswer"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var RuntimeQuizRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/{tenant.Tenant}/runtime/user/${runtimeuser.RuntimeUser}/quiz/{runtimequiz.RuntimeQuiz}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/{tenant.Tenant}/runtime/user/${runtimeuser.RuntimeUser}/quizes",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:RuntimeQuizRestAPISpec
type RuntimeQuiz struct {
	nexus.Node
	Quiz    quiz.Quiz                   `nexus:"link"`
	Answers runtimeanswer.RuntimeAnswer `nexus:"children"`
	Status  RuntimeQuizStatus           `nexus:"status"`
}

type RuntimeQuizStatus struct {
	TotalScore int
}
