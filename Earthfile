VERSION 0.6

all:
    BUILD +lint
    BUILD +test

code:
    FROM +golang-base

    # download deps
    COPY go.mod go.sum .
    RUN go mod download -x
    # copy code
    COPY --dir cos .
    COPY --dir cospb .
    COPY --dir cosmocks .

vendor:
    FROM +code

    RUN go mod vendor
    SAVE ARTIFACT /app /files


lint:
    FROM +vendor

    # Runs golangci-lint with settings:
    RUN golangci-lint run --timeout 10m --skip-dirs-use-default

test:
    FROM +vendor

    RUN go test -mod=vendor ./... -race -coverprofile=coverage.out -covermode=atomic

    SAVE ARTIFACT coverage.out AS LOCAL coverage.out

protogen:
    FROM +golang-base

    # copy the proto files to generate
    COPY --dir protos/ ./
    COPY buf.work.yaml buf.gen.yaml ./

    RUN ls ./
    RUN ls protos/

    # generate the pbs
    RUN buf generate \
            --template buf.gen.yaml \
            --path protos/chief-of-state-protos/chief_of_state/v1

    # save artifact to
    SAVE ARTIFACT gen gen AS LOCAL cospb

mock:
    FROM +golang-base

    # download deps
    COPY go.mod go.sum .
    RUN go mod download -x
    # copy code
    COPY --dir cospb .
    # TODO: Turn back on when vektra/mockery fixes generics for interfaces
    COPY --dir cos/client.go .

	# generates chief of state mocks
	RUN mockery --all --output ./cosmocks --case snake

    SAVE ARTIFACT ./cosmocks cosmocks AS LOCAL cosmocks

golang-base:
    FROM golang:1.18.2-alpine

    WORKDIR /app

    # install gcc dependencies into alpine for CGO
    RUN apk add gcc musl-dev curl git openssh

    # install docker tools
    # https://docs.docker.com/engine/install/debian/
    RUN apk add --update --no-cache docker

    # install the go generator plugins
    RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    RUN export PATH="$PATH:$(go env GOPATH)/bin"

    # install vektra/mockery
    RUN go install github.com/vektra/mockery/v2@v2.13.0-beta.1

    # install buf from source
    RUN GO111MODULE=on GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/buf@v1.4.0

    # install linter
    # binary will be $(go env GOPATH)/bin/golangci-lint
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
    RUN ls -la $(which golangci-lint)
