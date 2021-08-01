package schema

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"student-housing-backend/ent"
	"student-housing-backend/graph/generated"
)

func (r *userResolver) Friend(ctx context.Context, obj *ent.User) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
