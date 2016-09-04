from openjdk:jre
maintainer Cory Buecker <email@corybuecker.com>

run wget http://apache.claz.org/zookeeper/zookeeper-3.4.8/zookeeper-3.4.8.tar.gz -O /tmp/zookeeper-3.4.8.tar.gz && \
    tar -zxvf /tmp/zookeeper-3.4.8.tar.gz -C /opt && \
    mv /opt/zookeeper-3.4.8 /opt/zookeeper && \
    rm -rf /tmp/zoo*

copy ["zoo.cfg", "/opt/zookeeper/conf/zoo.cfg"]

expose 2181
entrypoint ["/opt/zookeeper/bin/zkServer.sh", "start-foreground"]
