NAME=mssql
VERSION=0.0.1
BINARY=terraform-provider-${NAME}_${VERSION}

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=linux GOARCH=amd64 go install

install: build
	mv ${BINARY} ~/.terraform.d/plugins/${BINARY}

