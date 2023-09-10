import logging

from comment_service.models.proto import user_pb2
from comment_service.events.base_consumer import EventConsumer


class UserEventConsumer(EventConsumer):
    CONSUMER_TOPIC = "user"
    CONSUMER_EVENT_TYPE = user_pb2.UserEvent

    async def _process_event(self, event: user_pb2.UserEvent) -> None:
        if event.type == "deleted":
            assert event.data.id != ""
            await self._db_repo.delete_user_comments(event.data.id)
            logging.info("succesfully processed UserEvent (type: 'deleted')")