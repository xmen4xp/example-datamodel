package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"

	nexus_client "example/build/nexus-client"
	"example/build/nexus-gql/graph/model"
)

var c = GrpcClients{
	mtx:     sync.Mutex{},
	Clients: map[string]GrpcClient{},
}
var nc *nexus_client.Clientset

func getParentName(parentLabels map[string]interface{}, key string) string {
	if v, ok := parentLabels[key]; ok && v != nil {
		return v.(string)
	}
	return ""
}

type NodeMetricTypeEnum string
type ServiceMetricTypeEnum string
type ServiceGroupByEnum string
type HTTPMethodEnum string
type EventSeverityEnum string
type AnalyticsMetricEnum string
type AnalyticsSubMetricEnum string
type TrafficDirectionEnum string
type SloDetailsEnum string

// ////////////////////////////////////
// Nexus K8sAPIEndpointConfig
// ////////////////////////////////////
func getK8sAPIEndpointConfig() *rest.Config {
	var (
		config *rest.Config
		err    error
	)
	filePath := os.Getenv("KUBECONFIG")
	if filePath != "" {
		config, err = clientcmd.BuildConfigFromFlags("", filePath)
		if err != nil {
			return nil
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil
		}
	}
	config.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(200, 300)
	return config
}

// ////////////////////////////////////
// Singleton Resolver for Parent Node
// PKG: Root, NODE: Root
// ////////////////////////////////////
func getRootResolver() (*model.RootRoot, error) {
	if nc == nil {
		k8sApiConfig := getK8sAPIEndpointConfig()
		nexusClient, err := nexus_client.NewForConfig(k8sApiConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to get k8s client config: %s", err)
		}
		nc = nexusClient
		nc.SubscribeAll()
		log.Debugf("Subscribed to all nodes in datamodel")
	}

	vRoot, err := nc.GetRootRoot(context.TODO())
	if err != nil {
		log.Errorf("[getRootResolver]Error getting Root node %s", err)
		return nil, nil
	}
	dn := vRoot.DisplayName()
	parentLabels := map[string]interface{}{"roots.root.example.com": dn}

	ret := &model.RootRoot{
		Id:           &dn,
		ParentLabels: parentLabels,
	}
	log.Debugf("[getRootResolver]Output Root object %+v", ret)
	return ret, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Tenant Node: Root PKG: Root
// ////////////////////////////////////
func getRootRootTenantResolver(obj *model.RootRoot, id *string) ([]*model.TenantTenant, error) {
	log.Debugf("[getRootRootTenantResolver]Parent Object %+v", obj)
	var vTenantTenantList []*model.TenantTenant
	if id != nil && *id != "" {
		log.Debugf("[getRootRootTenantResolver]Id %q", *id)
		vTenant, err := nc.RootRoot().GetTenant(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getRootRootTenantResolver]Error getting Tenant node %q : %s", *id, err)
			return vTenantTenantList, nil
		}
		dn := vTenant.DisplayName()
		parentLabels := map[string]interface{}{"tenants.tenant.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.TenantTenant{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vTenantTenantList = append(vTenantTenantList, ret)

		log.Debugf("[getRootRootTenantResolver]Output Tenant objects %v", vTenantTenantList)

		return vTenantTenantList, nil
	}

	log.Debug("[getRootRootTenantResolver]Id is empty, process all Tenants")

	vTenantParent, err := nc.GetRootRoot(context.TODO())
	if err != nil {
		log.Errorf("[getRootRootTenantResolver]Error getting parent node %s", err)
		return vTenantTenantList, nil
	}
	vTenantAllObj, err := vTenantParent.GetAllTenant(context.TODO())
	if err != nil {
		log.Errorf("[getRootRootTenantResolver]Error getting Tenant objects %s", err)
		return vTenantTenantList, nil
	}
	for _, i := range vTenantAllObj {
		vTenant, err := nc.RootRoot().GetTenant(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getRootRootTenantResolver]Error getting Tenant node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vTenant.DisplayName()
		parentLabels := map[string]interface{}{"tenants.tenant.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.TenantTenant{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vTenantTenantList = append(vTenantTenantList, ret)
	}

	log.Debugf("[getRootRootTenantResolver]Output Tenant objects %v", vTenantTenantList)

	return vTenantTenantList, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Evaluation Node: Root PKG: Root
// ////////////////////////////////////
func getRootRootEvaluationResolver(obj *model.RootRoot, id *string) ([]*model.EvaluationEvaluation, error) {
	log.Debugf("[getRootRootEvaluationResolver]Parent Object %+v", obj)
	var vEvaluationEvaluationList []*model.EvaluationEvaluation
	if id != nil && *id != "" {
		log.Debugf("[getRootRootEvaluationResolver]Id %q", *id)
		vEvaluation, err := nc.RootRoot().GetEvaluation(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getRootRootEvaluationResolver]Error getting Evaluation node %q : %s", *id, err)
			return vEvaluationEvaluationList, nil
		}
		dn := vEvaluation.DisplayName()
		parentLabels := map[string]interface{}{"evaluations.evaluation.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.EvaluationEvaluation{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vEvaluationEvaluationList = append(vEvaluationEvaluationList, ret)

		log.Debugf("[getRootRootEvaluationResolver]Output Evaluation objects %v", vEvaluationEvaluationList)

		return vEvaluationEvaluationList, nil
	}

	log.Debug("[getRootRootEvaluationResolver]Id is empty, process all Evaluations")

	vEvaluationParent, err := nc.GetRootRoot(context.TODO())
	if err != nil {
		log.Errorf("[getRootRootEvaluationResolver]Error getting parent node %s", err)
		return vEvaluationEvaluationList, nil
	}
	vEvaluationAllObj, err := vEvaluationParent.GetAllEvaluation(context.TODO())
	if err != nil {
		log.Errorf("[getRootRootEvaluationResolver]Error getting Evaluation objects %s", err)
		return vEvaluationEvaluationList, nil
	}
	for _, i := range vEvaluationAllObj {
		vEvaluation, err := nc.RootRoot().GetEvaluation(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getRootRootEvaluationResolver]Error getting Evaluation node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vEvaluation.DisplayName()
		parentLabels := map[string]interface{}{"evaluations.evaluation.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.EvaluationEvaluation{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vEvaluationEvaluationList = append(vEvaluationEvaluationList, ret)
	}

	log.Debugf("[getRootRootEvaluationResolver]Output Evaluation objects %v", vEvaluationEvaluationList)

	return vEvaluationEvaluationList, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Quiz Node: Evaluation PKG: Evaluation
// ////////////////////////////////////
func getEvaluationEvaluationQuizResolver(obj *model.EvaluationEvaluation, id *string) ([]*model.QuizQuiz, error) {
	log.Debugf("[getEvaluationEvaluationQuizResolver]Parent Object %+v", obj)
	var vQuizQuizList []*model.QuizQuiz
	if id != nil && *id != "" {
		log.Debugf("[getEvaluationEvaluationQuizResolver]Id %q", *id)
		vQuiz, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).GetQuiz(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getEvaluationEvaluationQuizResolver]Error getting Quiz node %q : %s", *id, err)
			return vQuizQuizList, nil
		}
		dn := vQuiz.DisplayName()
		parentLabels := map[string]interface{}{"quizes.quiz.example.com": dn}
		vDefaultScorePerQuestion := int(vQuiz.Spec.DefaultScorePerQuestion)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizQuiz{
			Id:                      &dn,
			ParentLabels:            parentLabels,
			DefaultScorePerQuestion: &vDefaultScorePerQuestion,
		}
		vQuizQuizList = append(vQuizQuizList, ret)

		log.Debugf("[getEvaluationEvaluationQuizResolver]Output Quiz objects %v", vQuizQuizList)

		return vQuizQuizList, nil
	}

	log.Debug("[getEvaluationEvaluationQuizResolver]Id is empty, process all Quizs")

	vQuizParent, err := nc.RootRoot().GetEvaluation(context.TODO(), getParentName(obj.ParentLabels, "evaluations.evaluation.example.com"))
	if err != nil {
		log.Errorf("[getEvaluationEvaluationQuizResolver]Error getting parent node %s", err)
		return vQuizQuizList, nil
	}
	vQuizAllObj, err := vQuizParent.GetAllQuiz(context.TODO())
	if err != nil {
		log.Errorf("[getEvaluationEvaluationQuizResolver]Error getting Quiz objects %s", err)
		return vQuizQuizList, nil
	}
	for _, i := range vQuizAllObj {
		vQuiz, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).GetQuiz(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getEvaluationEvaluationQuizResolver]Error getting Quiz node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vQuiz.DisplayName()
		parentLabels := map[string]interface{}{"quizes.quiz.example.com": dn}
		vDefaultScorePerQuestion := int(vQuiz.Spec.DefaultScorePerQuestion)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizQuiz{
			Id:                      &dn,
			ParentLabels:            parentLabels,
			DefaultScorePerQuestion: &vDefaultScorePerQuestion,
		}
		vQuizQuizList = append(vQuizQuizList, ret)
	}

	log.Debugf("[getEvaluationEvaluationQuizResolver]Output Quiz objects %v", vQuizQuizList)

	return vQuizQuizList, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Question Node: Quiz PKG: Quiz
// ////////////////////////////////////
func getQuizQuizQuestionResolver(obj *model.QuizQuiz, id *string) ([]*model.QuizquestionQuizQuestion, error) {
	log.Debugf("[getQuizQuizQuestionResolver]Parent Object %+v", obj)
	var vQuizquestionQuizQuestionList []*model.QuizquestionQuizQuestion
	if id != nil && *id != "" {
		log.Debugf("[getQuizQuizQuestionResolver]Id %q", *id)
		vQuizQuestion, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).GetQuestion(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getQuizQuizQuestionResolver]Error getting Question node %q : %s", *id, err)
			return vQuizquestionQuizQuestionList, nil
		}
		dn := vQuizQuestion.DisplayName()
		parentLabels := map[string]interface{}{"quizquestions.quizquestion.example.com": dn}
		vQuestion := string(vQuizQuestion.Spec.Question)
		vHint := string(vQuizQuestion.Spec.Hint)
		vAnswer := string(vQuizQuestion.Spec.Answer)
		vFormat := string(vQuizQuestion.Spec.Format)
		vScore := int(vQuizQuestion.Spec.Score)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizquestionQuizQuestion{
			Id:           &dn,
			ParentLabels: parentLabels,
			Question:     &vQuestion,
			Hint:         &vHint,
			Answer:       &vAnswer,
			Format:       &vFormat,
			Score:        &vScore,
		}
		vQuizquestionQuizQuestionList = append(vQuizquestionQuizQuestionList, ret)

		log.Debugf("[getQuizQuizQuestionResolver]Output Question objects %v", vQuizquestionQuizQuestionList)

		return vQuizquestionQuizQuestionList, nil
	}

	log.Debug("[getQuizQuizQuestionResolver]Id is empty, process all Questions")

	vQuizQuestionParent, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).GetQuiz(context.TODO(), getParentName(obj.ParentLabels, "quizes.quiz.example.com"))
	if err != nil {
		log.Errorf("[getQuizQuizQuestionResolver]Error getting parent node %s", err)
		return vQuizquestionQuizQuestionList, nil
	}
	vQuizQuestionAllObj, err := vQuizQuestionParent.GetAllQuestion(context.TODO())
	if err != nil {
		log.Errorf("[getQuizQuizQuestionResolver]Error getting Question objects %s", err)
		return vQuizquestionQuizQuestionList, nil
	}
	for _, i := range vQuizQuestionAllObj {
		vQuizQuestion, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).GetQuestion(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getQuizQuizQuestionResolver]Error getting Question node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vQuizQuestion.DisplayName()
		parentLabels := map[string]interface{}{"quizquestions.quizquestion.example.com": dn}
		vQuestion := string(vQuizQuestion.Spec.Question)
		vHint := string(vQuizQuestion.Spec.Hint)
		vAnswer := string(vQuizQuestion.Spec.Answer)
		vFormat := string(vQuizQuestion.Spec.Format)
		vScore := int(vQuizQuestion.Spec.Score)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizquestionQuizQuestion{
			Id:           &dn,
			ParentLabels: parentLabels,
			Question:     &vQuestion,
			Hint:         &vHint,
			Answer:       &vAnswer,
			Format:       &vFormat,
			Score:        &vScore,
		}
		vQuizquestionQuizQuestionList = append(vQuizquestionQuizQuestionList, ret)
	}

	log.Debugf("[getQuizQuizQuestionResolver]Output Question objects %v", vQuizquestionQuizQuestionList)

	return vQuizquestionQuizQuestionList, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Choice Node: QuizQuestion PKG: Quizquestion
// ////////////////////////////////////
func getQuizquestionQuizQuestionChoiceResolver(obj *model.QuizquestionQuizQuestion, id *string) ([]*model.QuizchoiceQuizChoice, error) {
	log.Debugf("[getQuizquestionQuizQuestionChoiceResolver]Parent Object %+v", obj)
	var vQuizchoiceQuizChoiceList []*model.QuizchoiceQuizChoice
	if id != nil && *id != "" {
		log.Debugf("[getQuizquestionQuizQuestionChoiceResolver]Id %q", *id)
		vQuizChoice, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).Question(getParentName(obj.ParentLabels, "quizquestions.quizquestion.example.com")).GetChoice(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getQuizquestionQuizQuestionChoiceResolver]Error getting Choice node %q : %s", *id, err)
			return vQuizchoiceQuizChoiceList, nil
		}
		dn := vQuizChoice.DisplayName()
		parentLabels := map[string]interface{}{"quizchoices.quizchoice.example.com": dn}
		vChoice := string(vQuizChoice.Spec.Choice)
		vHint := string(vQuizChoice.Spec.Hint)
		vPictureName := string(vQuizChoice.Spec.PictureName)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizchoiceQuizChoice{
			Id:           &dn,
			ParentLabels: parentLabels,
			Choice:       &vChoice,
			Hint:         &vHint,
			PictureName:  &vPictureName,
		}
		vQuizchoiceQuizChoiceList = append(vQuizchoiceQuizChoiceList, ret)

		log.Debugf("[getQuizquestionQuizQuestionChoiceResolver]Output Choice objects %v", vQuizchoiceQuizChoiceList)

		return vQuizchoiceQuizChoiceList, nil
	}

	log.Debug("[getQuizquestionQuizQuestionChoiceResolver]Id is empty, process all Choices")

	vQuizChoiceParent, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).GetQuestion(context.TODO(), getParentName(obj.ParentLabels, "quizquestions.quizquestion.example.com"))
	if err != nil {
		log.Errorf("[getQuizquestionQuizQuestionChoiceResolver]Error getting parent node %s", err)
		return vQuizchoiceQuizChoiceList, nil
	}
	vQuizChoiceAllObj, err := vQuizChoiceParent.GetAllChoice(context.TODO())
	if err != nil {
		log.Errorf("[getQuizquestionQuizQuestionChoiceResolver]Error getting Choice objects %s", err)
		return vQuizchoiceQuizChoiceList, nil
	}
	for _, i := range vQuizChoiceAllObj {
		vQuizChoice, err := nc.RootRoot().Evaluation(getParentName(obj.ParentLabels, "evaluations.evaluation.example.com")).Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).Question(getParentName(obj.ParentLabels, "quizquestions.quizquestion.example.com")).GetChoice(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getQuizquestionQuizQuestionChoiceResolver]Error getting Choice node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vQuizChoice.DisplayName()
		parentLabels := map[string]interface{}{"quizchoices.quizchoice.example.com": dn}
		vChoice := string(vQuizChoice.Spec.Choice)
		vHint := string(vQuizChoice.Spec.Hint)
		vPictureName := string(vQuizChoice.Spec.PictureName)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizchoiceQuizChoice{
			Id:           &dn,
			ParentLabels: parentLabels,
			Choice:       &vChoice,
			Hint:         &vHint,
			PictureName:  &vPictureName,
		}
		vQuizchoiceQuizChoiceList = append(vQuizchoiceQuizChoiceList, ret)
	}

	log.Debugf("[getQuizquestionQuizQuestionChoiceResolver]Output Choice objects %v", vQuizchoiceQuizChoiceList)

	return vQuizchoiceQuizChoiceList, nil
}

// ////////////////////////////////////
// CHILD RESOLVER (Singleton)
// FieldName: Config Node: Tenant PKG: Tenant
// ////////////////////////////////////
func getTenantTenantConfigResolver(obj *model.TenantTenant) (*model.ConfigConfig, error) {
	log.Debugf("[getTenantTenantConfigResolver]Parent Object %+v", obj)
	vConfig, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).GetConfig(context.TODO())
	if err != nil {
		log.Errorf("[getTenantTenantConfigResolver]Error getting Tenant node %s", err)
		return &model.ConfigConfig{}, nil
	}
	dn := vConfig.DisplayName()
	parentLabels := map[string]interface{}{"configs.config.example.com": dn}

	for k, v := range obj.ParentLabels {
		parentLabels[k] = v
	}
	ret := &model.ConfigConfig{
		Id:           &dn,
		ParentLabels: parentLabels,
	}

	log.Debugf("[getTenantTenantConfigResolver]Output object %+v", ret)
	return ret, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Interest Node: Tenant PKG: Tenant
// ////////////////////////////////////
func getTenantTenantInterestResolver(obj *model.TenantTenant, id *string) ([]*model.InterestInterest, error) {
	log.Debugf("[getTenantTenantInterestResolver]Parent Object %+v", obj)
	var vInterestInterestList []*model.InterestInterest
	if id != nil && *id != "" {
		log.Debugf("[getTenantTenantInterestResolver]Id %q", *id)
		vInterest, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).GetInterest(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getTenantTenantInterestResolver]Error getting Interest node %q : %s", *id, err)
			return vInterestInterestList, nil
		}
		dn := vInterest.DisplayName()
		parentLabels := map[string]interface{}{"interests.interest.example.com": dn}
		vName := string(vInterest.Spec.Name)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.InterestInterest{
			Id:           &dn,
			ParentLabels: parentLabels,
			Name:         &vName,
		}
		vInterestInterestList = append(vInterestInterestList, ret)

		log.Debugf("[getTenantTenantInterestResolver]Output Interest objects %v", vInterestInterestList)

		return vInterestInterestList, nil
	}

	log.Debug("[getTenantTenantInterestResolver]Id is empty, process all Interests")

	vInterestParent, err := nc.RootRoot().GetTenant(context.TODO(), getParentName(obj.ParentLabels, "tenants.tenant.example.com"))
	if err != nil {
		log.Errorf("[getTenantTenantInterestResolver]Error getting parent node %s", err)
		return vInterestInterestList, nil
	}
	vInterestAllObj, err := vInterestParent.GetAllInterest(context.TODO())
	if err != nil {
		log.Errorf("[getTenantTenantInterestResolver]Error getting Interest objects %s", err)
		return vInterestInterestList, nil
	}
	for _, i := range vInterestAllObj {
		vInterest, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).GetInterest(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getTenantTenantInterestResolver]Error getting Interest node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vInterest.DisplayName()
		parentLabels := map[string]interface{}{"interests.interest.example.com": dn}
		vName := string(vInterest.Spec.Name)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.InterestInterest{
			Id:           &dn,
			ParentLabels: parentLabels,
			Name:         &vName,
		}
		vInterestInterestList = append(vInterestInterestList, ret)
	}

	log.Debugf("[getTenantTenantInterestResolver]Output Interest objects %v", vInterestInterestList)

	return vInterestInterestList, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: User Node: Config PKG: Config
// ////////////////////////////////////
func getConfigConfigUserResolver(obj *model.ConfigConfig, id *string) ([]*model.UserUser, error) {
	log.Debugf("[getConfigConfigUserResolver]Parent Object %+v", obj)
	var vUserUserList []*model.UserUser
	if id != nil && *id != "" {
		log.Debugf("[getConfigConfigUserResolver]Id %q", *id)
		vUser, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().GetUser(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getConfigConfigUserResolver]Error getting User node %q : %s", *id, err)
			return vUserUserList, nil
		}
		dn := vUser.DisplayName()
		parentLabels := map[string]interface{}{"users.user.example.com": dn}
		vUsername := string(vUser.Spec.Username)
		vMail := string(vUser.Spec.Mail)
		vFirstName := string(vUser.Spec.FirstName)
		vLastName := string(vUser.Spec.LastName)
		vPassword := string(vUser.Spec.Password)
		vRealm := string(vUser.Spec.Realm)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.UserUser{
			Id:           &dn,
			ParentLabels: parentLabels,
			Username:     &vUsername,
			Mail:         &vMail,
			FirstName:    &vFirstName,
			LastName:     &vLastName,
			Password:     &vPassword,
			Realm:        &vRealm,
		}
		vUserUserList = append(vUserUserList, ret)

		log.Debugf("[getConfigConfigUserResolver]Output User objects %v", vUserUserList)

		return vUserUserList, nil
	}

	log.Debug("[getConfigConfigUserResolver]Id is empty, process all Users")

	vUserParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).GetConfig(context.TODO())
	if err != nil {
		log.Errorf("[getConfigConfigUserResolver]Error getting parent node %s", err)
		return vUserUserList, nil
	}
	vUserAllObj, err := vUserParent.GetAllUser(context.TODO())
	if err != nil {
		log.Errorf("[getConfigConfigUserResolver]Error getting User objects %s", err)
		return vUserUserList, nil
	}
	for _, i := range vUserAllObj {
		vUser, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().GetUser(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getConfigConfigUserResolver]Error getting User node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vUser.DisplayName()
		parentLabels := map[string]interface{}{"users.user.example.com": dn}
		vUsername := string(vUser.Spec.Username)
		vMail := string(vUser.Spec.Mail)
		vFirstName := string(vUser.Spec.FirstName)
		vLastName := string(vUser.Spec.LastName)
		vPassword := string(vUser.Spec.Password)
		vRealm := string(vUser.Spec.Realm)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.UserUser{
			Id:           &dn,
			ParentLabels: parentLabels,
			Username:     &vUsername,
			Mail:         &vMail,
			FirstName:    &vFirstName,
			LastName:     &vLastName,
			Password:     &vPassword,
			Realm:        &vRealm,
		}
		vUserUserList = append(vUserUserList, ret)
	}

	log.Debugf("[getConfigConfigUserResolver]Output User objects %v", vUserUserList)

	return vUserUserList, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Event Node: Config PKG: Config
// ////////////////////////////////////
func getConfigConfigEventResolver(obj *model.ConfigConfig, id *string) ([]*model.EventEvent, error) {
	log.Debugf("[getConfigConfigEventResolver]Parent Object %+v", obj)
	var vEventEventList []*model.EventEvent
	if id != nil && *id != "" {
		log.Debugf("[getConfigConfigEventResolver]Id %q", *id)
		vEvent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().GetEvent(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getConfigConfigEventResolver]Error getting Event node %q : %s", *id, err)
			return vEventEventList, nil
		}
		dn := vEvent.DisplayName()
		parentLabels := map[string]interface{}{"events.event.example.com": dn}
		vDescription := string(vEvent.Spec.Description)
		vMeetingLink := string(vEvent.Spec.MeetingLink)
		Time, _ := json.Marshal(vEvent.Spec.Time)
		TimeData := string(Time)
		vPublic := bool(vEvent.Spec.Public)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.EventEvent{
			Id:           &dn,
			ParentLabels: parentLabels,
			Description:  &vDescription,
			MeetingLink:  &vMeetingLink,
			Time:         &TimeData,
			Public:       &vPublic,
		}
		vEventEventList = append(vEventEventList, ret)

		log.Debugf("[getConfigConfigEventResolver]Output Event objects %v", vEventEventList)

		return vEventEventList, nil
	}

	log.Debug("[getConfigConfigEventResolver]Id is empty, process all Events")

	vEventParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).GetConfig(context.TODO())
	if err != nil {
		log.Errorf("[getConfigConfigEventResolver]Error getting parent node %s", err)
		return vEventEventList, nil
	}
	vEventAllObj, err := vEventParent.GetAllEvent(context.TODO())
	if err != nil {
		log.Errorf("[getConfigConfigEventResolver]Error getting Event objects %s", err)
		return vEventEventList, nil
	}
	for _, i := range vEventAllObj {
		vEvent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().GetEvent(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getConfigConfigEventResolver]Error getting Event node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vEvent.DisplayName()
		parentLabels := map[string]interface{}{"events.event.example.com": dn}
		vDescription := string(vEvent.Spec.Description)
		vMeetingLink := string(vEvent.Spec.MeetingLink)
		Time, _ := json.Marshal(vEvent.Spec.Time)
		TimeData := string(Time)
		vPublic := bool(vEvent.Spec.Public)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.EventEvent{
			Id:           &dn,
			ParentLabels: parentLabels,
			Description:  &vDescription,
			MeetingLink:  &vMeetingLink,
			Time:         &TimeData,
			Public:       &vPublic,
		}
		vEventEventList = append(vEventEventList, ret)
	}

	log.Debugf("[getConfigConfigEventResolver]Output Event objects %v", vEventEventList)

	return vEventEventList, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Wanna Node: User PKG: User
// ////////////////////////////////////
func getUserUserWannaResolver(obj *model.UserUser, id *string) ([]*model.WannaWanna, error) {
	log.Debugf("[getUserUserWannaResolver]Parent Object %+v", obj)
	var vWannaWannaList []*model.WannaWanna
	if id != nil && *id != "" {
		log.Debugf("[getUserUserWannaResolver]Id %q", *id)
		vWanna, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().User(getParentName(obj.ParentLabels, "users.user.example.com")).GetWanna(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getUserUserWannaResolver]Error getting Wanna node %q : %s", *id, err)
			return vWannaWannaList, nil
		}
		dn := vWanna.DisplayName()
		parentLabels := map[string]interface{}{"wannas.wanna.example.com": dn}
		vName := string(vWanna.Spec.Name)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.WannaWanna{
			Id:           &dn,
			ParentLabels: parentLabels,
			Name:         &vName,
		}
		vWannaWannaList = append(vWannaWannaList, ret)

		log.Debugf("[getUserUserWannaResolver]Output Wanna objects %v", vWannaWannaList)

		return vWannaWannaList, nil
	}

	log.Debug("[getUserUserWannaResolver]Id is empty, process all Wannas")

	vWannaParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().GetUser(context.TODO(), getParentName(obj.ParentLabels, "users.user.example.com"))
	if err != nil {
		log.Errorf("[getUserUserWannaResolver]Error getting parent node %s", err)
		return vWannaWannaList, nil
	}
	vWannaAllObj, err := vWannaParent.GetAllWanna(context.TODO())
	if err != nil {
		log.Errorf("[getUserUserWannaResolver]Error getting Wanna objects %s", err)
		return vWannaWannaList, nil
	}
	for _, i := range vWannaAllObj {
		vWanna, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().User(getParentName(obj.ParentLabels, "users.user.example.com")).GetWanna(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getUserUserWannaResolver]Error getting Wanna node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vWanna.DisplayName()
		parentLabels := map[string]interface{}{"wannas.wanna.example.com": dn}
		vName := string(vWanna.Spec.Name)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.WannaWanna{
			Id:           &dn,
			ParentLabels: parentLabels,
			Name:         &vName,
		}
		vWannaWannaList = append(vWannaWannaList, ret)
	}

	log.Debugf("[getUserUserWannaResolver]Output Wanna objects %v", vWannaWannaList)

	return vWannaWannaList, nil
}

// ////////////////////////////////////
// LINK RESOLVER
// FieldName: Interest Node: Wanna PKG: Wanna
// ////////////////////////////////////
func getWannaWannaInterestResolver(obj *model.WannaWanna) (*model.InterestInterest, error) {
	log.Debugf("[getWannaWannaInterestResolver]Parent Object %+v", obj)
	vInterestParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Config().User(getParentName(obj.ParentLabels, "users.user.example.com")).GetWanna(context.TODO(), getParentName(obj.ParentLabels, "wannas.wanna.example.com"))
	if err != nil {
		log.Errorf("[getWannaWannaInterestResolver]Error getting parent node %s", err)
		return &model.InterestInterest{}, nil
	}
	vInterest, err := vInterestParent.GetInterest(context.TODO())
	if err != nil {
		log.Errorf("[getWannaWannaInterestResolver]Error getting Interest object %s", err)
		return &model.InterestInterest{}, nil
	}
	dn := vInterest.DisplayName()
	parentLabels := map[string]interface{}{"interests.interest.example.com": dn}
	vName := string(vInterest.Spec.Name)

	for k, v := range obj.ParentLabels {
		parentLabels[k] = v
	}
	ret := &model.InterestInterest{
		Id:           &dn,
		ParentLabels: parentLabels,
		Name:         &vName,
	}
	log.Debugf("[getWannaWannaInterestResolver]Output object %v", ret)

	return ret, nil
}
