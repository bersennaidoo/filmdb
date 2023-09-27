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

build:
	go build -o rest cmd/rest/main.go

migrateup:
	migrate -path=./pkg/migrations -database="postgresql://bersen:bersen@localhost/filmdb" up

migratedown:
	migrate -path=./pkg/migrations -database="postgresql://bersen:bersen@localhost/filmdb" down 
