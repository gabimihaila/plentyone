all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/microservice microservice/main.go
	docker buildx build --tag microservice -f Dockerfile_microservice .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/client client/main.go
	docker buildx build --tag client -f Dockerfile_client .
	cd api_gateway && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ../bin/api_gateway main.go
	docker buildx build --tag api_gateway -f Dockerfile_api_gateway .

clean:
	rm bin/*
