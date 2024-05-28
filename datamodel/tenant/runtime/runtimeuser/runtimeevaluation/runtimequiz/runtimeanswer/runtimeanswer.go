package runtimeanswer

import (
	"example/evaluation/quiz/quizquestion/quizchoice"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var RuntimeAnswerRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/tenant/{tenant.Tenant}/runtime/user/${runtimeuser.RuntimeUser}/quiz/{runtimequiz.RuntimeQuiz}/answer/{runtimeanswer.RuntimeAnswer}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/tenant/{tenant.Tenant}/runtime/user/${runtimeuser.RuntimeUser}/quiz/{runtimequiz.RuntimeQuiz}/answers",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:RuntimeAnswerRestAPISpec
type RuntimeAnswer struct {
	nexus.Node
	ProvidedAnswer string
	Answer         quizchoice.QuizChoice `nexus:"link"`
}
