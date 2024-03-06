package main

import (
	"io"
	"log"
	"bytes"
	"net/http"
	"encoding/json"
)

func main() {
	const debeziumHost string = "debezium:8083"

	client := &http.Client{}

	connectorCfg := map[string]string{
		"connector.class": "io.debezium.connector.postgresql.PostgresConnector",
		"tasks.max": "1",
		"database.hostname": "test-service-postgres",
		"database.port": "5432",
		"database.user": "postgres",
		"database.password": "postgres",
		"database.dbname": "postgres",
		"topic.prefix": "test-service",
		"table.include.list": "public.event_outbox",
		"plugin.name": "pgoutput",
		"transforms": "outbox",
		"transforms.outbox.type": "io.debezium.transforms.outbox.EventRouter",
		"transforms.outbox.route.topic.replacement": "${routedByValue}",
		"value.converter": "io.debezium.converters.BinaryDataConverter",
	}

	payloadB, err := json.Marshal(connectorCfg)
	if err != nil {
		panic("failed to marshal connector config to JSON")
	}
	
	url := "http://" + debeziumHost + "/connectors/" + "test-service-outbox" + "/config"
	req, err := http.NewRequest("PUT", url, bytes.NewReader(payloadB))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// response
	bodyB, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Println(bodyB)
	log.Println(string(bodyB))

	var body struct{}
	err = json.Unmarshal(bodyB, &body)
	if err == nil { 
		log.Println(body)
	}

	//if res.StatusCode == http.StatusOk {}
}