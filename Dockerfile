FROM golang:1.23-bullseye AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
COPY Makefile ./
RUN make deps
COPY . ./
RUN make build
RUN chmod +x start.sh

FROM alpine:3.15
RUN apk --update add ca-certificates && rm -rf /var/cache/apk/*
RUN adduser -D appuser
USER appuser
COPY --from=builder /usr/src/app/bin/search-keyword-service /home/appuser/search-keyword-service
COPY --from=builder /usr/src/app/start.sh /home/appuser/start.sh
WORKDIR /home/appuser/
CMD ["./start.sh"]