build:
	@go build main/execute.go

run: build
	@execute.exe