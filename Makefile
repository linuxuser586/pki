.PHONY: build
build:
	docker build -t linuxuser586/pki .

.PHONY: test
test:
	go test -race ./...

.PHONY: test-fast
test-fast:
	go test ./...

.PHONY: cover
cover:
	go test -cover ./...

.PHONY: cover-html
cover-html:
	@mkdir -p build
	go test -coverprofile=build/coverage.out ./...
	@go tool cover -html=build/coverage.out -o build/coverage.html
	@firefox build/coverage.html


.PHONY: release
release: build
	docker push linuxuser586/pki
