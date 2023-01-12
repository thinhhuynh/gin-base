FROM golang:1.19-alpine3.17

# Required because go requires gcc to build
RUN apk add build-base
RUN apk add inotify-tools
RUN apk add git
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN echo $GOPATH

ENV APP_PATH='/app'

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

RUN go mod download

CMD sh ${APP_PATH}/docker/run.sh