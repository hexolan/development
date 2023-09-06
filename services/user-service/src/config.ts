function getMongoURI(): string {
  if (process.env.MONGODB_URI === undefined) {
    throw new Error("mongodb_uri configuration not provided");
  } else {
    return process.env.MONGODB_URI
  }
}

function getKafkaBrokers(): string[] {
  if (process.env.KAFKA_BROKERS === undefined) {
    throw new Error("kafka_brokers configuration not provided");
  } else {
    return process.env.KAFKA_BROKERS.split(",")
  }
}

const mongodb_uri = getMongoURI();
const kafka_brokers = getKafkaBrokers();

export { mongodb_uri, kafka_brokers }