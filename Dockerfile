FROM golang:1.14.2

RUN apt-get update -y
#install neccesairy ssl packages and modules for kerberos support. Install kinit and other kafka utilities.
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y libssl-dev libsasl2-dev libsasl2-modules-gssapi-mit krb5-user

#github.com/confluentinc/confluent-kafka-go/kafka does not include ssl libraries. Need to install librdkafka from source with ssl support
WORKDIR /opt
RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR /opt/librdkafka
RUN ./configure
RUN make
RUN make install
RUN ldconfig

RUN mkdir -p /go/src/hellokafka

ADD . /go/src/hellokafka
WORKDIR /go/src/hellokafka

#build and install go with dynamic tag. Required for ssl libraries to load properly
RUN go build -v -tags dynamic
RUN go install -v -tags dynamic

CMD ["hellokafka"]
