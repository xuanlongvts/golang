MAIN_FILE=main.go

dev:
	go run ${MAIN_FILE}

# ---------- Start DB
DB_FILE=db/migration
DB_CONN=postgresql://postgres:123456@localhost:5432/db-pooling?sslmode=disable

define run_migrate
	migrate -source "file://${DB_FILE}" -database "${DB_CONN}" $(1) $(2)
endef

mig-new:
	migrate create -ext sql -dir "${DB_FILE}" -seq init_schema

mig-force:
	$(call run_migrate,force,1)

mig-up: # make mig-up 1
	$(call run_migrate,up,1)

mig-down: # make mig-down 1
	$(call run_migrate,down,1)
# ///---------- End DB

# benkmark
benk-normal:
	~/go/bin/go-wrk -c 200 -d 10 http://localhost:8080/products/normal

benk-pooled:
	~/go/bin/go-wrk -c 200 -d 10 http://localhost:8080/products/pooled

