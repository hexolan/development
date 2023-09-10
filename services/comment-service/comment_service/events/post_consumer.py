import logging

from comment_service.models.proto import post_pb2
from comment_service.events.base_consumer import EventConsumer


class PostEventConsumer(EventConsumer):
    CONSUMER_TOPIC = "post"
    CONSUMER_EVENT_TYPE = post_pb2.PostEvent

    async def _process_event(self, event: post_pb2.PostEvent) -> None:
        if event.type == "deleted":
            assert event.data.id != ""
            await self._db_repo.delete_post_comments(event.data.id)
            logging.info("succesfully processed PostEvent (type: 'deleted')")