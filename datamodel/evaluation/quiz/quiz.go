package quiz

import (
	"example/evaluation/quiz/quizquestion"

	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type Quiz struct {
	nexus.Node

	DefaultScorePerQuestion int `json:",omitempty"`

	// Name for this children will be numerical string to help sequence the choices.
	Question quizquestion.QuizQuestion `nexus:"children"`
}
