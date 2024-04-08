

all: build

build:
	@echo "Building UpdateDNS...."
	@go build -o bin/nm_updatedns main.go
	@echo "All done! 😁"

run:
	@./bin/nm_updatedns

clean:
	@rm bin/*
	@echo "All Clean! 😊"
