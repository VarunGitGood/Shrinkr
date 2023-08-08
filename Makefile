build:
	@echo "Building shrinkr..."
	@echo "Generating completion files..."
	@go run main.go completion bash > comp.sh
	@go build -o shrinkr
	@sudo mv shrinkr /usr/local/bin/
	@echo "Almost there..."
	@echo "To enable autocomplete run: source comp.sh"