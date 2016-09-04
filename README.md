# Building Zookeeper

    docker build -t zookeeper zookeeper

# Running Zookeeper with Volume Persistence

    docker run -d -p 2181:2181 --name zookeeper --restart always -v zookeeper:/var/zookeeper zookeeper

# Building Kafka

    docker build -t kafka kafka

# Running Kafka with Volume Persistence

    docker run -d -v kafka:/var/kafka -p 9092:9092 --link zookeeper:zookeeper --name kafka --restart always kafka --override zookeeper.connect=zookeeper:2181
