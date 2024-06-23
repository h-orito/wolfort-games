package graphql

import (
	"context"
	"errors"
	"wolfort-games/adaptor/auth"
	"wolfort-games/middleware/graph"

	"github.com/99designs/gqlgen/graphql"
)

var Directive graph.DirectiveRoot = graph.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	user := auth.GetUser(ctx)
	if user == nil {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
