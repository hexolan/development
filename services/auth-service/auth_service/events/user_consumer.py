import logging
from typing import Type

from aiokafka import AIOKafkaConsumer

from auth_service.models import Config, AuthDBRepository
from auth_service.models.proto import user_pb2


class UserEventConsumer:
    def __init__(self, config: Config, db_repo: Type[AuthDBRepository]) -> None:
        self._db_repo = db_repo
        self._consumer = AIOKafkaConsumer(
            "user",
            bootstrap_servers=config.kafka_brokers,
            group_id="auth-service"
        )

    async def start(self) -> None:
        await self._consumer.start()
        try:
            async for msg in self._consumer:
                await self._process_msg(msg)
        except Exception as e:
            logging.error("error whilst consuming messages:", e)
        finally:
            await self._consumer.stop()

    async def _process_msg(self, msg) -> None:
        # Attempt to sereialise the msg value as a UserEvent
        try:    
            event = user_pb2.UserEvent()
            event.ParseFromString(msg.value)
            assert event.type != ""
            await self._process_event(event)
        except AssertionError:
            logging.error("invalid UserEvent recieved")
            return
        except Exception as e:
            logging.error("error whilst processing UserEvent:", e)
            return

    async def _process_event(self, event: user_pb2.UserEvent) -> None:
        if event.type == "deleted":
            await self._db_repo.delete_password_auth_method(event.data.id)
            logging.info("succesfully processed UserEvent - type: 'deleted'")