pg_migration_folder = scripts/pgbo/migrations/
database_link = postgres://postgres:403201aa@localhost:5432/postgres?sslmode=disable

gen.sql.pg.migration:
	migrate create -ext sql -dir $(pg_migration_folder) -seq $(name)

pg.migration.up:
	migrate -path $(pg_migration_folder) -database $(database_link) -verbose up

pg.migration.down:
	migrate -path $(pg_migration_folder) -database $(database_link) -verbose down

pg.migration.fix:
	migrate -path $(pg_migration_folder) -database $(database_link) force $(version)

gen.pg.model:
	cd scripts/pgbo && sqlc generate