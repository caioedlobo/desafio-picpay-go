# ==================================================================================== #
# HELPERS
# ==================================================================================== #

### help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

### docker/up: start the docker container on docker-compose.yml
.PHONY: docker/up
docker/up:
	start "Docker Desktop" "C:\Program Files\Docker\Docker\Docker Desktop.exe"
	docker-compose up -d

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd --db-dsn=${PICPAY_DB_DSN}

## db/migration/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${PICPAY_DB_DSN} up

## db/migrations/down: rollback all up database migrations
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Rolling back all migrations...'
	migrate -path ./migrations -database ${PICPAY_DB_DSN} down

## db/migrations/goto version=$1: Go to a specific version of database
.PHONY: db/migrations/goto
db/migrations/goto: confirm
	@echo 'Going to version ${version} of database'
	migrate -path ./migrations -database ${PICPAY_DB_DSN} goto ${version}

## db/migrations/force version=$1: force command on migrate
.PHONY: db/migrations/force
db/migrations/force: confirm
	@echo 'Forcing migration to version ${version}'
	migrate -path ./migrations -database ${PICPAY_DB_DSN} force ${version}