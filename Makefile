

GO = docker run --rm -e GOOS=linux -e GOARCH=amd64 -v $(PWD):/usr/src/app -w /usr/src/app golang:1.14 go

.PHONY: test


test:
	@echo "Testing..."
	$(GO) test ./pkg/java_util_concurrent/...