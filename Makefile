SHELL:=/bin/sh

all: run_test

run_test:
	@cd algorithms && go test