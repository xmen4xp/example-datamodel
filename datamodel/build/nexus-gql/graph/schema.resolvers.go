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

// Tenant is the resolver for the Tenant field.
func (r *root_RootResolver) Tenant(ctx context.Context, obj *model.RootRoot, id *string) ([]*model.TenantTenant, error) {
	return getRootRootTenantResolver(obj, id)
}

// Interest is the resolver for the Interest field.
func (r *tenant_TenantResolver) Interest(ctx context.Context, obj *model.TenantTenant, id *string) ([]*model.InterestInterest, error) {
	return getTenantTenantInterestResolver(obj, id)
}

// Config is the resolver for the Config field.
func (r *tenant_TenantResolver) Config(ctx context.Context, obj *model.TenantTenant) (*model.ConfigConfig, error) {
	return getTenantTenantConfigResolver(obj)
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

// Root_Root returns generated.Root_RootResolver implementation.
func (r *Resolver) Root_Root() generated.Root_RootResolver { return &root_RootResolver{r} }

// Tenant_Tenant returns generated.Tenant_TenantResolver implementation.
func (r *Resolver) Tenant_Tenant() generated.Tenant_TenantResolver { return &tenant_TenantResolver{r} }

// User_User returns generated.User_UserResolver implementation.
func (r *Resolver) User_User() generated.User_UserResolver { return &user_UserResolver{r} }

// Wanna_Wanna returns generated.Wanna_WannaResolver implementation.
func (r *Resolver) Wanna_Wanna() generated.Wanna_WannaResolver { return &wanna_WannaResolver{r} }

type queryResolver struct{ *Resolver }
type config_ConfigResolver struct{ *Resolver }
type root_RootResolver struct{ *Resolver }
type tenant_TenantResolver struct{ *Resolver }
type user_UserResolver struct{ *Resolver }
type wanna_WannaResolver struct{ *Resolver }
