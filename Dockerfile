FROM golang:1.10-alpine AS builder

# update latest version
RUN apk add git \
    && go get -u github.com/golang/dep/cmd/dep

# create folder repository
RUN mkdir -p "$GOPATH/src/github.com/azharprabudi/api-plastik"

# set working directory
WORKDIR /go/src/github.com/azharprabudi/api-plastik

# set volume
COPY . /go/src/github.com/azharprabudi/api-plastik

# install dep
RUN dep ensure --update

# build docker
RUN go build .

# runner
FROM golang:1.10-alpine AS runner

# set workdir
WORKDIR /app

# copy build file to runner
COPY --from=builder /go/src/github.com/azharprabudi/api-plastik/api-plastik .

CMD [ "./api-plastik" ]

