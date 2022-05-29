## Description
Small hello world app with minimal requiements. Sends in a Hello, world message to specified topic. 

## Requirements
- pem file
- kafka keytab file
- kafka krb5.conf file
- kafka servers
- kafka principal
- kafka topic
		

## Env Variables
- bootstrap.servers (which servers are we connecting to(producer or consumer))
- sasl.kerberos.principal (what principal are we using)
- topic (which topic are we sending a message to)
- KRB5_CONFIG (krb5 config location)

## build the container
`docker build -t test_kafka_hello_world .`

## run the container

`docker run -v [local path to credentials]:/credentials -e bootstrap.servers=[producing server] -e sasl.kerberos.principal=[principal] -e topic=[topic] -e KRB5_CONFIG=[path inside containter to krb5.conf]  test_kafka_hello_world`

## example
example of a running container with all env variables populated 

*Note: Code assumes authentication files are mounted in /credentials*

`docker run -v /opt/FONPatrol/kafka-gateway/truststore.pem:/credentials/truststore.pem -v /opt/FONPatrol/kafka-gateway/svc_prd_patrol_kafka.keytab:/credentials/kafka.keytab -v /opt/FONPatrol/kafka-gateway/krb5.conf:/credentials/krb5.conf -e bootstrap.servers=stgpldbfx0006.unix.gsm1900.org:9093 -e sasl.kerberos.principal=svc_prd_patrol_kafka@DBFX.STAGE.POL.CDH.T-MOBILE.COM -e topic=oneconsole.sprint.patrol.network.alarm -e KRB5_CONFIG=/credentials/krb5.conf  test_kafka_hello_world`