include .env

tailwind:
	@bun run tailwindcss --config tailwind.config.js -i index.css -o ${TAILWIND_CSS}
	@echo "Generating tailwindcss files..."

build: tailwind
	@if [ -f "${BINARY}" ]; then \
		echo "Deleting old build file"; \
		rm -rf ${BINARY}; \
	fi 
	@go build -o ${BINARY} main.go;
	@echo "Building app..." 

docker_build: tailwind
	docker build -t ${DOCKER_CONTAINER} .

docker_push:
	docker image push awoelf/go-web-app:latest

docker_run:
	docker run -it -p ${PORT}:${PORT} ${DOCKER_CONTAINER}

docker_stop:
	docker stop ${DOCKER_CONTAINER}

run:
	air