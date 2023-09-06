import logging
from typing import Type

import grpc
from grpc_health.v1 import health, health_pb2_grpc

from auth_service.models import AuthRepository
from auth_service.models.proto import auth_pb2_grpc
from auth_service.rpc.servicer import AuthServicer


class RPCServerWrapper:
    def __init__(self, svc_repo: Type[AuthRepository]) -> None:
        self._grpc_server = grpc.aio.server()
        self._grpc_server.add_insecure_port("[::]:9090")

        auth_servicer = AuthServicer(svc_repo)
        auth_pb2_grpc.add_AuthServiceServicer_to_server(auth_servicer, self._grpc_server)

        health_pb2_grpc.add_HealthServicer_to_server(health.aio.HealthServicer(), self._grpc_server)

    async def start(self) -> None:
        logging.info("attempting to serve RPC...")
        await self._grpc_server.start()
        await self._grpc_server.wait_for_termination()


def create_rpc_server(svc_repo: Type[AuthRepository]) -> RPCServerWrapper:
    return RPCServerWrapper(svc_repo)