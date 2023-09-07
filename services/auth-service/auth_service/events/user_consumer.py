import logging
from typing import Type

from aiokafka import AIOKafkaConsumer

from auth_service.models import Config, AuthDBRepository
from auth_service.models.proto import user_pb2


class UserEventConsumer:
    """Consumer class responsible for 'user' events.
    
    Attributes:
        _db_repo (AuthDBRepository): The repository interface for modifying data.
        _consumer (aiokafka.AIOKafkaConsumer): The underlying Kafka instance.

    """

    def __init__(self, config: Config, db_repo: Type[AuthDBRepository]) -> None:
        # todo: make a base consumer class for abstraction (_process_msg, with class types e.g. PROTO_MSG = user_pb2.UserEvent to convert using that attr)
        """Initialise the event consumer.
        
        Args:
            config (Config): The app configuration instance (to access brokers list).
            db_repo (AuthDBRepository): The repository interface for updating data.

        """
        self._db_repo = db_repo
        self._consumer = AIOKafkaConsumer(
            "user",
            bootstrap_servers=config.kafka_brokers,
            group_id="auth-service"
        )

    async def start(self) -> None:
        """Begin consuming messages."""
        await self._consumer.start()
        try:
            async for msg in self._consumer:
                await self._process_msg(msg)
        except Exception as e:
            logging.error("error whilst consuming messages:", e)
        finally:
            await self._consumer.stop()

    async def _process_msg(self, msg) -> None:
        """Process a recieved message.
        
        The messages are deserialise from bytes into their protobuf form,
        then passed to the `_process_event` method to handle the logic.
        
        Args:
            msg (kafka.Message): The event to process.
        """
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
        """Process a recieved event.

        In response to a User deleted event, delete any auth methods
        this service has in relation to that user.

        Args:
            event (user_pb2.UserEvent): The decoded protobuf message.
        
        """
        if event.type == "deleted":
            await self._db_repo.delete_password_auth_method(event.data.id)
            logging.info("succesfully processed UserEvent - type: 'deleted'")