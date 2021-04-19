GOPATH:=$(shell go env GOPATH)

.PHONY: db\:drop
db\:drop:
	go run cmd/db_drop/main.go

.PHONY: db\:seed
db\:seed:
	go run cmd/db_seed/main.go

.PHONY: test
test:
	go test ./... -cover -count=1 -v
