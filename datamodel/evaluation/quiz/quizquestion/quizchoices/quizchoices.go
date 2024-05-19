package quizchoice

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var QuizChoicesRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/eval/quiz/{quiz.Quiz}/question/{quizquestion.QuizQuestion}/choice/{quizchoice.QuizChoice}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/eval/quiz/{quiz.Quiz}/question/{quizquestion.QuizQuestion}/choices",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:QuizChoicesRestAPISpec
type QuizChoice struct {
	nexus.Node

	Choice      string
	Hint        string
	PictureName string
}
