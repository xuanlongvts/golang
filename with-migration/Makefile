migra-create: # make migra-create name=int_schema
	migrate create -ext sql -dir ./migrations -seq $(name)

migra: # make migra step=1
	@echo $(step)
	go run main.go --step=$(step)

go-mod:
	go mod tidy
	go mod verify
	go mod vendor


