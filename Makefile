# Build all
all: deps
	@ echo "Building..."
	go build .

run: deps
	@ echo "Running..."
	go run .

# Get repo dependencies
deps:
	@ echo "Checking dependencies..."
	which go	
	go get -d ./... 

