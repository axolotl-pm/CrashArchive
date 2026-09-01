FROM --platform=$BUILDPLATFORM golang:1.27-trixie AS build
ARG TARGETARCH
ENV CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH
WORKDIR /go/crasharchive
COPY . .
RUN go build -o ./bin/crasharchive ./cmd/crasharchive
RUN go build -o ./bin/crasharchive-adduser ./cmd/crasharchive-adduser

FROM debian:trixie-slim

WORKDIR /app
COPY --from=build /go/crasharchive/bin/* ./

CMD ["./crasharchive"]
