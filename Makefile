all: run

build:
	go build -o ./bin/crasharchive ./cmd/crasharchive
	go build -o ./bin/crasharchive-adduser ./cmd/crasharchive-adduser

run: build
	./bin/crasharchive

cli/mysql:
	docker compose exec db mysql -p -D crash_archive

adduser:
	docker compose exec ca /app/crasharchive-adduser $(ARGS)

defaultconfig:
	cp ./default-docker-compose.yml ./docker-compose.yml
	cp ./config/default-config.json ./config/config.json
