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

package main

import (
	"bytes"
	"net/http"
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

func applyPostgresOutbox(cfg *InitConfig, conf *config.PostgresConfig) {
	payloadB, err := json.Marshal(map[string]string{
		"connector.class": "io.debezium.connector.postgresql.PostgresConnector",
		"plugin.name": "pgoutput",
		"tasks.max": "1",
		"table.include.list": "public.event_outbox",
		"transforms": "outbox",
		"transforms.outbox.type": "io.debezium.transforms.outbox.EventRouter",
		"transforms.outbox.route.topic.replacement": "${routedByValue}",
		"value.converter": "io.debezium.converters.BinaryDataConverter",
		
		"topic.prefix": cfg.ServiceName,
		"database.hostname": conf.Host,
		"database.port": conf.Port,
		"database.user": conf.Username,
		"database.password": conf.Password,
		"database.dbname": conf.Database,
	})
	log.Info().Bytes("payload", payloadB).Msg("payload bytes")
	log.Info().Any("abc", cfg.ServiceName).Msg("a")
	log.Info().Any("conf", conf).Msg("b")
	log.Info().Any("cfg", cfg).Msg("c")
	if err != nil {
		log.Panic().Err(err).Msg("debezium connect: failed to marshal debezium cfg")
	}
	
	url := cfg.DebeziumHost + "/connectors/" + cfg.ServiceName + "-outbox/config"
	log.Info().Str("url", url).Msg("debezium url")
	req, err := http.NewRequest(
		"PUT",
		url,
		bytes.NewReader(payloadB),
	)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Panic().Err(err).Msg("debezium connect: failed to make debezium request")
	}

	log.Info().Str("status", res.Status).Msg("debezium connect: applied outbox config")
}