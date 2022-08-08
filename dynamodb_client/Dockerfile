FROM ubuntu:20.04

RUN apt update && mkdir /secret-volume

COPY ./config ./
COPY ./dynamodbclient ./
COPY ./aws_signing_helper /bin/


ENV AWS_SDK_LOAD_CONFIG=1
ENV AWS_CONFIG_FILE=/config
CMD [ "/bin/bash" ]
