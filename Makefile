

all: build

build:
	@echo "Building UpdateDNS...."
	@go build -o bin/updatedns src/main.go
	@echo "All done! 😁"

clean:
	@rm bin/*
	@echo "All Clean! 😊"
