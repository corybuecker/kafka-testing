# Building Zookeeper

    docker build -t zookeeper zookeeper

# Running Zookeeper with Volume Persistence

    docker run -d -p 2181:2181 --name zookeeper -v zookeeper:/var/zookeeper zookeeper

# Building Kafka

    docker build -t kafka kafka

# Running Two Kafka Brokers with Volume Persistence

The advertised listener and zookeeper IP should be the IP address of your Docker host system.

    docker run -d -v kafka-1:/var/kafka -p 9092:9092 --name kafka-1 kafka --override advertised.listeners=PLAINTEXT://192.168.0.9:9092 --override zookeeper.connect=192.168.0.9:2181
    docker run -d -v kafka-2:/var/kafka -p 9093:9092 --name kafka-2 kafka --override advertised.listeners=PLAINTEXT://192.168.0.9:9093 --override zookeeper.connect=192.168.0.9:2181

# Running the Go-based producer

    cd go-kafka-producer
    docker build -t go-kafka-producer .
    docker run -ti --link kafka:kafka go-kafka-producer --kafka_host localhost:9093,localhost:9092 --number_of_messages 25000
