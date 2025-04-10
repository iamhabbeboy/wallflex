# Makefile
SCHEDULER=scheduler/runner.go
SCHEDULER_BIN=wallflex_scheduler
TMP_PATH=scheduler/wallflex_scheduler
TARGET_PATH=/usr/local/bin/$(SCHEDULER_BIN)

build-scheduler:
	@echo "Building worker..."
	go build -o $(TMP_PATH) $(SCHEDULER)
	# sudo cp $(TMP_PATH) $(TARGET_PATH)
	#sudo chmod +x $(TARGET_PATH)

build-wails:
	build-scheduler
	wails build

release-app:
	wails build -nsis

# Build both the worker and the desktop app
build: build-scheduler build-wails

# Run both the worker and the desktop app
run:
	@echo "Running desktop app with background worker..."
	# Run worker in the background
	./$(SCHEDULER_BIN) & \
	# Run desktop app
	wails dev

# Clean up build artifacts
clean:
	@echo "Cleaning up build artifacts..."
	rm -f $(SCHEDULER_BIN)

lint:
	@echo "Linting..."
	golangci-lint run
