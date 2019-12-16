package errMess

import "errors"

var (
	UserConflict = errors.New("The user has exist")
	SingupFailed = errors.New("Signup failed")
	UserNotFound = errors.New("User not found")
)
