from enum import Enum, auto

from grpc import RpcContext, StatusCode


class ServiceErrorCode(Enum):
    INVALID_ARGUMENT = auto()
    CONFLICT = auto()
    NOT_FOUND = auto()
    INVALID_CREDENTIALS = auto()
    SERVICE_ERROR = auto()

    __RPC_CODE_MAP__ = {
        INVALID_ARGUMENT: StatusCode.INVALID_ARGUMENT,
        CONFLICT: StatusCode.ALREADY_EXISTS,
        NOT_FOUND: StatusCode.NOT_FOUND,
        INVALID_CREDENTIALS: StatusCode.UNAUTHENTICATED,
        SERVICE_ERROR: StatusCode.INTERNAL
    }

    def to_rpc_code(self) -> StatusCode:
        return self.__class__.__RPC_CODE_MAP__.get(self.value, StatusCode.UNKNOWN)


class ServiceException(Exception):
    def __init__(self, message: str, error_code: ServiceErrorCode) -> None:
        super().__init__(message)
        self.message = message
        self.error_code = error_code

    def apply_to_rpc(self, context: RpcContext) -> None:
        context.set_code(self.error_code.to_rpc_code())
        context.set_details(self.message)
