FROM golang:1.16

RUN cd /usr/src && mkdir myapp
ENV GOPATH=/workspaces/docker
ENV PATH=$PATH:/workspaces/docker/bin
EXPOSE 8080:8080
WORKDIR /usr/src/myapp