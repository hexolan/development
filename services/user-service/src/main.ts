import mongoose from "mongoose";

import { mongodb_uri } from "./config";
import userProducer from "./kafka/producer";
import serveRPC from "./rpc/server";

async function main(): Promise<void> {
  await mongoose.connect(mongodb_uri);
  await userProducer.connect()

  void serveRPC();
}

void main();