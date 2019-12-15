package repository

import (
	"backend_github_trending/model"
	"context"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}