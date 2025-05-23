build:
	@docker compose build code

bash:
	@docker compose run --remove-orphans --service-ports code bash

# **************************************************

.PHONY: build bash
