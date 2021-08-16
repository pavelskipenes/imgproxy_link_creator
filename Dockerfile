FROM alpine

RUN apk update

RUN apk add go git

RUN git clone https://github.com/pavelskipenes/imgproxy_link_creator

WORKDIR /imgproxy_link_creator

RUN go build .

ENV SERVER="http://192.168.60.4:9001"
ENV KEY="943b421c9eb07c830af81030552c86009268de4e532ba2ee2eab8247c6da0881"
ENV SALT="520f986b998545b4785e0defbc4f3c1203f22de2374a3d53cb7a7fe9fea309c5"
ENV EXTENSION="jpg"

ENTRYPOINT ["./imgproxy_link_creator", "--path"]