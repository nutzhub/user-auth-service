FROM golang:1.4.2
RUN apt-get update && apt-get install --no-install-recommends -y -q build-essential git

# golang puts its go install here (weird but true)
ENV GOROOT_BOOTSTRAP /usr/src/go

# golang sets GOPATH=/go
ADD . /go/src/userauth
RUN go install userauth
ENTRYPOINT ["/go/bin/userauth"]

# Kubernetes expects us to listen on port 8080
EXPOSE 8080