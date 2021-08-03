package schema

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"student-housing-backend/ent"
	"student-housing-backend/graph/generated"
	"student-housing-backend/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.Client.User.Query().All(context.Background())
}

func (r *queryResolver) Properties(ctx context.Context) ([]*model.Property, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
