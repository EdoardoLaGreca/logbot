# Build all
all: build


build: deps
	@ echo "Building..."
	go build -o bin/logbot ./src

run: deps
	@ echo "Running..."
	- go run ./src

# Get repo dependencies
deps:
	@ echo "Checking dependencies..."
	which go	
	go get -d ./... 

