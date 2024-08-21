# step 1:	export DEBEZIUM_VERSION=2.7

run:
	docker compose -f docker-compose-mysql.yml up

regis-mysql:
	curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d @register-mysql.json