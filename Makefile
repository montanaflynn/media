.PHONY: all

default: test lint

format: 
	go fmt .

test:
	go test -race 
	
check: format test

benchmark:
	go test -bench=. -benchmem

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html="coverage.out"

lint: format
	golangci-lint run .

docs:
	godoc2md github.com/montanaflynn/media | sed -e s#src/target/##g > DOC.md

changelog:
	git-chglog -o CHANGELOG.md
