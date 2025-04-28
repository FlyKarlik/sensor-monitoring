include .env

.PHONY: prepare
prepare:
	go mod tidy

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o ./sensor-monitoring ./cmd/sensor-monitoring/main.go
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o ./migrator ./cmd/migrator/main.go

.PHONY: run-sensor-monitoring
run-sensor-monitoring:
	./sensor-monitoring

.PHONY: run-migrator-up
run-migrator-up:
	./migrator up

.PHONY: run-migrator-down
run-migrator-down:
	./migrator down

.PHONY: clean
clean:
	rm -f ./sensor-monitoring
	rm -f ./migrator

.PHONY: migrate_all_up migrate_all_down migrate_force migrate_version migrate_up migrate_down
migrate_all_up:
	migrate -database ${INFRA__POSTGRES__CONN_STR} -path migrations/ up

migrate_all_down:
	migrate -database ${INFRA__POSTGRES__CONN_STR} -path migrations/ down

migrate_force:
	migrate -database ${INFRA__POSTGRES__CONN_STR} -path migrations/ force 1

migrate_version:
	migrate -database ${INFRA__POSTGRES__CONN_STR} -path migrations/ version

migrate_up:
	migrate -database ${INFRA__POSTGRES__CONN_STR} -path migrations/ up 1

migrate_down:
	migrate -database ${INFRA__POSTGRES__CONN_STR} -path migrations/ down 1
