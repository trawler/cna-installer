FROM golang:alpine as BUILD_IMAGE
MAINTAINER "Petr Ruzicka <ruzickap@mirantis.com>"

RUN apk add --update --no-cache curl git make unzip

RUN git clone https://github.com/trawler/cna-installer.git && \
    cd cna-installer && \
    make build

####################################################################

FROM alpine
COPY --from=BUILD_IMAGE /go/cna-installer/build/bin/* /
ENTRYPOINT ["/cna-installer"]
