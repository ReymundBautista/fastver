FROM golang:stretch
LABEL maintainer="Reymund Bautista"
USER root
ARG PROJECTNAME
ARG PROJECTPATH=/build/${PROJECTNAME}
RUN mkdir -p ${PROJECTPATH}
WORKDIR ${PROJECTPATH}
COPY ./ ${PROJECTPATH}
RUN GOOS=darwin GO111MODULE=on CGO_ENABLED=0 go build -a -ldflags="-s -w"
RUN GOOS=darwin go install
RUN ls -l /go/bin/darwin_amd64
