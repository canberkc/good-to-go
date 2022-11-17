FROM golang:1.19-bullseye as builder

ENV GO111MODULE=on
ENV APP_NAME echo_app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY configs ./configs
COPY pkg ./pkg

RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o $APP_NAME ./cmd

FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/echo_app /echo_app
COPY --from=builder /app/configs /configs
EXPOSE 8080

CMD [ "./echo_app" ]
