include .env

tailwind:
	bun run tailwindcss --config tailwind.config.js -i index.css -o public/css/styles.css
	echo "Generating tailwindcss files..." \

build: tailwind
	docker build -t awoelf/go-web-app .

run_docker:
	docker run -it -p ${PORT}:${PORT} awoelf/go-web-app

stop_docker:
	docker stop awoelf/go-web-app

run: tailwind
	air
