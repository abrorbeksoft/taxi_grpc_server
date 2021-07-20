CURRENT_DIR=$(shell pwd)

run:
	echo "Hello Abrorbek"

compile:
		./scripts/proto-gen.sh ${CURRENT_DIR}

.PHONY:compile
