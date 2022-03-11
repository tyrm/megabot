PROJECT_NAME=megabot

build-snapshot: clean
	goreleaser build --snapshot

clean:
	rm -Rvf coverage.txt dist gosec.xml megabot
	find . -name ".DS_Store" -exec rm -v {} \;

clean-npm:
	rm -Rvf web/static/css/bootstrap.css web/static/js/bootstrap.bundle.js

fmt:
	go fmt ./...

gosec:
	gosec ./...

i18n-extract:
	goi18n extract

lint:
	golint ./...

npm-install:
	cd web/bootstrap && npm install

npm-install-jenkins:
	cd web/bootstrap && npm install --cache=/.npm

npm-scss: clean-npm
	cd web/bootstrap && npm run scss

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