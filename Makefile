ENV_FILE = .env

ENV_VARS = \
    POSTGRES_DB=buckets \
    POSTGRES_USER=user \
    POSTGRES_PASSWORD=pass \
    POSTGRES_HOST=db \
    POSTGRES_PORT=5432 \
	SSL_MODE=disable \
	PGBOUNCER_HOST=telebot_pgbouncer \
	PGBOUNCER_PORT=6432

env:
	@$(eval SHELL:=/bin/bash)
	@printf "%s\n" $(ENV_VARS) > $(ENV_FILE)
	@echo "$(ENV_FILE) file created"

run:
	@chmod +x scripts/pgbouncer/entrypoint.sh
	@docker compose up --build -d

off:
	@docker compose down