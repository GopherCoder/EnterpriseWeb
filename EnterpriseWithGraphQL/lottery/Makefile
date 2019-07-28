PROJECT=lottery

default:
	go build -o ${PROJECT} main.go


run:
	@go run main.go


sync: default
	@chmod +x ${PROJECT}
	./${PROJECT} sync

import: default
	@chmod +x ${PROJECT}
	./${PROJECT} winner


.PHONY: default run sync import
