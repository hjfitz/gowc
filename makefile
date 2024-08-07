# Set the default target
all: build

# Build the Go binary
build:
	go build -o wc main.go

# Run the built binary
run:
	./wc

# Clean up the built binary
clean:
	rm -f wc
