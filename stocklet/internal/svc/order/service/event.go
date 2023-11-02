package service

// todo:
// sort layout of project out
// API
// 
// 


// High Level Layers:
// Database Interface
// > Service Interface
// > > Event Processor (^ call service)

// API: (/api)
// > * Inbound RPC = Service
// > * Inbound HTTP | ^ Service Interface
// > * Inbound Kafka | ^ Event Process Funcs

// top level
// > svc interface
// >

// > /rpc
// > /http
// > /kafka


// - /order.go (svc inteface)
// - /database.go
// - /producer.go
// - /config.go

//		> /api 