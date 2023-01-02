POSTGRES_USER := $(shell printenv POSTGRES_USER)
POSTGRES_PASSWORD := $(shell printenv POSTGRES_PASSWORD)
POSTGRES_DB := $(shell printenv POSTGRES_DB)

start-dev:
	air ./cmd/main.go -b 0.0.0.0

migrate-create:
	migrate create -ext sql -dir migrations -seq $(table)

migrate-up:
	migrate -database postgres://root:root@postgres:5432/DS?sslmode=disable -path ./migrations up ${v}

migrate-down:
	migrate -database postgres://root:root@postgres:5432/DS?sslmode=disable -path ./migrations down ${a}

# парсинг из всех источниов
# make parse-adreses
parse-adreses:
	go run ./cmd/parsers/adreses/main.go

# make parse-stores
parse-stores:
	go run ./cmd/parsers/stores/main.go

