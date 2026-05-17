.PHONY: dev build setup tailwind templ

setup:
	go mod tidy
	sqlc generate
	templ generate
	cd ./internal/web && bun install && bun run build:css

dev:
	make -j2 tailwind templ

tailwind:
	cd ./internal/web && bun run tailwind

templ:
	templ generate --watch --proxy="http://localhost:8081" --cmd="go run ."

build:
	templ generate && cd ./internal/web && bunx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css && cd - && go build -o ./bin/app .
