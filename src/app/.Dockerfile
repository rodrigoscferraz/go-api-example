FROM golang:1.21 as build

WORKDIR /app
COPY ./src/app/*  /app
RUN pwd && ls -lah
RUN go mod tidy && \
    go vet -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12:latest
COPY --from=build /go/bin/app /
EXPOSE 8081

ENTRYPOINT ["/app"]