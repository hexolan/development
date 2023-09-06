import logging
from typing import Type

from aiokafka import AIOKafkaConsumer

from comment_service.models import Config, CommentDBRepository
from comment_service.models.proto import post_pb2


class PostEventConsumer:
    def __init__(self, config: Config, db_repo: Type[CommentDBRepository]) -> None:
        self._db_repo = db_repo
        self._consumer = AIOKafkaConsumer(
            "post",
            bootstrap_servers=config.kafka_brokers,
            group_id="comment-service"
        )

    async def start(self) -> None:
        await self._consumer.start()
        try:
            async for msg in self._consumer:
                await self._process_msg(msg)
        except Exception as e:
            logging.error("error whilst attempting to consume PostEvents:", e)
            await self._consumer.stop()
            raise

    async def _process_msg(self, msg) -> None:
        # Attempt to serialise the msg value as a PostEvent
        try:    
            event = post_pb2.PostEvent.ParseFromString(msg.value)
            assert event.data is not None
            await self._process_event(event)
        except AssertionError:
            logging.error("invalid PostEvent recieved")
            pass
        except Exception as e:
            logging.error("error whilst processing PostEvent:", e)
            pass

    async def _process_event(self, event: post_pb2.PostEvent) -> None:
        if event.type == "deleted":
            assert event.data.id != ""
            await self._db_repo.delete_post_comments(event.data.id)
            logging.info("succesfully processed PostEvent - type: 'deleted'")