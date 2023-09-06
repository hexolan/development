import logging
import traceback
from typing import Type

from google.protobuf import empty_pb2
from grpc import RpcContext, StatusCode

from comment_service.models import CommentRepository, ServiceException, Comment, CommentCreate, CommentUpdate
from comment_service.models.proto import comment_pb2, comment_pb2_grpc


class CommentServicer(comment_pb2_grpc.CommentServiceServicer):
    def __init__(self, svc_repo: Type[CommentRepository]) -> None:
        self._svc_repo = svc_repo
    
    def _apply_error(self, context: RpcContext, code: StatusCode, msg: str) -> None:
        context.set_code(code)
        context.set_details(msg)
    
    def _apply_unknown_error(self, context: RpcContext) -> None:
        self._apply_error(context, StatusCode.UNKNOWN, "unknown error occured")

    async def CreateComment(self, request: comment_pb2.CreateCommentRequest, context: RpcContext) -> comment_pb2.Comment:
        # vaLidate the request inputs
        if request.post_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="post not provided"
            )
            return
        
        if request.author_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="author not provided"
            )
            return
        
        if request.data == None:
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="malformed request"
            )
            return
        
        if request.data.message == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment message not provided"
            )
            return
        
        # convert to service model from protobuf
        try:
            data = CommentCreate.from_protobuf(request)
            comment = await self._svc_repo.create_comment(data)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return

        # convert comment to protobuf form
        return Comment.to_protobuf(comment)

    async def UpdateComment(self, request: comment_pb2.UpdateCommentRequest, context: RpcContext) -> comment_pb2.Comment:
        # vaLidate the request inputs
        if request.id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment not provided"
            )
            return
        elif not request.id.isnumeric():
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="invalid comment id provided"
            )
            return
        
        if request.data == None:
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="malformed request"
            )
            return
        
        if request.data.message == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment message not provided"
            )
            return
        
        # convert to service model from protobuf
        try:
            comment_id = int(request.id)
            data = CommentUpdate.from_protobuf(request)
            comment = await self._svc_repo.update_comment(comment_id, data)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return

        # convert comment to protobuf form
        return Comment.to_protobuf(comment)

    async def DeleteComment(self, request: comment_pb2.DeleteCommentRequest, context: RpcContext) -> empty_pb2.Empty:
        # vaLidate the request inputs
        if request.id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment not provided"
            )
            return
        
        if not request.id.isnumeric():
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="invalid comment id provided"
            )
            return

        # attempt to delete the comment
        try:
            comment_id = int(request.id)
            await self._svc_repo.delete_comment(comment_id)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return

        return empty_pb2.Empty()

    async def GetPostComments(self, request: comment_pb2.GetPostCommentsRequest, context: RpcContext) -> comment_pb2.PostComments:
        # vaLidate the request inputs
        if request.post_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="post id not provided"
            )
            return
        
        # attempt to get the comments
        try:
            comments = await self._svc_repo.get_post_comments(request.post_id)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # convert to protobuf
        return comment_pb2.PostComments(data=[Comment.to_protobuf(comment) for comment in comments])