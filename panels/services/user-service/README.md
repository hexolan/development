# User Service

## Event Documentation

* Events Produced:
  * **Topic:** "``user``" | **Schema:** "``UserEvent``" protobuf
    * Type: ``"created"`` | Data: ``User``
    * Type: ``"updated"`` | Data: ``User``
    * Type: ``"deleted"`` | Data: ``User``

* Events Consumed:
  * N/A

## Configuration

### Environment Variables

**MongoDB:**

``MONGODB_URI`` (Required)

* e.g. "mongodb://mongo:mongo@127.0.0.1:27017/"

---

**Kafka:**

``KAFKA_BROKERS`` (Required)

* e.g. "127.0.0.1:9092" or "127.0.0.1:9092,127.0.0.1:9093"
