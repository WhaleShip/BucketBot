ENV_FILE = .env

ENV_VARS = \
    POSTGRES_DB=buckets \
    POSTGRES_USER=user \
    POSTGRES_PASSWORD=pass \
    POSTGRES_HOST=db \
    POSTGRES_PORT=5432 \
	SSL_MODE=disable

env:
	@$(eval SHELL:=/bin/bash)
	@printf "%s\n" $(ENV_VARS) > $(ENV_FILE)
	@echo "$(ENV_FILE) file created"

run:
	@docker compose up --build -d

off:
	@docker compose down