FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/kg/codonOptimizer/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/poly_optimization_app

FROM scratch
COPY --from=builder /go/bin/poly_optimization_app /poly_optimization_app
ENTRYPOINT ["/poly_optimization_app"]
EXPOSE 8080
