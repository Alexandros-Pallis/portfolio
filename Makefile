container:
	docker compose -f docker-compose.yml down && \
	docker compose --project-name portfolio -f docker-compose.yml up -d --build
templ:
	cd src && templ generate --watch
