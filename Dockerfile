FROM golang:1.10-alpine

ENV APP_DIR $GOPATH/src/github.com/nenadstojanovikj/ghremind

RUN apk update \
    && apk add curl git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && go get github.com/canthefason/go-watcher \
    && go install github.com/canthefason/go-watcher/cmd/watcher

COPY . ${APP_DIR}

WORKDIR ${APP_DIR}

RUN dep ensure

ARG GITHUB_TOKEN=${GITHUB_TOKEN}

CMD watcher -run github.com/nenadstojanovikj/ghremind
