import logging
from typing import Type

from databases import Database

from comment_service.models import Config, CommentDBRepository
from comment_service.events import CommentEventProducer
from comment_service.postgres.repository import ServiceDBRepository


async def connect_database(config: Config) -> Database:
    db = Database(config.postgres_dsn)
    try:
        await db.connect()
    except Exception:
        logging.error("failed to connect to postgresql database")
        raise
    
    return db


async def create_db_repository(config: Config, event_producer: CommentEventProducer) -> Type[CommentDBRepository]:
    db = await connect_database(config)
    return ServiceDBRepository(db, event_producer)