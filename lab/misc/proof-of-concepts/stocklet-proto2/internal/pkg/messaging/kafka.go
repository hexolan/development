// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package messaging

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kadm"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

func NewKafkaConn(conf *config.KafkaConfig, opts ...kgo.Opt) (*kgo.Client, error) {
	opts = append(opts, kgo.SeedBrokers(conf.Brokers...))
	kCl, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to connect to Kafka", err)
	}

	return kCl, nil
}

func EnsureKafkaTopics(kcl *kgo.Client, topics ...string) error {
	ctx := context.Background()
	kadmCl := kadm.NewClient(kcl)

	_, err := kadmCl.CreateTopics(ctx, -1, -1, nil, topics...)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to create Kafka topics", err)
	}

	return nil
}
