package app

import (
	"fmt"
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

func StartConsuming(kcl *kgo.Client) {
	ctx := context.Background()

	fmt.Println("starting to consume")
	for {
		fetches := kcl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			panic(fmt.Sprint(errs))
		}

		fetches.EachTopic(func(ft kgo.FetchTopic) {
			switch ft.Topic {
			case "topic1":
				consumeTopic1(ft)
			case "topic2":
				consumeTopic2(ft)
			default:
				panic("recieved fetch from unexpected topic")
			}
		})
	}
}

func consumeTopic1(ft kgo.FetchTopic) {
	fmt.Println("recieved from topic 1")
	ft.EachRecord(func(record *kgo.Record) {
		fmt.Println(string(record.Value))
	})
}

func consumeTopic2(ft kgo.FetchTopic) {
	fmt.Println("recieved from topic 2")
	ft.EachRecord(func(record *kgo.Record) {
		fmt.Println(string(record.Value))
	})
}

/*
base example:
func StartConsuming(kcl *kgo.Client) {
	ctx := context.Background()
	for {
		fetches := kcl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			panic(fmt.Sprint(errs))
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}

		// or a callback function.
		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			for _, record := range p.Records {
				fmt.Println(string(record.Value), "from range inside a callback!")
			}

			// We can even use a second callback!
			p.EachRecord(func(record *kgo.Record) {
				fmt.Println(string(record.Value), "from a second callback!")
			})
		})
	}
}
*/