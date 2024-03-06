// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package outboxie

import (
	"context"

	"github.com/hexolan/outboxie/pkg/database"
	"github.com/hexolan/outboxie/pkg/messaging"
)

type Processor interface {
	Start() error
	Stop() error
}

// temp: definitelyt refactor

// todo:
// create with rough outlines and segregate the useful abstractions after
// like the database ReceiverFunc is called when DB sees new event msg
// there will need to be another callback for "event succesfully dispatched" or "dispatch failed"

type ProcessorConfig struct {
	DatabaseDriver database.Driver
	MessagingDriver messaging.Driver

	DatabaseUrl string
	MessagingUrl string
}

type OutboxProcessor struct {
	ctx context.Context
	ctxCancel context.CancelFunc

	bridgeFunc database.ReceiverCallback
	
	dbDriver database.Driver
	msgDriver messaging.Driver
}

func NewProcessor(cfg ProcessorConfig) (Processor, error) {
	ctx, ctxCancel := context.WithCancel(context.Background())

	dbDriver, err := cfg.DatabaseDriver.Setup(ctx, cfg.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	msgDriver, err := cfg.MessagingDriver.Setup(ctx, cfg.MessagingUrl)
	if err != nil {
		return nil, err
	}

	// todo: checks for electioner
	// msg processing func
	
	op := OutboxProcessor{
		ctx: ctx,
		ctxCancel: ctxCancel,

		dbDriver: dbDriver,
		msgDriver: msgDriver,
	}
	op.bridgeFunc = defaultBridgeFunc(op)

	return &op, nil
}

func defaultBridgeFunc(op OutboxProcessor) database.ReceiverFunc {
	// todo: adding electioner wrapper
	return func(topic string, msg []byte) {
		// todo: catching error and whatever
		op.msgDriver.Dispatch(topic, msg)
	}
}

func (op *OutboxProcessor) Start() error {
	ctx := context.Background()


	return nil
}

func (op *OutboxProcessor) Stop() error {
	return nil
}