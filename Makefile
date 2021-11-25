# Build all
all: deps
	@ echo "Building..."
	go build ./src

run: deps
	@ echo "Running..."
	- go run ./src

# Get repo dependencies
deps:
	@ echo "Checking dependencies..."
	which go	
	go get -d ./... 

