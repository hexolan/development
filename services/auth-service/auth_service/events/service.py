from typing import Type

from auth_service.models import Config, AuthDBRepository
from auth_service.events.user_consumer import UserEventConsumer


class EventConsumersWrapper:
    def __init__(self, user_consumer: UserEventConsumer) -> None:
        self._user_consumer = user_consumer

    async def start(self) -> None:
        await self._user_consumer.start()


def create_consumers(config: Config, db_repo: Type[AuthDBRepository]) -> EventConsumersWrapper:
    user_consumer = UserEventConsumer(config, db_repo)
    return EventConsumersWrapper(user_consumer=user_consumer)