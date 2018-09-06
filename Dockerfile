FROM scratch

COPY go-env-server-amd64 /go-env-server

EXPOSE 8888

CMD ["/go-env-server"]
