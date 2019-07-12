TARGET=tencent_vote

LDFLAGS=-ldflags "-X main.Env=dev -s -w"

default:
	@go build -o ${TARGET} -v ${LDFLAGS} -tags=jsoniter
run:
	@go run -v ${LDFLAGS} -tags=jsoniter  main.go

clean:
	@if [ -f ${TARGET} ] ; then rm ${TARGET} ; fi

.PHONY: default run clean
