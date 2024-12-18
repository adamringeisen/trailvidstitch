# Define the program name
PROGRAM = trailvidstitch

# Define local bin directory
LOCAL_BIN = ./bin

# Define install directory
INSTALL_DIR = /usr/local/bin

# Go compiler flags (optional)
GOFLAGS = -v

# Default target (builds locally)
all: build

# Build locally (to ./bin)
build:
	mkdir -p $(LOCAL_BIN)
	go build $(GOFLAGS) -o $(LOCAL_BIN)/$(PROGRAM)

# Install to /usr/local/bin (using the local binary)
install: build
	sudo install -m 0755 $(LOCAL_BIN)/$(PROGRAM) $(INSTALL_DIR)

# Uninstall from /usr/local/bin
uninstall:
	sudo rm -f $(INSTALL_DIR)/$(PROGRAM)

# Clean up local build
clean:
	rm -rf $(LOCAL_BIN)

#Format the code
fmt:
	go fmt ./...
