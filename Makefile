NAME=mssql
VERSION=0.0.2
BINARY=${NAME}_${VERSION}

default: build

build:
	go build -o ${BINARY}

release:
	mkdir -p ${NAME}/{VERSION}/linux_amd64
	mkdir -p ${NAME}/{VERSION}/windows_amd64
	mkdir -p ${NAME}/{VERSION}/darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o ${BINARY}
	mv ${BINARY} ${NAME}/{VERSION}/linux_amd64
	GOOS=windows GOARCH=amd64 go build -o ${BINARY}
	mv ${BINARY} ${NAME}/{VERSION}/windows_amd64
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY}
	mv ${BINARY} ${NAME}/{VERSION}/darwin_amd64
