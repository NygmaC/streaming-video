package consumer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/consumer"
	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
	"github.com/NygmaC/streamming-video/stream-reader/internal/reader"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var proccessConsumer *kafka.Consumer

func Init() {
	// {"videoName":"teste1", "session":"aaaaaa", "connection": {}}
	proccessConsumer = consumer.CreateConsumer("process-group", []string{"video-stream-proccess"})

	start()
}

func start() {
	run()

}

func run() {
	fmt.Println("Consumer OK")

	for {
		ev := proccessConsumer.Poll(3000)

		switch e := ev.(type) {
		case *kafka.Message:

			var streamProccess = model.Proccess{}

			parse(e.Value, &streamProccess)

			go reader.Proccess(streamProccess)

		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)

			//proccessConsumer.Close()

		}
	}
}

func parse(value []byte, p *model.Proccess) {
	err := json.Unmarshal(value, p)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%% Error convert Proccess data: %v\n", err)
	}

}
