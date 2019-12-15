package repo_impl

import (
	"backend_github_trending/db"
	"backend_github_trending/log"
	"backend_github_trending/model"
	"backend_github_trending/errMess"
	"backend_github_trending/model/req"
	"backend_github_trending/repository"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{sql:sql}
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errMess.UserConflict
			}
			return user, errMess.SingupFailed
		}
	}

	return user, nil
}

func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqSignin) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errMess.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}