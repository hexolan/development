import logging
from typing import Type

from databases import Database

from auth_service.models import Config, AuthDBRepository
from auth_service.postgres.repository import ServiceDBRepository


async def connect_database(config: Config) -> Database:
    db = Database(config.postgres_dsn)
    try:
        await db.connect()
    except Exception:
        logging.error("failed to connect to postgresql database")
        raise
    
    return db


async def create_db_repository(config: Config) -> Type[AuthDBRepository]:
    db = await connect_database(config)
    return ServiceDBRepository(db)