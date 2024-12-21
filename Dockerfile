FROM golang:1.21.3-alpine

RUN apk add --no-cache ca-certificates

COPY ./backend/bin/poller /home/poller-application/poller
COPY ./frontend/dist/ /home/poller-application/dist/

WORKDIR /home/poller-application
ENTRYPOINT ["/home/poller-application/poller"]