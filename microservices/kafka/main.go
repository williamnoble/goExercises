package main

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"log"
	"sync"
)

// franz-go with redpanda

var (
	fooTopic      = "fooTopic"
	asyncFooTopic = "asyncFooTopic"
)

func main() {
	seeds := []string{"127.0.0.1:50070"}

	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		//kgo.ConsumerGroup("std-group"),
		kgo.ConsumeTopics(fooTopic, asyncFooTopic),
	)

	if err != nil {
		log.Fatal("failed to connect to Kafka Client", err)
	}

	defer client.Close()
	ctx := context.Background()

	// 1. Producer: sync
	var wg sync.WaitGroup
	wg.Add(1)
	record := &kgo.Record{
		Value: []byte("bar"),
		Topic: fooTopic,
	}
	client.Produce(ctx, record, func(r *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			log.Printf("record produced an error: %v\n\n", err)
		}
		fmt.Printf("Record: %s\n", r.Value)
	})
	wg.Wait()

	// 2. Producer: async
	var records = []*kgo.Record{
		{
			Value: []byte("sync-value-1"),
			Topic: asyncFooTopic,
		},
		{
			Value: []byte("sync-value-2"),
			Topic: asyncFooTopic},
	}
	if err := client.ProduceSync(ctx, records...).FirstErr(); err != nil {
		fmt.Printf("record produced an errro with async: %v\n", err)
	}

	//3.

	for {

		fetches := client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Fatalf(fmt.Sprint("err when fetching: ", errs))
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), " from an iterator!")
		}

		fetches.EachPartition(func(fetchTopicPartition kgo.FetchTopicPartition) {
			for _, record := range fetchTopicPartition.Records {
				fmt.Println(string(record.Value), "from range inside a callback")
			}

			fetchTopicPartition.EachRecord(func(record *kgo.Record) {
				fmt.Println(string(record.Value), "from a second callback")
			})
		})

	}

	fmt.Println("END OF MAIN")
}
