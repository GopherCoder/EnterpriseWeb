BINARY=target_notes

VERSION=0.0.1

PRODUCTION_FLAGS=-ldflags "-X target_notes.Env=production -s -w"
DEV_FLAGS=-ldflags "-X target_notes.Env=dev -s -w"
TEST_FLAGS=-ldflags "-X target_notes.Env=test -s -w"


default:
	go build -o ${BINARY} -v ${DEV_FLAGS} -tags=jsoniter
prod:
	go build -o ${BINARY} -v ${PRODUCTION_FLAGS} -tags=jsoniter
test:
	go build -o ${BINARY} -v ${TEST_FLAGS} -tags=jsoniter

.PHONY: default prod test
