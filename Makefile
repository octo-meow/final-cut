debug: 
	@dlv debug --headless --listen=:2345 --api-version=2 cmd/main.go

.PHONY: clean all
run: 
	@go run cmd/main.go

watch:
	@tail -f log.txt -n 10
