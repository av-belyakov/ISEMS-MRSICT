FROM golang:1.17-alpine AS temporary_image
WORKDIR /go/src/
ENV PATH /usr/local/go/bin:$PATH
RUN apk update && \
    apk add --no-cache git && \
    git clone https://github.com/av-belyakov/ISEMS-MRSICT.git /go/src/
RUN go build

FROM alpine
LABEL author="Belaykov Artemy" application="ISEMS-MRSICT"
WORKDIR /opt/isems-mrsict
ENV MY_KEYS=~/_isems_docker_containers/isems-mrsict/keys
RUN mkdir /opt/isems-mrsict/defaultsettingsfiles && \
    mkdir /home/logs && \
    mkdir /home/keys
COPY --from=temporary_image /go/src/ISEMS-MRSICT /opt/isems-mrsict/ 
COPY --from=temporary_image /go/src/defaultsettingsfiles/* /opt/isems-mrsict/defaultsettingsfiles/
#COPY keys/* /home/keys/
COPY $MY_KEYS/* /home/keys/
COPY config.json README.md /opt/isems-mrsict/
EXPOSE 13000
ENTRYPOINT [ "./ISEMS-MRSICT" ]