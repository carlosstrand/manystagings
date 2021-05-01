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

.PHONY: build
build:
	export NODE_ENV=production && \
	export ZEPTO_ENV=production && \
	rm -rf public/build && \
	npm run build && \
	rm -rf build && mkdir build && \
	go build -o build/app-service *.go &&\
	cp -r templates public ./build
