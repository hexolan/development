package main

import (
	"bytes"
	"net/http"
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/hexolan/stocklet/internal/pkg/config"
)

func main() {
	
}

func configurePostgresOutbox(svcName string, dbzHost string, conf config.PostgresConfig) {
	payloadB, err := json.Marshal(map[string]string{
		"connector.class": "io.debezium.connector.postgresql.PostgresConnector",
		"plugin.name": "pgoutput",
		"tasks.max": "1",
		"table.include.list": "public.event_outbox",
		"transforms": "outbox",
		"transforms.outbox.type": "io.debezium.transforms.outbox.EventRouter",
		"transforms.outbox.route.topic.replacement": "${routedByValue}",
		"value.converter": "io.debezium.converters.BinaryDataConverter",
		
		"topic.prefix": svcName,
		"database.hostname": conf.Host,
		"database.port": conf.Port,
		"database.user": conf.Username,
		"database.password": conf.Password,
		"database.dbname": conf.Database,
	})
	if err != nil {
		log.Panic().Err(err).Msg("failed to marshal debezium cfg")
	}
	
	req, err := http.NewRequest(
		"PUT",
		"http://" + dbzHost + "/connectors/" + svcName + "-outbox/config",
		bytes.NewReader(payloadB),
	)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Panic().Err(err).Msg("failed to make debezium request")
	}
}