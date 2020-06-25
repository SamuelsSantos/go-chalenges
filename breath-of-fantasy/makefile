NAME := breath-of-fantasy
breath-of-fantasy:
	echo "Game Breath of Fantasy"
	
build:
	go build -o bin/$(NAME) main.go

test:
	mkdir -p ./coverage
	go test -v -coverprofile=./coverage/tests.out
	go tool cover -html=./coverage/tests.out -o ./coverage/coverage-report.html

run:
	go run main.go


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/$(NAME)-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/$(NAME)-numero-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/$(NAME)-numero-freebsd-386 main.go

all: $(NAME) build