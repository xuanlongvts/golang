package errMess

import "errors"

var (
	//repo
	RepoNotUpdated = errors.New("Update information Repo failed")
	RepoNotFound   = errors.New("Repo not found")
	RepoConflict   = errors.New("Repo has exist")
	RepoInsertFail = errors.New("Add Repo failed")

	//bookmark
	BookmarkNotFound = errors.New("Bookmark not found")
	BookmarkFail = errors.New("Bookmark failed")
	DelBookmarkFail = errors.New("DelBookmark failed")
	BookmarkConflic = errors.New("Bookmark has exist")

	//genneral
	ErrorSql       = errors.New("Error SQL")
)
