PWD := $(PWD)
export PATH := $(PWD)/bin:$(PATH)
export GOFLAGS := -mod=mod

lint:
	golangci-lint run --config ./.golangci.yml ./...

mock: mock
	find ./internal -type d -name "*mock*" | xargs rm -dfR
	find ./pkg -type d -name "*mock*" | xargs rm -dfR
	@PATH="${ROOTDIR}/bin:${PATH}:${GOPATH}" go generate ./...
	find . -type d -name "*mock*" -exec sh -c 'cd "$$1" && touch .coverignore && cd -' sh {} \;

test:
	go test ./... -count=1 -race
