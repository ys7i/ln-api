include .env

.PHONY: gen-api
gen-api:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	oapi-codegen --config api/config.yaml ./api/openapi.yaml

.PHONY: migrate-up
migrate-up:
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -database "mysql://$(DB_USER):$(DB_PS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -path db up $C
	
.PHONY: migrate-down
migrate-down:
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -database "mysql://$(DB_USER):$(DB_PS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -path db down $C

.PHONY: migrate-force
migrate-force:
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -database "mysql://$(DB_USER):$(DB_PS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -path db force $V

.PHONY: gen-xo
gen-xo:
	go install github.com/xo/xo@42b11c7999bc6ac5be620949723f44bd0ec63e02
	xo schema  --out=./db/daocore "mysql://$(DB_USER):$(DB_PS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)"