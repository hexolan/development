from typing import Type

import redis.asyncio as redis

from comment_service.models import Config, CommentRepository
from comment_service.redis.repository import ServiceRedisRepository


async def connect_redis(config: Config) -> None:
    host, port = config.redis_host, 5432
    try:
        host, port = config.redis_host.split(":")
        port = int(port)
    except Exception:
        pass

    conn = redis.Redis(host=host, port=port, password=config.redis_pass)
    await conn.ping()
    return conn


async def create_redis_repository(config: Config, downstream_repo: Type[CommentRepository]) -> Type[CommentRepository]:
    redis_conn = await connect_redis(config)
    return ServiceRedisRepository(redis_conn=redis_conn, downstream_repo=downstream_repo)