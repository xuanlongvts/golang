MAIN_FILE=main.go

dev:
	go run ${MAIN_FILE}

# ---------- Start DB
DB_FILE=db/migration
DB_CONN=postgresql://postgres:123456@localhost:5432/db-pooling?sslmode=disable

mig-new-file:
	migrate create -ext sql -dir "${DB_FILE}" -seq init_schema

mig-up-force:
	migrate -source "file://${DB_FILE}" -database "${DB_CONN}" force 1

mig-up:
	migrate -source "file://${DB_FILE}" -database "${DB_CONN}" up 1

mig-down-force:
	migrate -source "file://${DB_FILE}" -database "${DB_CONN}" force 1

mig-down:
	migrate -source "file://${DB_FILE}" -database "${DB_CONN}" down 1
# ///---------- End DB

