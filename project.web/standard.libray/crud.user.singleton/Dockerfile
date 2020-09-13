# tart by building the application.
# Build em api.user com distroless
FROM golang:1.13.4 as builder

WORKDIR /go/src/api.user

COPY . .

ENV GO111MODULE=on

#RUN go install -v ./...
#RUN GOOS=linux go  build -ldflags="-s -w" -o api.user main.go
RUN cp api.user /go/bin/api.user

RUN ls -lh

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/api.user /
CMD ["/api.user"]