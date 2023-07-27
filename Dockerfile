FROM golang:latest as go

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app
RUN chmod +x ./app

FROM scratch

COPY --from=go /go/src/app /app
ENTRYPOINT [ "./app" ]
