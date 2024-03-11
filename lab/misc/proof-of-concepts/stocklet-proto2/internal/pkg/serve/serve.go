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

package serve

// Port Definitions
const (
	grpcPort string = "9090"
	gatewayPort string = "90"
)

// Get an address to a gRPC server
func AddrToGrpc(host string) string {
	return host + ":" + grpcPort
}

// Get an address to a gRPC-gateway interface
func AddrToGateway(host string) string {
	return host + ":" + gatewayPort
}