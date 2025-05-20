build:
	@docker compose build statch

bash:
	@docker compose run --remove-orphans --service-ports statch bash

# **************************************************

.PHONY: build bash
