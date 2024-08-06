postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb-orders:
	docker exec -it postgres createdb --username=root --owner=root orders

dropdb-orders:
	docker exec -it postgres dropdb orders

createdb-payments:
	docker exec -it postgres createdb --username=root --owner=root payments

dropdb-payments:
	docker exec -it postgres dropdb payments

.PHONY: postgres createdb-orders dropdb-orders created-payments dropdb-payments