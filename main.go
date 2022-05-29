package main

import (
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

//Start the patrolEvents Producer
func main() {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":          os.Getenv("bootstrap.servers"),
			"ssl.ca.location":            "/credentials/truststore.pem",
			"sasl.kerberos.keytab":       "/credentials/kafka.keytab",
			"sasl.kerberos.principal":    os.Getenv("sasl.kerberos.principal"),
			"sasl.kerberos.service.name": "kafka",
			"security.protocol":          "SASL_SSL",
			"sasl.mechanism":             "GSSAPI",
			"debug":                      "all",
		},
	)

	if err != nil {
		logrus.Errorf("Failed to create Kafka Producer: %s", err)
		os.Exit(1)
	}

	wg := sync.WaitGroup{}
	go handleEventResponses(p, &wg)

	topic := os.Getenv("topic")

	wg.Add(1)
	go func() {
		p.ProduceChannel() <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte("Hello, Kafka!"),
		}
	}()
	wg.Wait()
}

func handleEventResponses(p *kafka.Producer, wg *sync.WaitGroup) {
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				logrus.Errorf("Delivery failed: %v", ev.TopicPartition.Error)
				continue
			}

			logrus.Infof("Delivered message for topic %s [%d] at offset %v Metadata[%v]",
				*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset, ev.TopicPartition.Metadata)

		default:
			logrus.Errorf("Could not understand status of event: %s", ev)
		}
		wg.Done()
	}
}
