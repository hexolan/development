from comment_service.models import Config, CommentDBRepository
from comment_service.events.producer import CommentEventProducer
from comment_service.events.post_consumer import PostEventConsumer
from comment_service.events.user_consumer import UserEventConsumer


class EventConsumersWrapper:
    def __init__(self, post_consumer: PostEventConsumer, user_consumer: UserEventConsumer) -> None:
        self._post_consumer = post_consumer
        self._user_consumer = user_consumer

    async def start(self) -> None:
        await self._post_consumer.start()
        await self._user_consumer.start()


def create_consumers(config: Config, db_repo: CommentDBRepository) -> EventConsumersWrapper:
    post_consumer = PostEventConsumer(config, db_repo)
    user_consumer = UserEventConsumer(config, db_repo)
    return EventConsumersWrapper(post_consumer=post_consumer, user_consumer=user_consumer)


async def create_producer(config: Config) -> CommentEventProducer:
    producer = CommentEventProducer(config)
    await producer.start()
    return producer