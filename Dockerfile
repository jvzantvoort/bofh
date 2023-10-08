FROM golang:1.21

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v . && go install ./...

EXPOSE 2609
CMD ["bofh"]
