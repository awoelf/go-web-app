include .env

tailwind:
	@bun run tailwindcss --config tailwind.config.js -i index.css -o public/css/styles.css
	@echo "Generating tailwindcss files..."

build: tailwind
	@if [ -f "./tmp/main" ]; then \
		echo "Deleting old build file"; \
		rm -rf ./tmp/main; \
	fi 
	@go build -o ./tmp/ main.go;
	@echo "Building app..." 

build_docker: tailwind
	docker build -t awoelf/go-web-app .

run_docker:
	docker run -it -p ${PORT}:${PORT} awoelf/go-web-app

stop_docker:
	docker stop awoelf/go-web-app

run:
	air
