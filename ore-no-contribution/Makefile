NAME     := ore-no-contribution
VERSION  := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

$(NAME): $(SRCS)
	go build  $(LDFLAGS) -o $(NAME)

