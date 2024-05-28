package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/build/nexus-gql/graph/generated"
	"example/build/nexus-gql/graph/model"
)

// Root is the resolver for the root field.
func (r *queryResolver) Root(ctx context.Context) (*model.RootRoot, error) {
	return getRootResolver()
}

// User is the resolver for the User field.
func (r *config_ConfigResolver) User(ctx context.Context, obj *model.ConfigConfig, id *string) ([]*model.UserUser, error) {
	return getConfigConfigUserResolver(obj, id)
}

// Event is the resolver for the Event field.
func (r *config_ConfigResolver) Event(ctx context.Context, obj *model.ConfigConfig, id *string) ([]*model.EventEvent, error) {
	return getConfigConfigEventResolver(obj, id)
}

// Quiz is the resolver for the Quiz field.
func (r *evaluation_EvaluationResolver) Quiz(ctx context.Context, obj *model.EvaluationEvaluation, id *string) ([]*model.QuizQuiz, error) {
	return getEvaluationEvaluationQuizResolver(obj, id)
}

// Question is the resolver for the Question field.
func (r *quiz_QuizResolver) Question(ctx context.Context, obj *model.QuizQuiz, id *string) ([]*model.QuizquestionQuizQuestion, error) {
	return getQuizQuizQuestionResolver(obj, id)
}

// Choice is the resolver for the Choice field.
func (r *quizquestion_QuizQuestionResolver) Choice(ctx context.Context, obj *model.QuizquestionQuizQuestion, id *string) ([]*model.QuizchoiceQuizChoice, error) {
	return getQuizquestionQuizQuestionChoiceResolver(obj, id)
}

// Tenant is the resolver for the Tenant field.
func (r *root_RootResolver) Tenant(ctx context.Context, obj *model.RootRoot, id *string) ([]*model.TenantTenant, error) {
	return getRootRootTenantResolver(obj, id)
}

// Evaluation is the resolver for the Evaluation field.
func (r *root_RootResolver) Evaluation(ctx context.Context, obj *model.RootRoot) (*model.EvaluationEvaluation, error) {
	return getRootRootEvaluationResolver(obj)
}

// User is the resolver for the User field.
func (r *runtime_RuntimeResolver) User(ctx context.Context, obj *model.RuntimeRuntime, id *string) ([]*model.RuntimeuserRuntimeUser, error) {
	return getRuntimeRuntimeUserResolver(obj, id)
}

// Answer is the resolver for the Answer field.
func (r *runtimeanswer_RuntimeAnswerResolver) Answer(ctx context.Context, obj *model.RuntimeanswerRuntimeAnswer) (*model.QuizchoiceQuizChoice, error) {
	return getRuntimeanswerRuntimeAnswerAnswerResolver(obj)
}

// Quiz is the resolver for the Quiz field.
func (r *runtimeevaluation_RuntimeEvaluationResolver) Quiz(ctx context.Context, obj *model.RuntimeevaluationRuntimeEvaluation, id *string) ([]*model.RuntimequizRuntimeQuiz, error) {
	return getRuntimeevaluationRuntimeEvaluationQuizResolver(obj, id)
}

// Quiz is the resolver for the Quiz field.
func (r *runtimequiz_RuntimeQuizResolver) Quiz(ctx context.Context, obj *model.RuntimequizRuntimeQuiz) (*model.QuizQuiz, error) {
	return getRuntimequizRuntimeQuizQuizResolver(obj)
}

// Answers is the resolver for the Answers field.
func (r *runtimequiz_RuntimeQuizResolver) Answers(ctx context.Context, obj *model.RuntimequizRuntimeQuiz, id *string) ([]*model.RuntimeanswerRuntimeAnswer, error) {
	return getRuntimequizRuntimeQuizAnswersResolver(obj, id)
}

// User is the resolver for the User field.
func (r *runtimeuser_RuntimeUserResolver) User(ctx context.Context, obj *model.RuntimeuserRuntimeUser) (*model.UserUser, error) {
	return getRuntimeuserRuntimeUserUserResolver(obj)
}

// Evaluation is the resolver for the Evaluation field.
func (r *runtimeuser_RuntimeUserResolver) Evaluation(ctx context.Context, obj *model.RuntimeuserRuntimeUser) (*model.RuntimeevaluationRuntimeEvaluation, error) {
	return getRuntimeuserRuntimeUserEvaluationResolver(obj)
}

// Interest is the resolver for the Interest field.
func (r *tenant_TenantResolver) Interest(ctx context.Context, obj *model.TenantTenant, id *string) ([]*model.InterestInterest, error) {
	return getTenantTenantInterestResolver(obj, id)
}

// Config is the resolver for the Config field.
func (r *tenant_TenantResolver) Config(ctx context.Context, obj *model.TenantTenant) (*model.ConfigConfig, error) {
	return getTenantTenantConfigResolver(obj)
}

// Runtime is the resolver for the Runtime field.
func (r *tenant_TenantResolver) Runtime(ctx context.Context, obj *model.TenantTenant) (*model.RuntimeRuntime, error) {
	return getTenantTenantRuntimeResolver(obj)
}

// Wanna is the resolver for the Wanna field.
func (r *user_UserResolver) Wanna(ctx context.Context, obj *model.UserUser, id *string) ([]*model.WannaWanna, error) {
	return getUserUserWannaResolver(obj, id)
}

// Interest is the resolver for the Interest field.
func (r *wanna_WannaResolver) Interest(ctx context.Context, obj *model.WannaWanna) (*model.InterestInterest, error) {
	return getWannaWannaInterestResolver(obj)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Config_Config returns generated.Config_ConfigResolver implementation.
func (r *Resolver) Config_Config() generated.Config_ConfigResolver { return &config_ConfigResolver{r} }

// Evaluation_Evaluation returns generated.Evaluation_EvaluationResolver implementation.
func (r *Resolver) Evaluation_Evaluation() generated.Evaluation_EvaluationResolver {
	return &evaluation_EvaluationResolver{r}
}

// Quiz_Quiz returns generated.Quiz_QuizResolver implementation.
func (r *Resolver) Quiz_Quiz() generated.Quiz_QuizResolver { return &quiz_QuizResolver{r} }

// Quizquestion_QuizQuestion returns generated.Quizquestion_QuizQuestionResolver implementation.
func (r *Resolver) Quizquestion_QuizQuestion() generated.Quizquestion_QuizQuestionResolver {
	return &quizquestion_QuizQuestionResolver{r}
}

// Root_Root returns generated.Root_RootResolver implementation.
func (r *Resolver) Root_Root() generated.Root_RootResolver { return &root_RootResolver{r} }

// Runtime_Runtime returns generated.Runtime_RuntimeResolver implementation.
func (r *Resolver) Runtime_Runtime() generated.Runtime_RuntimeResolver {
	return &runtime_RuntimeResolver{r}
}

// Runtimeanswer_RuntimeAnswer returns generated.Runtimeanswer_RuntimeAnswerResolver implementation.
func (r *Resolver) Runtimeanswer_RuntimeAnswer() generated.Runtimeanswer_RuntimeAnswerResolver {
	return &runtimeanswer_RuntimeAnswerResolver{r}
}

// Runtimeevaluation_RuntimeEvaluation returns generated.Runtimeevaluation_RuntimeEvaluationResolver implementation.
func (r *Resolver) Runtimeevaluation_RuntimeEvaluation() generated.Runtimeevaluation_RuntimeEvaluationResolver {
	return &runtimeevaluation_RuntimeEvaluationResolver{r}
}

// Runtimequiz_RuntimeQuiz returns generated.Runtimequiz_RuntimeQuizResolver implementation.
func (r *Resolver) Runtimequiz_RuntimeQuiz() generated.Runtimequiz_RuntimeQuizResolver {
	return &runtimequiz_RuntimeQuizResolver{r}
}

// Runtimeuser_RuntimeUser returns generated.Runtimeuser_RuntimeUserResolver implementation.
func (r *Resolver) Runtimeuser_RuntimeUser() generated.Runtimeuser_RuntimeUserResolver {
	return &runtimeuser_RuntimeUserResolver{r}
}

// Tenant_Tenant returns generated.Tenant_TenantResolver implementation.
func (r *Resolver) Tenant_Tenant() generated.Tenant_TenantResolver { return &tenant_TenantResolver{r} }

// User_User returns generated.User_UserResolver implementation.
func (r *Resolver) User_User() generated.User_UserResolver { return &user_UserResolver{r} }

// Wanna_Wanna returns generated.Wanna_WannaResolver implementation.
func (r *Resolver) Wanna_Wanna() generated.Wanna_WannaResolver { return &wanna_WannaResolver{r} }

type queryResolver struct{ *Resolver }
type config_ConfigResolver struct{ *Resolver }
type evaluation_EvaluationResolver struct{ *Resolver }
type quiz_QuizResolver struct{ *Resolver }
type quizquestion_QuizQuestionResolver struct{ *Resolver }
type root_RootResolver struct{ *Resolver }
type runtime_RuntimeResolver struct{ *Resolver }
type runtimeanswer_RuntimeAnswerResolver struct{ *Resolver }
type runtimeevaluation_RuntimeEvaluationResolver struct{ *Resolver }
type runtimequiz_RuntimeQuizResolver struct{ *Resolver }
type runtimeuser_RuntimeUserResolver struct{ *Resolver }
type tenant_TenantResolver struct{ *Resolver }
type user_UserResolver struct{ *Resolver }
type wanna_WannaResolver struct{ *Resolver }
