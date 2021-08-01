package schema

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"student-housing-backend/ent"
	"student-housing-backend/graph/generated"
	"student-housing-backend/graph/model"
	"time"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*ent.User, error) {
	num := time.Now().UnixNano()
	fmt.Println("CreateUser" + strconv.FormatInt(num, 10))
	return r.Client.User.Create().SetAge(30).SetNickname(input.Name + strconv.FormatInt(num, 10)).SetName(input.Name).Save(context.Background())
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
