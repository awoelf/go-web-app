tailwind:
	bun run tailwindcss --config tailwind.config.js -i index.css -o public/css/styles.css
	echo "Generating tailwindcss files..." \

build: tailwind
	go build -o bin ./cmd/server/main.go

build_docker: tailwind
	docker build -t awoelf/go-web-app .

run_docker:
	docker run -it -p 3000:3000 awoelf/go-web-app

stop_docker:
	docker stop awoelf/go-web-app

run: tailwind
	air