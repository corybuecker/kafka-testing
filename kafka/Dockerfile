from openjdk:jre
maintainer cory buecker <email@corybuecker.com>

run wget http://apache.claz.org/kafka/0.9.0.1/kafka_2.11-0.9.0.1.tgz -O /tmp/kafka_2.11-0.9.0.1.tgz && \
    tar zxvf /tmp/kafka_2.11-0.9.0.1.tgz -C /opt/ && \
    mv /opt/kafka_2.11-0.9.0.1 /opt/kafka && \
    rm -rf /tmp/kafka*

entrypoint ["/opt/kafka/bin/kafka-server-start.sh", "/opt/kafka/config/server.properties"]
