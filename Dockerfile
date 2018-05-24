FROM golang as compiler
RUN go get github.com/davidkarim/go_hello_world

FROM scratch
COPY --from=compiler /go/bin/go_hello_world .
EXPOSE 8080
CMD ["./go_hello_world"]
