package app

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kadm"
)

func EnsureKafkaTopics(kcl *kgo.Client, topics ...string) error {
	kadmcl := kadm.NewClient(kcl)

	ctx := context.Background()
	// or something: context.WithTimeout(ctx, time.Second * 30)

	// todo: check functionality
	_, err := kadmcl.CreateTopics(ctx, -1, -1, nil, topics...)
	if err != nil {
		return err
	}

	return nil
}