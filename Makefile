PROJECT_NAME=megabot

build-snapshot:
	goreleaser build --snapshot --single-target --rm-dist

clean:
	rm -Rvf dist megabot

fmt:
	go fmt ./...

gosec:
	gosec ./...

lint:
	golint ./...

test-docker-restart: test-docker-stop test-docker-start

test-docker-start:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml up -d

test-docker-stop:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml down

test: tidy fmt lint gosec
	go test -cover ./...

test-race: tidy fmt lint gosec
	go test -race -cover ./...

test-verbose: tidy fmt lint gosec
	go test -v -cover ./...

tidy:
	go mod tidy -compat=1.17

.PHONY: fmt lint test