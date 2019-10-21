CURDIR   := $(shell pwd)
GO       := GO111MODULE=on go
GOBUILD  := CGO_ENABLED=0 $(GO) build
GOTEST   := CGO_ENABLED=1 $(GO) test

.PHONY: build chaos-agent chaos-master agent-cli master-cli

build: chaos-agent chaos-master agent-cli master-cli

chaos-agent:
	$(GOBUILD) -o bin/chaos-agent ./cmd/chaos-agent

chaos-master:
	$(GOBUILD) -o bin/chaos-master ./cmd/chaos-master

agent-cli:
	$(GOBUILD) -o bin/agent-cli ./cmd/tool/chaos-agent

master-cli:
	$(GOBUILD) -o bin/master-cli ./cmd/tool/chaos-master
