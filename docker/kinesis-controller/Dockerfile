FROM bitnami/minideb:jessie

RUN install_packages ca-certificates

ADD kinesis-controller /kinesis-controller

ENTRYPOINT ["/kinesis-controller"]
