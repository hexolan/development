import logging
from typing import Type

from aiokafka import AIOKafkaConsumer

from comment_service.models import Config, CommentDBRepository
from comment_service.models.proto import user_pb2


class UserEventConsumer:
    def __init__(self, config: Config, db_repo: Type[CommentDBRepository]) -> None:
        self._db_repo = db_repo
        self._consumer = AIOKafkaConsumer(
            "user",
            bootstrap_servers=config.kafka_brokers,
            group_id="comment-service"
        )

    async def start(self) -> None:
        await self._consumer.start()
        try:
            async for msg in self._consumer:
                await self._process_msg(msg)
        except Exception as e:
            logging.error("error whilst attempting to consume UserEvents:", e)
            await self._consumer.stop()
            raise

    async def _process_msg(self, msg) -> None:
        # Attempt to serialise the msg value as a UserEvent
        try:    
            event = user_pb2.UserEvent.ParseFromString(msg.value)
            assert event.data is not None
            await self._process_event(event)
        except AssertionError:
            logging.error("invalid UserEvent recieved")
            pass
        except Exception as e:
            logging.error("error whilst processing UserEvent:", e)
            pass

    async def _process_event(self, event: user_pb2.UserEvent) -> None:
        if event.type == "deleted":
            assert event.data.id != ""
            await self._db_repo.delete_user_comments(event.data.id)
            logging.info("succesfully processed UserEvent - type: 'deleted'")