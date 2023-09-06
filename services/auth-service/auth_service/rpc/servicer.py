import logging
import traceback
from typing import Type

from google.protobuf import empty_pb2
from grpc import RpcContext, StatusCode

from auth_service.models import AuthRepository, ServiceException, AuthToken
from auth_service.models.proto import auth_pb2, auth_pb2_grpc


class AuthServicer(auth_pb2_grpc.AuthServiceServicer):
    def __init__(self, svc_repo: Type[AuthRepository]) -> None:
        self._svc_repo = svc_repo
    
    def _apply_error(self, context: RpcContext, code: StatusCode, msg: str) -> None:
        context.set_code(code)
        context.set_details(msg)
    
    def _apply_unknown_error(self, context: RpcContext) -> None:
        self._apply_error(context, StatusCode.UNKNOWN, "unknown error occured")

    async def AuthWithPassword(self, request: auth_pb2.PasswordAuthRequest, context: RpcContext) -> auth_pb2.AuthToken:
        # validate the request inputs
        if request.user_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                details="user not provided"
            )
            return

        if request.password == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                details="password not provided"
            )
            return

        # Attempt to authenticate the user
        try:
            token = await self._svc_repo.auth_with_password(request.user_id, request.password)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # Convert token to protobuf
        return AuthToken.to_protobuf(token)

    async def SetPasswordAuth(self, request: auth_pb2.SetPasswordAuthMethod, context: RpcContext) -> empty_pb2.Empty:
        # validate the request inputs
        if request.user_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                details="user id not provided"
            )
            return
        
        if request.password == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                details="password not provided"
            )
            return

        # Attempt to create the auth method
        try:
            await self._svc_repo.set_password_auth_method(request.user_id, request.password)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # Success
        return empty_pb2.Empty()

    async def DeletePasswordAuth(self, request: auth_pb2.DeletePasswordAuthMethod, context: RpcContext) -> empty_pb2.Empty:
        # Ensure a user id is provided
        if request.user_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                details="user id not provided"
            )
            return
        
        # Attempt to delete the auth method
        try:
            await self._svc_repo.delete_password_auth_method(request.user_id, request.password)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # Success
        return empty_pb2.Empty()