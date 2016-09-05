package messages

import (
	"log"
	"runtime"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s: %s", name, elapsed)
}

func GenerateRawMessages(numberOfMessages int) [][]byte {
	defer timeTrack(time.Now(), "building the messages took")

	log.Printf("numberOfMessages is %d", numberOfMessages)

	var mem runtime.MemStats

	messages := make([]Message, numberOfMessages)
	rawMessages := make([][]byte, numberOfMessages)

	for i, message := range messages {
		data, _ := message.randomMessageAsBytes()
		rawMessages[i] = data
	}

	runtime.ReadMemStats(&mem)
	log.Printf("building the messages used %d bytes of memory", mem.Alloc)

	return rawMessages
}
