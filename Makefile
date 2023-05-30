#!/bin/bash

migrate:
	@go run main.go db:migrate

seed:
	@go run main.go db:seed

user-list:
	@go run main.go user:list

user-create:
	@go run main.go user:create

user-update:
	@go run main.go user:update

user-delete:
	@go run main.go user:delete

user-detail:
	@go run main.go user:detail