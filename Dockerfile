FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" \
    CGO_ENABLED=0 \
    go build -o main .

# execution phase
FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/main /app
EXPOSE 8080
ENTRYPOINT [ "./main" ]
