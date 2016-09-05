from golang:1.7

copy ["src", "/go/src/"] . /go/src/kafka-producer

run go install github.com/corybuecker/go-kafka-producer

entrypoint ["/go/bin/go-kafka-producer"]
