package quiz

import (
	"example/evaluation/quiz/quizquestion"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var QuizRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/eval/quiz/{quiz.Quiz}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/eval/quizes",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:QuizRestAPISpec
type Quiz struct {
	nexus.Node

	DefaultScorePerQuestion int `json:",omitempty"`

	// Name for this children will be numerical string to help sequence the choices.
	Question quizquestion.QuizQuestion `nexus:"children"`

	Status QuizStatus `nexus:"status"`
}

type QuizStatus struct {
	TotalScore int
}
