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
// CHILD RESOLVER (Singleton)
// FieldName: Evaluation Node: Root PKG: Root
// ////////////////////////////////////
func getRootRootEvaluationResolver(obj *model.RootRoot) (*model.EvaluationEvaluation, error) {
	log.Debugf("[getRootRootEvaluationResolver]Parent Object %+v", obj)
	vEvaluation, err := nc.RootRoot().GetEvaluation(context.TODO())
	if err != nil {
		log.Errorf("[getRootRootEvaluationResolver]Error getting Root node %s", err)
		return &model.EvaluationEvaluation{}, nil
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

	log.Debugf("[getRootRootEvaluationResolver]Output object %+v", ret)
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
// FieldName: Category Node: Root PKG: Root
// ////////////////////////////////////
func getRootRootCategoryResolver(obj *model.RootRoot, id *string) ([]*model.CategoryCategory, error) {
	log.Debugf("[getRootRootCategoryResolver]Parent Object %+v", obj)
	var vCategoryCategoryList []*model.CategoryCategory
	if id != nil && *id != "" {
		log.Debugf("[getRootRootCategoryResolver]Id %q", *id)
		vCategory, err := nc.RootRoot().GetCategory(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getRootRootCategoryResolver]Error getting Category node %q : %s", *id, err)
			return vCategoryCategoryList, nil
		}
		dn := vCategory.DisplayName()
		parentLabels := map[string]interface{}{"categories.category.example.com": dn}
		vDesription := string(vCategory.Spec.Desription)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.CategoryCategory{
			Id:           &dn,
			ParentLabels: parentLabels,
			Desription:   &vDesription,
		}
		vCategoryCategoryList = append(vCategoryCategoryList, ret)

		log.Debugf("[getRootRootCategoryResolver]Output Category objects %v", vCategoryCategoryList)

		return vCategoryCategoryList, nil
	}

	log.Debug("[getRootRootCategoryResolver]Id is empty, process all Categorys")

	vCategoryParent, err := nc.GetRootRoot(context.TODO())
	if err != nil {
		log.Errorf("[getRootRootCategoryResolver]Error getting parent node %s", err)
		return vCategoryCategoryList, nil
	}
	vCategoryAllObj, err := vCategoryParent.GetAllCategory(context.TODO())
	if err != nil {
		log.Errorf("[getRootRootCategoryResolver]Error getting Category objects %s", err)
		return vCategoryCategoryList, nil
	}
	for _, i := range vCategoryAllObj {
		vCategory, err := nc.RootRoot().GetCategory(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getRootRootCategoryResolver]Error getting Category node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vCategory.DisplayName()
		parentLabels := map[string]interface{}{"categories.category.example.com": dn}
		vDesription := string(vCategory.Spec.Desription)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.CategoryCategory{
			Id:           &dn,
			ParentLabels: parentLabels,
			Desription:   &vDesription,
		}
		vCategoryCategoryList = append(vCategoryCategoryList, ret)
	}

	log.Debugf("[getRootRootCategoryResolver]Output Category objects %v", vCategoryCategoryList)

	return vCategoryCategoryList, nil
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
		vQuiz, err := nc.RootRoot().Evaluation().GetQuiz(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getEvaluationEvaluationQuizResolver]Error getting Quiz node %q : %s", *id, err)
			return vQuizQuizList, nil
		}
		dn := vQuiz.DisplayName()
		parentLabels := map[string]interface{}{"quizes.quiz.example.com": dn}
		Labels, _ := json.Marshal(vQuiz.Spec.Labels)
		LabelsData := string(Labels)
		vDefaultScorePerQuestion := int(vQuiz.Spec.DefaultScorePerQuestion)
		vDescription := string(vQuiz.Spec.Description)
		Categories, _ := json.Marshal(vQuiz.Spec.Categories)
		CategoriesData := string(Categories)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizQuiz{
			Id:                      &dn,
			ParentLabels:            parentLabels,
			Labels:                  &LabelsData,
			DefaultScorePerQuestion: &vDefaultScorePerQuestion,
			Description:             &vDescription,
			Categories:              &CategoriesData,
		}
		vQuizQuizList = append(vQuizQuizList, ret)

		log.Debugf("[getEvaluationEvaluationQuizResolver]Output Quiz objects %v", vQuizQuizList)

		return vQuizQuizList, nil
	}

	log.Debug("[getEvaluationEvaluationQuizResolver]Id is empty, process all Quizs")

	vQuizParent, err := nc.RootRoot().GetEvaluation(context.TODO())
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
		vQuiz, err := nc.RootRoot().Evaluation().GetQuiz(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getEvaluationEvaluationQuizResolver]Error getting Quiz node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vQuiz.DisplayName()
		parentLabels := map[string]interface{}{"quizes.quiz.example.com": dn}
		Labels, _ := json.Marshal(vQuiz.Spec.Labels)
		LabelsData := string(Labels)
		vDefaultScorePerQuestion := int(vQuiz.Spec.DefaultScorePerQuestion)
		vDescription := string(vQuiz.Spec.Description)
		Categories, _ := json.Marshal(vQuiz.Spec.Categories)
		CategoriesData := string(Categories)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizQuiz{
			Id:                      &dn,
			ParentLabels:            parentLabels,
			Labels:                  &LabelsData,
			DefaultScorePerQuestion: &vDefaultScorePerQuestion,
			Description:             &vDescription,
			Categories:              &CategoriesData,
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
		vQuizQuestion, err := nc.RootRoot().Evaluation().Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).GetQuestion(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getQuizQuizQuestionResolver]Error getting Question node %q : %s", *id, err)
			return vQuizquestionQuizQuestionList, nil
		}
		dn := vQuizQuestion.DisplayName()
		parentLabels := map[string]interface{}{"quizquestions.quizquestion.example.com": dn}
		vQuestion := string(vQuizQuestion.Spec.Question)
		vHint := string(vQuizQuestion.Spec.Hint)
		vFormat := string(vQuizQuestion.Spec.Format)
		vScore := int(vQuizQuestion.Spec.Score)
		vAnimationFilePath := string(vQuizQuestion.Spec.AnimationFilePath)
		vPictureFilePath := string(vQuizQuestion.Spec.PictureFilePath)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizquestionQuizQuestion{
			Id:                &dn,
			ParentLabels:      parentLabels,
			Question:          &vQuestion,
			Hint:              &vHint,
			Format:            &vFormat,
			Score:             &vScore,
			AnimationFilePath: &vAnimationFilePath,
			PictureFilePath:   &vPictureFilePath,
		}
		vQuizquestionQuizQuestionList = append(vQuizquestionQuizQuestionList, ret)

		log.Debugf("[getQuizQuizQuestionResolver]Output Question objects %v", vQuizquestionQuizQuestionList)

		return vQuizquestionQuizQuestionList, nil
	}

	log.Debug("[getQuizQuizQuestionResolver]Id is empty, process all Questions")

	vQuizQuestionParent, err := nc.RootRoot().Evaluation().GetQuiz(context.TODO(), getParentName(obj.ParentLabels, "quizes.quiz.example.com"))
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
		vQuizQuestion, err := nc.RootRoot().Evaluation().Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).GetQuestion(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getQuizQuizQuestionResolver]Error getting Question node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vQuizQuestion.DisplayName()
		parentLabels := map[string]interface{}{"quizquestions.quizquestion.example.com": dn}
		vQuestion := string(vQuizQuestion.Spec.Question)
		vHint := string(vQuizQuestion.Spec.Hint)
		vFormat := string(vQuizQuestion.Spec.Format)
		vScore := int(vQuizQuestion.Spec.Score)
		vAnimationFilePath := string(vQuizQuestion.Spec.AnimationFilePath)
		vPictureFilePath := string(vQuizQuestion.Spec.PictureFilePath)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizquestionQuizQuestion{
			Id:                &dn,
			ParentLabels:      parentLabels,
			Question:          &vQuestion,
			Hint:              &vHint,
			Format:            &vFormat,
			Score:             &vScore,
			AnimationFilePath: &vAnimationFilePath,
			PictureFilePath:   &vPictureFilePath,
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
		vQuizChoice, err := nc.RootRoot().Evaluation().Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).Question(getParentName(obj.ParentLabels, "quizquestions.quizquestion.example.com")).GetChoice(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getQuizquestionQuizQuestionChoiceResolver]Error getting Choice node %q : %s", *id, err)
			return vQuizchoiceQuizChoiceList, nil
		}
		dn := vQuizChoice.DisplayName()
		parentLabels := map[string]interface{}{"quizchoices.quizchoice.example.com": dn}
		vChoice := string(vQuizChoice.Spec.Choice)
		vHint := string(vQuizChoice.Spec.Hint)
		vPictureName := string(vQuizChoice.Spec.PictureName)
		vAnswer := bool(vQuizChoice.Spec.Answer)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizchoiceQuizChoice{
			Id:           &dn,
			ParentLabels: parentLabels,
			Choice:       &vChoice,
			Hint:         &vHint,
			PictureName:  &vPictureName,
			Answer:       &vAnswer,
		}
		vQuizchoiceQuizChoiceList = append(vQuizchoiceQuizChoiceList, ret)

		log.Debugf("[getQuizquestionQuizQuestionChoiceResolver]Output Choice objects %v", vQuizchoiceQuizChoiceList)

		return vQuizchoiceQuizChoiceList, nil
	}

	log.Debug("[getQuizquestionQuizQuestionChoiceResolver]Id is empty, process all Choices")

	vQuizChoiceParent, err := nc.RootRoot().Evaluation().Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).GetQuestion(context.TODO(), getParentName(obj.ParentLabels, "quizquestions.quizquestion.example.com"))
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
		vQuizChoice, err := nc.RootRoot().Evaluation().Quiz(getParentName(obj.ParentLabels, "quizes.quiz.example.com")).Question(getParentName(obj.ParentLabels, "quizquestions.quizquestion.example.com")).GetChoice(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getQuizquestionQuizQuestionChoiceResolver]Error getting Choice node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vQuizChoice.DisplayName()
		parentLabels := map[string]interface{}{"quizchoices.quizchoice.example.com": dn}
		vChoice := string(vQuizChoice.Spec.Choice)
		vHint := string(vQuizChoice.Spec.Hint)
		vPictureName := string(vQuizChoice.Spec.PictureName)
		vAnswer := bool(vQuizChoice.Spec.Answer)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.QuizchoiceQuizChoice{
			Id:           &dn,
			ParentLabels: parentLabels,
			Choice:       &vChoice,
			Hint:         &vHint,
			PictureName:  &vPictureName,
			Answer:       &vAnswer,
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
// CHILD RESOLVER (Singleton)
// FieldName: Runtime Node: Tenant PKG: Tenant
// ////////////////////////////////////
func getTenantTenantRuntimeResolver(obj *model.TenantTenant) (*model.RuntimeRuntime, error) {
	log.Debugf("[getTenantTenantRuntimeResolver]Parent Object %+v", obj)
	vRuntime, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).GetRuntime(context.TODO())
	if err != nil {
		log.Errorf("[getTenantTenantRuntimeResolver]Error getting Tenant node %s", err)
		return &model.RuntimeRuntime{}, nil
	}
	dn := vRuntime.DisplayName()
	parentLabels := map[string]interface{}{"runtimes.runtime.example.com": dn}

	for k, v := range obj.ParentLabels {
		parentLabels[k] = v
	}
	ret := &model.RuntimeRuntime{
		Id:           &dn,
		ParentLabels: parentLabels,
	}

	log.Debugf("[getTenantTenantRuntimeResolver]Output object %+v", ret)
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
		Type, _ := json.Marshal(vWanna.Spec.Type)
		TypeData := string(Type)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.WannaWanna{
			Id:           &dn,
			ParentLabels: parentLabels,
			Name:         &vName,
			Type:         &TypeData,
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
		Type, _ := json.Marshal(vWanna.Spec.Type)
		TypeData := string(Type)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.WannaWanna{
			Id:           &dn,
			ParentLabels: parentLabels,
			Name:         &vName,
			Type:         &TypeData,
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

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: User Node: Runtime PKG: Runtime
// ////////////////////////////////////
func getRuntimeRuntimeUserResolver(obj *model.RuntimeRuntime, id *string) ([]*model.RuntimeuserRuntimeUser, error) {
	log.Debugf("[getRuntimeRuntimeUserResolver]Parent Object %+v", obj)
	var vRuntimeuserRuntimeUserList []*model.RuntimeuserRuntimeUser
	if id != nil && *id != "" {
		log.Debugf("[getRuntimeRuntimeUserResolver]Id %q", *id)
		vRuntimeUser, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().GetUser(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getRuntimeRuntimeUserResolver]Error getting User node %q : %s", *id, err)
			return vRuntimeuserRuntimeUserList, nil
		}
		dn := vRuntimeUser.DisplayName()
		parentLabels := map[string]interface{}{"runtimeusers.runtimeuser.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.RuntimeuserRuntimeUser{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vRuntimeuserRuntimeUserList = append(vRuntimeuserRuntimeUserList, ret)

		log.Debugf("[getRuntimeRuntimeUserResolver]Output User objects %v", vRuntimeuserRuntimeUserList)

		return vRuntimeuserRuntimeUserList, nil
	}

	log.Debug("[getRuntimeRuntimeUserResolver]Id is empty, process all Users")

	vRuntimeUserParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).GetRuntime(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimeRuntimeUserResolver]Error getting parent node %s", err)
		return vRuntimeuserRuntimeUserList, nil
	}
	vRuntimeUserAllObj, err := vRuntimeUserParent.GetAllUser(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimeRuntimeUserResolver]Error getting User objects %s", err)
		return vRuntimeuserRuntimeUserList, nil
	}
	for _, i := range vRuntimeUserAllObj {
		vRuntimeUser, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().GetUser(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getRuntimeRuntimeUserResolver]Error getting User node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vRuntimeUser.DisplayName()
		parentLabels := map[string]interface{}{"runtimeusers.runtimeuser.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.RuntimeuserRuntimeUser{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vRuntimeuserRuntimeUserList = append(vRuntimeuserRuntimeUserList, ret)
	}

	log.Debugf("[getRuntimeRuntimeUserResolver]Output User objects %v", vRuntimeuserRuntimeUserList)

	return vRuntimeuserRuntimeUserList, nil
}

// ////////////////////////////////////
// CHILD RESOLVER (Singleton)
// FieldName: Evaluation Node: RuntimeUser PKG: Runtimeuser
// ////////////////////////////////////
func getRuntimeuserRuntimeUserEvaluationResolver(obj *model.RuntimeuserRuntimeUser) (*model.RuntimeevaluationRuntimeEvaluation, error) {
	log.Debugf("[getRuntimeuserRuntimeUserEvaluationResolver]Parent Object %+v", obj)
	vRuntimeEvaluation, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).GetEvaluation(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimeuserRuntimeUserEvaluationResolver]Error getting RuntimeUser node %s", err)
		return &model.RuntimeevaluationRuntimeEvaluation{}, nil
	}
	dn := vRuntimeEvaluation.DisplayName()
	parentLabels := map[string]interface{}{"runtimeevaluations.runtimeevaluation.example.com": dn}

	for k, v := range obj.ParentLabels {
		parentLabels[k] = v
	}
	ret := &model.RuntimeevaluationRuntimeEvaluation{
		Id:           &dn,
		ParentLabels: parentLabels,
	}

	log.Debugf("[getRuntimeuserRuntimeUserEvaluationResolver]Output object %+v", ret)
	return ret, nil
}

// ////////////////////////////////////
// LINK RESOLVER
// FieldName: User Node: RuntimeUser PKG: Runtimeuser
// ////////////////////////////////////
func getRuntimeuserRuntimeUserUserResolver(obj *model.RuntimeuserRuntimeUser) (*model.UserUser, error) {
	log.Debugf("[getRuntimeuserRuntimeUserUserResolver]Parent Object %+v", obj)
	vUserParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().GetUser(context.TODO(), getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com"))
	if err != nil {
		log.Errorf("[getRuntimeuserRuntimeUserUserResolver]Error getting parent node %s", err)
		return &model.UserUser{}, nil
	}
	vUser, err := vUserParent.GetUser(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimeuserRuntimeUserUserResolver]Error getting User object %s", err)
		return &model.UserUser{}, nil
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
	log.Debugf("[getRuntimeuserRuntimeUserUserResolver]Output object %v", ret)

	return ret, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Quiz Node: RuntimeEvaluation PKG: Runtimeevaluation
// ////////////////////////////////////
func getRuntimeevaluationRuntimeEvaluationQuizResolver(obj *model.RuntimeevaluationRuntimeEvaluation, id *string) ([]*model.RuntimequizRuntimeQuiz, error) {
	log.Debugf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Parent Object %+v", obj)
	var vRuntimequizRuntimeQuizList []*model.RuntimequizRuntimeQuiz
	if id != nil && *id != "" {
		log.Debugf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Id %q", *id)
		vRuntimeQuiz, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).Evaluation().GetQuiz(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Error getting Quiz node %q : %s", *id, err)
			return vRuntimequizRuntimeQuizList, nil
		}
		dn := vRuntimeQuiz.DisplayName()
		parentLabels := map[string]interface{}{"runtimequizes.runtimequiz.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.RuntimequizRuntimeQuiz{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vRuntimequizRuntimeQuizList = append(vRuntimequizRuntimeQuizList, ret)

		log.Debugf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Output Quiz objects %v", vRuntimequizRuntimeQuizList)

		return vRuntimequizRuntimeQuizList, nil
	}

	log.Debug("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Id is empty, process all Quizs")

	vRuntimeQuizParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).GetEvaluation(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Error getting parent node %s", err)
		return vRuntimequizRuntimeQuizList, nil
	}
	vRuntimeQuizAllObj, err := vRuntimeQuizParent.GetAllQuiz(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Error getting Quiz objects %s", err)
		return vRuntimequizRuntimeQuizList, nil
	}
	for _, i := range vRuntimeQuizAllObj {
		vRuntimeQuiz, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).Evaluation().GetQuiz(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Error getting Quiz node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vRuntimeQuiz.DisplayName()
		parentLabels := map[string]interface{}{"runtimequizes.runtimequiz.example.com": dn}

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.RuntimequizRuntimeQuiz{
			Id:           &dn,
			ParentLabels: parentLabels,
		}
		vRuntimequizRuntimeQuizList = append(vRuntimequizRuntimeQuizList, ret)
	}

	log.Debugf("[getRuntimeevaluationRuntimeEvaluationQuizResolver]Output Quiz objects %v", vRuntimequizRuntimeQuizList)

	return vRuntimequizRuntimeQuizList, nil
}

// ////////////////////////////////////
// LINK RESOLVER
// FieldName: Quiz Node: RuntimeQuiz PKG: Runtimequiz
// ////////////////////////////////////
func getRuntimequizRuntimeQuizQuizResolver(obj *model.RuntimequizRuntimeQuiz) (*model.QuizQuiz, error) {
	log.Debugf("[getRuntimequizRuntimeQuizQuizResolver]Parent Object %+v", obj)
	vQuizParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).Evaluation().GetQuiz(context.TODO(), getParentName(obj.ParentLabels, "runtimequizes.runtimequiz.example.com"))
	if err != nil {
		log.Errorf("[getRuntimequizRuntimeQuizQuizResolver]Error getting parent node %s", err)
		return &model.QuizQuiz{}, nil
	}
	vQuiz, err := vQuizParent.GetQuiz(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimequizRuntimeQuizQuizResolver]Error getting Quiz object %s", err)
		return &model.QuizQuiz{}, nil
	}
	dn := vQuiz.DisplayName()
	parentLabels := map[string]interface{}{"quizes.quiz.example.com": dn}
	Labels, _ := json.Marshal(vQuiz.Spec.Labels)
	LabelsData := string(Labels)
	vDefaultScorePerQuestion := int(vQuiz.Spec.DefaultScorePerQuestion)
	vDescription := string(vQuiz.Spec.Description)
	Categories, _ := json.Marshal(vQuiz.Spec.Categories)
	CategoriesData := string(Categories)

	for k, v := range obj.ParentLabels {
		parentLabels[k] = v
	}
	ret := &model.QuizQuiz{
		Id:                      &dn,
		ParentLabels:            parentLabels,
		Labels:                  &LabelsData,
		DefaultScorePerQuestion: &vDefaultScorePerQuestion,
		Description:             &vDescription,
		Categories:              &CategoriesData,
	}
	log.Debugf("[getRuntimequizRuntimeQuizQuizResolver]Output object %v", ret)

	return ret, nil
}

// ////////////////////////////////////
// CHILDREN RESOLVER
// FieldName: Answers Node: RuntimeQuiz PKG: Runtimequiz
// ////////////////////////////////////
func getRuntimequizRuntimeQuizAnswersResolver(obj *model.RuntimequizRuntimeQuiz, id *string) ([]*model.RuntimeanswerRuntimeAnswer, error) {
	log.Debugf("[getRuntimequizRuntimeQuizAnswersResolver]Parent Object %+v", obj)
	var vRuntimeanswerRuntimeAnswerList []*model.RuntimeanswerRuntimeAnswer
	if id != nil && *id != "" {
		log.Debugf("[getRuntimequizRuntimeQuizAnswersResolver]Id %q", *id)
		vRuntimeAnswer, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).Evaluation().Quiz(getParentName(obj.ParentLabels, "runtimequizes.runtimequiz.example.com")).GetAnswers(context.TODO(), *id)
		if err != nil {
			log.Errorf("[getRuntimequizRuntimeQuizAnswersResolver]Error getting Answers node %q : %s", *id, err)
			return vRuntimeanswerRuntimeAnswerList, nil
		}
		dn := vRuntimeAnswer.DisplayName()
		parentLabels := map[string]interface{}{"runtimeanswers.runtimeanswer.example.com": dn}
		vProvidedAnswer := string(vRuntimeAnswer.Spec.ProvidedAnswer)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.RuntimeanswerRuntimeAnswer{
			Id:             &dn,
			ParentLabels:   parentLabels,
			ProvidedAnswer: &vProvidedAnswer,
		}
		vRuntimeanswerRuntimeAnswerList = append(vRuntimeanswerRuntimeAnswerList, ret)

		log.Debugf("[getRuntimequizRuntimeQuizAnswersResolver]Output Answers objects %v", vRuntimeanswerRuntimeAnswerList)

		return vRuntimeanswerRuntimeAnswerList, nil
	}

	log.Debug("[getRuntimequizRuntimeQuizAnswersResolver]Id is empty, process all Answerss")

	vRuntimeAnswerParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).Evaluation().GetQuiz(context.TODO(), getParentName(obj.ParentLabels, "runtimequizes.runtimequiz.example.com"))
	if err != nil {
		log.Errorf("[getRuntimequizRuntimeQuizAnswersResolver]Error getting parent node %s", err)
		return vRuntimeanswerRuntimeAnswerList, nil
	}
	vRuntimeAnswerAllObj, err := vRuntimeAnswerParent.GetAllAnswers(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimequizRuntimeQuizAnswersResolver]Error getting Answers objects %s", err)
		return vRuntimeanswerRuntimeAnswerList, nil
	}
	for _, i := range vRuntimeAnswerAllObj {
		vRuntimeAnswer, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).Evaluation().Quiz(getParentName(obj.ParentLabels, "runtimequizes.runtimequiz.example.com")).GetAnswers(context.TODO(), i.DisplayName())
		if err != nil {
			log.Errorf("[getRuntimequizRuntimeQuizAnswersResolver]Error getting Answers node %q : %s", i.DisplayName(), err)
			continue
		}
		dn := vRuntimeAnswer.DisplayName()
		parentLabels := map[string]interface{}{"runtimeanswers.runtimeanswer.example.com": dn}
		vProvidedAnswer := string(vRuntimeAnswer.Spec.ProvidedAnswer)

		for k, v := range obj.ParentLabels {
			parentLabels[k] = v
		}
		ret := &model.RuntimeanswerRuntimeAnswer{
			Id:             &dn,
			ParentLabels:   parentLabels,
			ProvidedAnswer: &vProvidedAnswer,
		}
		vRuntimeanswerRuntimeAnswerList = append(vRuntimeanswerRuntimeAnswerList, ret)
	}

	log.Debugf("[getRuntimequizRuntimeQuizAnswersResolver]Output Answers objects %v", vRuntimeanswerRuntimeAnswerList)

	return vRuntimeanswerRuntimeAnswerList, nil
}

// ////////////////////////////////////
// LINK RESOLVER
// FieldName: Answer Node: RuntimeAnswer PKG: Runtimeanswer
// ////////////////////////////////////
func getRuntimeanswerRuntimeAnswerAnswerResolver(obj *model.RuntimeanswerRuntimeAnswer) (*model.QuizchoiceQuizChoice, error) {
	log.Debugf("[getRuntimeanswerRuntimeAnswerAnswerResolver]Parent Object %+v", obj)
	vQuizChoiceParent, err := nc.RootRoot().Tenant(getParentName(obj.ParentLabels, "tenants.tenant.example.com")).Runtime().User(getParentName(obj.ParentLabels, "runtimeusers.runtimeuser.example.com")).Evaluation().Quiz(getParentName(obj.ParentLabels, "runtimequizes.runtimequiz.example.com")).GetAnswers(context.TODO(), getParentName(obj.ParentLabels, "runtimeanswers.runtimeanswer.example.com"))
	if err != nil {
		log.Errorf("[getRuntimeanswerRuntimeAnswerAnswerResolver]Error getting parent node %s", err)
		return &model.QuizchoiceQuizChoice{}, nil
	}
	vQuizChoice, err := vQuizChoiceParent.GetAnswer(context.TODO())
	if err != nil {
		log.Errorf("[getRuntimeanswerRuntimeAnswerAnswerResolver]Error getting Answer object %s", err)
		return &model.QuizchoiceQuizChoice{}, nil
	}
	dn := vQuizChoice.DisplayName()
	parentLabels := map[string]interface{}{"quizchoices.quizchoice.example.com": dn}
	vChoice := string(vQuizChoice.Spec.Choice)
	vHint := string(vQuizChoice.Spec.Hint)
	vPictureName := string(vQuizChoice.Spec.PictureName)
	vAnswer := bool(vQuizChoice.Spec.Answer)

	for k, v := range obj.ParentLabels {
		parentLabels[k] = v
	}
	ret := &model.QuizchoiceQuizChoice{
		Id:           &dn,
		ParentLabels: parentLabels,
		Choice:       &vChoice,
		Hint:         &vHint,
		PictureName:  &vPictureName,
		Answer:       &vAnswer,
	}
	log.Debugf("[getRuntimeanswerRuntimeAnswerAnswerResolver]Output object %v", ret)

	return ret, nil
}
