.PHONY: tests
tests:
	go test -v ./week1

.DEFAULT_GOAL := tests