default:
	@go run 'cmd/tccapp/main.go'

air:
	@air --build.cmd "go build -o out/tccapp cmd/tccapp/main.go" --build.bin "./out/tccapp"
