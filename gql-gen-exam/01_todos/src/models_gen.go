// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package src

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
