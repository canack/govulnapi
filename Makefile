IMAGE_TAG=localhost/govulnapi

build:
	docker build -t ${IMAGE_TAG} .

run:
	docker run --rm -it -p 127.0.0.1:8080:8080 -p 127.0.0.1:8081:8081 ${IMAGE_TAG}

swagger:
	swag init --pd -g cmd/govulnapi/main.go -o api/docs
