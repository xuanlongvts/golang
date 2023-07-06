//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import "github.com/google/wire"

func InitializeEvent(pharse string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
