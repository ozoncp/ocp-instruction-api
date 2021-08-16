
PROJECT_NAME = ocp-instruction-api

DIR_BIN := bin/$(PROJECT_NAME)
DIR_PKG := pkg/$(PROJECT_NAME)
DIR_CMD := cmd/$(PROJECT_NAME)
DIR_VENDOR := vendor.protogen/api/$(PROJECT_NAME)
DIR_API := api/$(PROJECT_NAME)


run:
	go run $(DIR_CMD)/main.go

lint:
	golint ./...

test:
	go test -v ./...

.PHONY: build
build: .vendor-proto .generate .build

.PHONY: gen
gen: .generate

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p $(DIR_PKG)

		#mkdir -p $(DIR_VENDOR)
		#cp $(DIR_API)/$(PROJECT_NAME).proto $(DIR_VENDOR)

		protoc -I vendor.protogen \
				-I $(DIR_API) \
				--go_out=$(DIR_PKG) --go_opt=paths=import \
				--go-grpc_out=$(DIR_PKG) --go-grpc_opt=paths=import \
				--grpc-gateway_out=$(DIR_PKG) --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=import \
				--validate_out lang=go:$(DIR_PKG) \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				$(DIR_API)/$(PROJECT_NAME).proto

		mv $(DIR_PKG)/github.com/ozoncp/$(PROJECT_NAME)/$(DIR_PKG)/* $(DIR_PKG)/
		rm -rf $(DIR_PKG)/github.com
		mkdir -p $(DIR_CMD)

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o $(DIR_BIN) $(DIR_CMD)/main.go

.PHONY: install
install: build .install

.PHONY: .install
.install:
		go install cmd/grpc-server/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi

		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi

.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/envoyproxy/protoc-gen-validate
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/golang/protobuf/proto
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install github.com/envoyproxy/protoc-gen-validate
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/golang/protobuf/protoc-gen-go
