GO_VERSION := 1.21  # <1>



.PHONY: install-go init-go

setup: install-go init-go # <2>

install-go: # <3>
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go: # <4>
    echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
    echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

upgrade-go: # <5>
	sudo rm -rf /usr/bin/go
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

migrateup:
	migrate -path=./pkg/migrations -database="postgresql://bersen:bersen@localhost/filmdb?sslmode=disable" up

migratedown:
	migrate -path=./pkg/migrations -database="postgresql://bersen:bersen@localhost/filmdb" down 

build:
	go build -o api cmd/rest/main.go


postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=bersen -e POSTGRES_PASSWORD=bersen -d postgres:latest

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)
