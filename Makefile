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
run/api: docker/up
	go run ./cmd/api --db-dsn=${PICPAY_DB_DSN}

## db/migration/new name=$1: create a new database migration
.PHONY: db/migrations/up
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${PICPAY_DB_DSN}