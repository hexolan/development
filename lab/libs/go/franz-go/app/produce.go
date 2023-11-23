package app

import (
	"fmt"
	"sync"
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

func ProduceRecords(kcl *kgo.Client) {
	fmt.Println("starting to produce records")

	ctx := context.Background()

	record := &kgo.Record{Topic: "topic1", Value: []byte("foo")}
	if err := kcl.ProduceSync(ctx, record).FirstErr(); err != nil {
		fmt.Printf("record had a produce error while synchronously producing: %v\n", err)
	}


	record = &kgo.Record{Topic: "topic2", Value: []byte("bar")}
	var wg sync.WaitGroup
	wg.Add(1)
	kcl.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
	wg.Wait()

	fmt.Println("records produced")
}