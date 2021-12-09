# Build all
all: build


build: deps
	@ echo "Building..."
	go build -o bin/logbot .

run: deps
	@ echo "Running..."
	- go run .

clean:
	rm -rf bin

update:
	@ echo "Updating repository..."
	git pull

# DO NOT CALL DIRECTLY THE TARGETS BELOW, they exist just to be called from other targets

deps:
	@ echo "Checking dependencies..."
	which go	
	go version
	go get -d ./...	
