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

build_docker: tailwind
	docker build -t ${DOCKER_CONTAINER} .

run_docker:
	docker run -it -p ${PORT}:${PORT} ${DOCKER_CONTAINER}

stop_docker:
	docker stop ${DOCKER_CONTAINER}

run:
	air
