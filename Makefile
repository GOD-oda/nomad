.PHONY: build
build:
	docker compose build --no-cache --force-rm

.PHONY: down
down:
	docker compose down -v --remove-orphans
