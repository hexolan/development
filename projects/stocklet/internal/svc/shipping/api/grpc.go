// Copyright (C) 2024 Declan Teevan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package api

import (
	"google.golang.org/grpc"

	"github.com/hexolan/stocklet/internal/svc/shipping"
	"github.com/hexolan/stocklet/internal/pkg/serve"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/shipping/v1"
)

func PrepareGrpc(cfg *shipping.ServiceConfig, svc *shipping.ShippingService) *grpc.Server {
	svr := serve.NewGrpcServeBase(&cfg.Shared)
	pb.RegisterShippingServiceServer(svr, svc)
	return svr
}