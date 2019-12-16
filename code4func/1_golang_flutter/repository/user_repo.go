package repository

import (
	"backend_github_trending/model"
	"backend_github_trending/model/req"
	"context"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignin) (model.User, error)
}
