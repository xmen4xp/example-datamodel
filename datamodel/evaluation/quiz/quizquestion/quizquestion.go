package quizquestion

import (
	quizchoice "example/evaluation/quiz/quizquestion/quizchoices"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

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

	// Name for this children will be numerical string to help sequence the choices.
	Choice quizchoice.QuizChoice `nexus:"children"`
}
