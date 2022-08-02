FROM golang:1.18 as builder

ENV GO111MODULE=on
ENV APP_NAME echo_app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o $APP_NAME ./cmd

FROM scratch

COPY --from=builder /app/echo_app /echo_app
EXPOSE 8080

CMD [ "./echo_app" ]
