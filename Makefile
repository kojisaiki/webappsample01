NAME     := webappsample01
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')

bin/$(NAME): $(SRCS)
		statik.exe -src=public
		go build -o ./bin/$(NAME).exe
