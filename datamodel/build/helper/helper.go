package helper

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/elliotchance/orderedmap"

	datamodel "example/build/client/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const DEFAULT_KEY = "default"
const DISPLAY_NAME_LABEL = "nexus/display_name"
const IS_NAME_HASHED_LABEL = "nexus/is_name_hashed"

func GetCRDParentsMap() map[string][]string {
	return map[string][]string{
		"categories.category.example.com":                  {"roots.root.example.com"},
		"configs.config.example.com":                       {"roots.root.example.com", "tenants.tenant.example.com"},
		"evaluations.evaluation.example.com":               {"roots.root.example.com"},
		"events.event.example.com":                         {"roots.root.example.com", "tenants.tenant.example.com", "configs.config.example.com"},
		"interests.interest.example.com":                   {"roots.root.example.com", "tenants.tenant.example.com"},
		"quizchoices.quizchoice.example.com":               {"roots.root.example.com", "evaluations.evaluation.example.com", "quizes.quiz.example.com", "quizquestions.quizquestion.example.com"},
		"quizes.quiz.example.com":                          {"roots.root.example.com", "evaluations.evaluation.example.com"},
		"quizquestions.quizquestion.example.com":           {"roots.root.example.com", "evaluations.evaluation.example.com", "quizes.quiz.example.com"},
		"roots.root.example.com":                           {},
		"runtimeanswers.runtimeanswer.example.com":         {"roots.root.example.com", "tenants.tenant.example.com", "runtimes.runtime.example.com", "runtimeusers.runtimeuser.example.com", "runtimeevaluations.runtimeevaluation.example.com", "runtimequizes.runtimequiz.example.com"},
		"runtimeevaluations.runtimeevaluation.example.com": {"roots.root.example.com", "tenants.tenant.example.com", "runtimes.runtime.example.com", "runtimeusers.runtimeuser.example.com"},
		"runtimequizes.runtimequiz.example.com":            {"roots.root.example.com", "tenants.tenant.example.com", "runtimes.runtime.example.com", "runtimeusers.runtimeuser.example.com", "runtimeevaluations.runtimeevaluation.example.com"},
		"runtimes.runtime.example.com":                     {"roots.root.example.com", "tenants.tenant.example.com"},
		"runtimeusers.runtimeuser.example.com":             {"roots.root.example.com", "tenants.tenant.example.com", "runtimes.runtime.example.com"},
		"tenants.tenant.example.com":                       {"roots.root.example.com"},
		"users.user.example.com":                           {"roots.root.example.com", "tenants.tenant.example.com", "configs.config.example.com"},
		"wannas.wanna.example.com":                         {"roots.root.example.com", "tenants.tenant.example.com", "configs.config.example.com", "users.user.example.com"},
	}
}

func GetObjectByCRDName(dmClient *datamodel.Clientset, crdName string, name string) interface{} {
	if crdName == "categories.category.example.com" {
		obj, err := dmClient.CategoryExampleV1().Categories().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "configs.config.example.com" {
		obj, err := dmClient.ConfigExampleV1().Configs().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "evaluations.evaluation.example.com" {
		obj, err := dmClient.EvaluationExampleV1().Evaluations().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "events.event.example.com" {
		obj, err := dmClient.EventExampleV1().Events().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "interests.interest.example.com" {
		obj, err := dmClient.InterestExampleV1().Interests().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "quizchoices.quizchoice.example.com" {
		obj, err := dmClient.QuizchoiceExampleV1().QuizChoices().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "quizes.quiz.example.com" {
		obj, err := dmClient.QuizExampleV1().Quizes().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "quizquestions.quizquestion.example.com" {
		obj, err := dmClient.QuizquestionExampleV1().QuizQuestions().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "roots.root.example.com" {
		obj, err := dmClient.RootExampleV1().Roots().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "runtimeanswers.runtimeanswer.example.com" {
		obj, err := dmClient.RuntimeanswerExampleV1().RuntimeAnswers().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "runtimeevaluations.runtimeevaluation.example.com" {
		obj, err := dmClient.RuntimeevaluationExampleV1().RuntimeEvaluations().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "runtimequizes.runtimequiz.example.com" {
		obj, err := dmClient.RuntimequizExampleV1().RuntimeQuizes().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "runtimes.runtime.example.com" {
		obj, err := dmClient.RuntimeExampleV1().Runtimes().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "runtimeusers.runtimeuser.example.com" {
		obj, err := dmClient.RuntimeuserExampleV1().RuntimeUsers().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "tenants.tenant.example.com" {
		obj, err := dmClient.TenantExampleV1().Tenants().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "users.user.example.com" {
		obj, err := dmClient.UserExampleV1().Users().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}
	if crdName == "wannas.wanna.example.com" {
		obj, err := dmClient.WannaExampleV1().Wannas().Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return nil
		}
		return obj
	}

	return nil
}

func ParseCRDLabels(crdName string, labels map[string]string) *orderedmap.OrderedMap {
	parents := GetCRDParentsMap()[crdName]

	m := orderedmap.NewOrderedMap()
	for _, parent := range parents {
		if label, ok := labels[parent]; ok {
			m.Set(parent, label)
		} else {
			m.Set(parent, DEFAULT_KEY)
		}
	}

	return m
}

func GetHashedName(crdName string, labels map[string]string, name string) string {
	orderedLabels := ParseCRDLabels(crdName, labels)

	var output string
	for i, key := range orderedLabels.Keys() {
		value, _ := orderedLabels.Get(key)

		output += fmt.Sprintf("%s:%s", key, value)
		if i < orderedLabels.Len()-1 {
			output += "/"
		}
	}

	output += fmt.Sprintf("%s:%s", crdName, name)

	h := sha1.New()
	_, _ = h.Write([]byte(output))
	return hex.EncodeToString(h.Sum(nil))
}
