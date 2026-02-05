run:
	@cd src ; go run *.go

build:
	@cd src ; go mod tidy ; go build -o ../bin/httpd *.go

image:
	@docker build -t web:latest .

start:
	@docker run -d -p 8888:8888 web:latest

stop:
	@docker rm -f httpd