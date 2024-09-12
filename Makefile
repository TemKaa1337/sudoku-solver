ifneq ("$(wildcard ./.env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif

.PHONY: setup snapshot tests login

docker-up:
	@docker compose build
	@docker compose up --detach

docker-go-cli:
	@docker compose exec app zsh

docker-down:
	@docker compose down
