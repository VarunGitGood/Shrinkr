build:
	@echo "Building shrinkr..."
	@echo "Generating completion files..."
	@go build -o shrinkr
	@sudo mv shrinkr /usr/local/bin/
	@echo "All done"