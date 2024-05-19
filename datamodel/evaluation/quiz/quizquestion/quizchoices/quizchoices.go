package quizchoice

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type QuizChoice struct {
	nexus.Node
	Choice      string
	Hint        string
	PictureName string
}
