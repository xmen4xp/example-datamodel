package quizquestion

import (
	quizchoice "example/evaluation/quiz/quizquestion/quizchoices"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

var QuizQuestionRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/eval/quiz/{quiz.Quiz}/question/{quizquestion.QuizQuestion}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/eval/quiz/{quiz.Quiz}/questions",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:QuizQuestionRestAPISpec
type QuizQuestion struct {
	nexus.Node
	Question string
	Hint     string
	Answer   string `json:",omitempty"`

	// A question may support multiple formats:
	// a. multiple-choice
	// b. q&a
	// The format field is a enum that specifies the format of the question.
	Format string

	Score int

	AnimationFilePath string
	PictureFilePath   string

	// Name for this children will be numerical string to help sequence the choices.
	Choice quizchoice.QuizChoice `nexus:"children"`
}
