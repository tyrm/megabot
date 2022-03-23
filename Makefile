PROJECT_NAME=megabot

.DEFAULT_GOAL := test

build-snapshot: clean
	goreleaser build --snapshot

clean:
	rm -Rvf coverage.txt dist gosec.xml megabot
	find . -name ".DS_Store" -exec rm -v {} \;

clean-npm:
	rm -Rvf web/bootstrap/dist web/static/css/bootstrap.min.css web/static/js/bootstrap.bundle.min.js

fmt:
	go fmt ./...

gosec:
	gosec ./...

i18n-extract:
	goi18n extract -outdir locales

i18n-merge:
	goi18n merge -outdir locales locales/active.*.toml locales/translate.*.toml

i18n-translations:
	goi18n merge -outdir locales locales/active.*.toml

lint:
	golint ./...

minify-static:
	minify web/static-src/css/error.css > web/static/css/error.min.css
	minify web/static-src/css/login.css > web/static/css/login.min.css
	minify web/bootstrap/dist/bootstrap.css > web/static/css/bootstrap.min.css
	minify web/bootstrap/node_modules/bootstrap/dist/js/bootstrap.bundle.js > web/static/js/bootstrap.bundle.min.js

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

test: tidy fmt lint #gosec
	MB_TLS_CERT=../../server.crt MB_TLS_KEY=../../server.key go test -cover ./...

test-ext: tidy fmt lint #gosec
	MB_TLS_CERT=../../server.crt MB_TLS_KEY=../../server.key go test --tags=postgres -cover ./...

test-race: tidy fmt lint #gosec
	MB_TLS_CERT=../../server.crt MB_TLS_KEY=../../server.key go test -race -cover ./...

test-race-ext: tidy fmt lint #gosec
	MB_TLS_CERT=../../server.crt MB_TLS_KEY=../../server.key go test --tags=postgres -race -count=1 -cover ./...

test-verbose: tidy fmt lint #gosec
	MB_TLS_CERT=../../server.crt MB_TLS_KEY=../../server.key go test -v -cover ./...

tidy:
	go mod tidy -compat=1.18

.PHONY: fmt lint test test-ext