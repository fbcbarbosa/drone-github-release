FROM golang:1.12.4-alpine3.9
WORKDIR /go/src/github.com/drone-plugins/drone-github-release
RUN apk add -U --no-cache ca-certificates git
ADD . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o /tmp/drone-github-release .

FROM scratch
LABEL maintainer="QuintoAndar <github.com/quintoandar>"
COPY --from=0 /tmp/drone-github-release /bin/
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
ENTRYPOINT ["/bin/drone-github-release"]
