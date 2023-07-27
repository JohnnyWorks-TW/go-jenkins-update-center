FROM golang:1.20.3-alpine3.17 as builder
ENV CGO_ENABLED=0
WORKDIR /data
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:3.17
WORKDIR /app
COPY --chown=1000:1000 --from=builder /data/go-jenkins-update-center .

RUN apk --no-cache add ca-certificates curl vim \
  && adduser -u 2000 -s /bin/bash -D iuser \
  && chown -R iuser:iuser /app \
  && rm -rf /var/cache/apk/*

USER iuser
EXPOSE 8080
CMD ["./go-jenkins-update-center"]