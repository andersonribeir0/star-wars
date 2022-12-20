.PHONY: build
build:
	@REPO_NAME=http://github.com/andersonribeir0/starfields docker build -f ./infra/Dockerfile .