build:
	@go build main/execute.go

run: build
	@execute.exe

test:
	@go test exercises/exercise1_test.go