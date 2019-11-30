package resolvers

import (
	"context"

	"github.com/cmelgarejo/go-gql-server/internal/gql"
	"github.com/cmelgarejo/go-gql-server/internal/gql/models"
)

type Resolver struct {}

type queryResolver struct {
	*Resolver
}

type mutationResolver struct {
	*Resolver
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}