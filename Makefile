# Variables
ENV_FILE=.env

.PHONY: all clean build

all: build

# Build the Docker image
build: $(REQUIREMENTS)
	docker-compose build

# Utility target to run the Docker container
run: $(ENV_FILE)
	docker-compose up
