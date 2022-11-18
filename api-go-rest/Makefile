fix:
	export GO111MODULE="on" 

run:
	go run main.go

init:
	go mod init github.com/ioannova/alura-golang/api-go-rest

install:
	go get -u github.com/gorilla/mux
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/gorilla/handlers

ip:
	docker inspect api-go-rest_postgres_1 | grep IPAddress