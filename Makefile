CURRENT_DIR=$(shell pwd)

start:
	 go run cmd/main.go

compile:
		./scripts/proto-gen.sh ${CURRENT_DIR}

.PHONY:compile
