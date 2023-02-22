.PHONY: build
build:
	docker compose build --no-cache --force-rm

.PHONY: down
down:
	docker compose down -v --remove-orphans

.PHONY: readme
readme:
	docker compose run --rm app go run main.go
	@cp src/README.md README.md
