package repo_impl

import (
	"backend_github_trending/db"
	"backend_github_trending/errMess"
	"backend_github_trending/log"
	"backend_github_trending/model"
	"backend_github_trending/repository"
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type GithubRepoImpl struct {
	sql *db.Sql
}

func NewGitHubRepo(sql *db.Sql) repository.GithubRepo {
	return &GithubRepoImpl{sql: sql}
}

func (g GithubRepoImpl) SelectRepoByName(context context.Context, name string) (model.GithubRepo, error) {
	var repo = model.GithubRepo{}
	err := g.sql.Db.GetContext(context, &repo, `SELECT * FROM repos WHERE name=$1`, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return repo, errMess.RepoNotFound
		}
		log.Error(err.Error())
	}
	return repo, nil
}

func (g GithubRepoImpl) SaveRepo(context context.Context, repo model.GithubRepo) (model.GithubRepo, error) {
	// name, description, url, color, lang, fork, stars, stars_today, build_by, created_at, updated_at
	var statment = `
		INSERT INTO repo(name, description, url, color, lang, fork, stars, stars_today, build_by, created_at, updated_at)
		VALUES(:name,:description, :url, :color, :lang, :fork, :stars, :stars_today, :build_by, :created_at, :updated_at)
	`
	repo.CreatedAt = time.Now()
	repo.UpdatedAt = time.Now()

	_, err := g.sql.Db.NamedExecContext(context, statment, repo)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return repo, errMess.RepoConflict
			}
		}
		log.Error(err.Error())
		return repo, errMess.RepoInsertFail
	}
	return repo, nil
}

func (g GithubRepoImpl) SelectRepos(context context.Context, userId string, limit int) ([]model.GithubRepo, error) {
	var repo []model.GithubRepo
	statment := `
		SELECT repos.name, repos.description, repos.url, repos.color, repos.lang, repos.fork, repos.stars, repos.stars_today, repos.build_by, repos.updated_at, 
		FROM repos
		FULL OUTER JOIN bookmarks
		ON repos.name = bookmarks.repo_name
		AND bookmarks.user_id=$1
		WHERE repos.name IS NOT NULL
		ORDER BY updated_at ASC LIMIT $2
	`
	err := g.sql.Db.SelectContext(context, &repo, statment, userId, limit)
	if err != nil {
		log.Error(err.Error())
		return repo, err
	}
	return repo, nil
}

func (g GithubRepoImpl) UpdateRepo(context context.Context, repo model.GithubRepo) (model.GithubRepo, error) {
	// name, description, url, color, lang, fork, stars, stars_today, build_by, created_at, updated_at
	statment := `
		UPDATE repos
		SET	
			stars = :stars,
			fork = :fork,
			stars_today = :stars_today,
			build_by = :build_by,
			updated_at = :updated_at
		WHERE name = :name
	`
	result, err := g.sql.Db.NamedExecContext(context, statment, repo)
	if err != nil {
		log.Error(err.Error())
		return repo, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return repo, errMess.RepoNotUpdated
	}
	if count == 0 {
		return repo, errMess.RepoNotUpdated
	}
	return repo, nil
}

func (g GithubRepoImpl) SelectAllBookmarks(context context.Context, userId string) ([]model.GithubRepo, error) {
	repos := []model.GithubRepo{}
	statment := `
		SELECT repos.name, repos.description, repos.url, repos.color, repos.lang, repos.fork, repos.stars, repos.stars_today, repos.build_by
		FROM bookmarks
		INNER JOIN repos
		ON bookmarks.user_id=$1
		AND repos.name = bookmarks.name
	`
	err := g.sql.Db.SelectContext(context, &repos, statment, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return repos, errMess.BookmarkNotFound
		}
		log.Error(err.Error())
		return repos, err
	}
	return repos, nil
}

func (g GithubRepoImpl) Bookmark(context context.Context, bid, nameRepo, userId string) error {
	statment := `
		INSERT INTO bookmarks(bid, user_id, repo_name, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5)
	`
	now := time.Now()
	_, err := g.sql.Db.ExecContext(context, statment, bid, userId, nameRepo, now, now)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errMess.BookmarkConflic
			}
		}
		log.Error(err.Error())
		return errMess.BookmarkFail
	}
	return nil
}

func (g GithubRepoImpl) DelBookmark(context context.Context, nameRepo, userId string) error {
	statment := `DELETE FROM bookmarks WHERE repo_name =  $1 AND user_id = $2`
	result := g.sql.Db.MustExecContext(context, statment, nameRepo, userId)

	_, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return errMess.DelBookmarkFail
	}
	return nil
}