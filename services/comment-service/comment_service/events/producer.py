from aiokafka import AIOKafkaProducer

from comment_service.models.config import Config
from comment_service.models.service import Comment
from comment_service.models.proto import comment_pb2


class CommentEventProducer:
    def __init__(self, config: Config) -> None:
        self._producer = AIOKafkaProducer(
            bootstrap_servers=config.kafka_brokers,
            client_id="auth-service"
        )

    async def start(self) -> None:
        await self._producer.start()

    async def _send_event(self, event: comment_pb2.CommentEvent) -> None:
        future = await self._producer.send(
            topic="comment",
            value=event.SerializeToString(),
        )
        # todo: ensure this future gets executed

    async def send_created_event(self, comment: Comment) -> None:
        event = comment_pb2.CommentEvent(
            type="created",
            data=Comment.to_protobuf(comment)
        )
        await self._send_event(event)

    async def send_updated_event(self, comment: Comment) -> None:
        event = comment_pb2.CommentEvent(
            type="updated",
            data=Comment.to_protobuf(comment)
        )
        await self._send_event(event)

    async def send_deleted_event(self, comment: Comment) -> None:
        event = comment_pb2.CommentEvent(
            type="deleted",
            data=Comment.to_protobuf(comment)
        )
        await self._send_event(event)