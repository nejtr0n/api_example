.PHONY: di proto mocks test all
uid=`id -u`
gid=`id -g`
ifdef CI_BUILD_REF_NAME
	APP_VERSION=$(CI_BUILD_REF_NAME)
else
	APP_VERSION=$(shell git symbolic-ref HEAD | sed -e 's,.*/\(.*\),\1,')
endif

ifdef CI_COMMIT_SHA
	APP_REVISION=$(CI_COMMIT_SHA)
else
	APP_REVISION=$(shell git rev-parse --short HEAD)
endif

start:
	@UID=$(uid) GID=$(gid) docker-compose up -d
	@echo "Starting docker-compose environment"

stop:
	@UID=$(uid) GID=$(gid) docker-compose down
	@echo "Started docker-compose environment"

di:
	@docker-compose exec -u $(uid):$(gid) -w /app/src/cmd/server -e GOCACHE=/go/.cache -e CGO_ENABLED=0 builder wire
	@echo "Generated di container"

proto:
	@docker-compose exec -u $(uid):$(gid) -e CGO_ENABLED=0 -e GOCACHE=/go/.cache builder protoc -I /app/src/proto --gofast_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:/app/src/ui/grpc api_example.proto
	@echo "Generated protobuf"

mocks:
	@docker run -it --rm -u $(uid):$(gid) -v `pwd`:/app -w="/app/src" vektra/mockery:v2.4
	@echo "Generated mocks"

test:
	@docker-compose exec -u $(uid):$(gid) -w /app/src -e CGO_ENABLED=0 -e GOCACHE=/go/.cache builder sh -c 'go clean -testcache && go test `go list ./... | grep -v /vendor/`'

server:
	@docker-compose exec -u $(uid):$(gid) -w /app/src -e CGO_ENABLED=0 -e GOOS=linux -e GOCACHE=/go/.cache builder sh -c "go build -a \
-ldflags '-X main.version=${APP_VERSION} -X main.revision=${APP_REVISION} -s -w -extldflags "-static"' \
-installsuffix cgo -o bin/server cmd/server/main.go cmd/server/wire_gen.go"
	@echo "Server client builded successfully"

fetch_client:
	@docker-compose exec -u $(uid):$(gid) -w /app/src -e CGO_ENABLED=0 -e GOOS=linux -e GOCACHE=/go/.cache builder sh -c "go build -a \
-ldflags '-X main.version=${APP_VERSION} -X main.revision=${APP_REVISION} -s -w -extldflags "-static"' \
-installsuffix cgo -o bin/fetch_client cmd/fetch_client/main.go"
	@echo "Fetch client builded successfully"

list_client:
	@docker-compose exec -u $(uid):$(gid) -w /app/src -e CGO_ENABLED=0 -e GOOS=linux -e GOCACHE=/go/.cache builder sh -c "go build -a \
-ldflags '-X main.version=${APP_VERSION} -X main.revision=${APP_REVISION} -s -w -extldflags "-static"' \
-installsuffix cgo -o bin/list_client cmd/list_client/main.go"
	@echo "List client builded successfully"

clients: fetch_client list_client

all: di proto mocks test clients