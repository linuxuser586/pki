# 1.12.6-stretch
FROM golang@sha256:35200a727dc44175d9221a6ece398eed7e4b8e17cb7f0d72b20bf2d5cf9dc00d AS build

WORKDIR /go/src/github.com/linuxuser586/pki

COPY . .
RUN go build cmd/pki/pki.go

# 2019-07-05
FROM linuxuser586/base@sha256:d0a7ebd7d97b29d34311aa0c3f4f1d7367ee7c6fc55f128208d77bce8d1376f7

COPY --from=build /go/src/github.com/linuxuser586/pki/pki /

EXPOSE 10042

ENTRYPOINT [ "dumb-init", "--", "/pki" ]