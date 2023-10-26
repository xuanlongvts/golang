Migrate:
    migrate create -ext sql -dir ./migrations -seq add_address_col_table_company

### run migration
    make migra --step=1